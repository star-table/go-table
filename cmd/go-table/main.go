package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/config/file"

	"github.com/spf13/cast"

	sentryClint "github.com/getsentry/sentry-go"
	nacosconfig "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	encoding "github.com/star-table/go-common/pkg/encoding"
	zapLog "github.com/star-table/go-common/pkg/log"
	"github.com/star-table/go-table/internal/conf"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	//TODO:生成新模板后请更改这个名字，这个名字和注册、配置联系在一起
	Name = "go-table"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()

	// 环境，local就不会去注册
	env = ""

	nacosHost, nacosPort, nacosNamespace, nacosUser, nacosPassword string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf config.yaml")

	nacosHost = os.Getenv("REGISTER_HOST")
	nacosPort = os.Getenv("REGISTER_PORT")
	nacosNamespace = os.Getenv("REGISTER_NAMESPACE")
	nacosUser = os.Getenv("REGISTER_USERNAME")
	nacosPassword = os.Getenv("REGISTER_PASSWORD")

	rand.Seed(time.Now().UnixNano())
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, registry registry.Registrar) *kratos.App {
	options := []kratos.Option{
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	}
	if env != "local" {
		options = append(options, kratos.Registrar(registry))
	}
	return kratos.New(options...)
}

func main() {
	bc := loadConfig()
	// 设置下环境
	env = bc.Env

	initThirdParty(bc, env)

	logger := zapLog.InitDefaultLog(bc.Log.Path, sentryClint.CurrentHub().Client())

	sc, cc := getNacosServerAndClientConfig()
	app, cleanup, err := initApp(bc.Server, bc.Data, sc, cc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	encoding.Init()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(jaegerConf *conf.ThirdParty_Jaeger) (*trace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerConf.HttpEndpoint)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(float64(jaegerConf.SamplerParam)))),
		trace.WithBatcher(exp),
		trace.WithResource(
			resource.NewSchemaless(
				semconv.ServiceNameKey.String(Name),
				attribute.String("version", Version),
			)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func initThirdParty(bc *conf.Bootstrap, env string) {
	if bc.ThirdParty.Jaeger != nil && bc.ThirdParty.Jaeger.HttpEndpoint != "" {
		_, err := tracerProvider(bc.ThirdParty.Jaeger)
		if err != nil {
			panic(err)
		}
	}

	initSentry(bc.ThirdParty.Sentry, env)
}

func loadConfig() *conf.Bootstrap {
	flag.Parse()
	if flagconf != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}

		var bc = &conf.Bootstrap{}
		if err := c.Scan(bc); err != nil {
			panic(err)
		}

		return bc
	}

	return loadNacosConfig()
}

func loadNacosConfig() *conf.Bootstrap {
	client, err := getNacosConfigClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	configSource := nacosconfig.NewConfigSource(client, nacosconfig.WithGroup("DEFAULT_GROUP"), nacosconfig.WithDataID(Name))

	c := config.New(
		config.WithSource(
			configSource,
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)

	if err := c.Load(); err != nil {
		log.Error(err)
		return nil
	}

	bc := &conf.Bootstrap{}
	if err := c.Scan(bc); err != nil {
		log.Error(err)
		return nil
	}

	return bc
}

func getNacosServerAndClientConfig() ([]constant.ServerConfig, constant.ClientConfig) {
	return []constant.ServerConfig{
			*constant.NewServerConfig(nacosHost, cast.ToUint64(nacosPort)),
		},
		constant.ClientConfig{
			AppName:             Name,
			NamespaceId:         nacosNamespace, //namespace id
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogLevel:            "error",
			Username:            nacosUser,
			Password:            nacosPassword,
		}
}

func getNacosConfigClient() (config_client.IConfigClient, error) {
	sc, cc := getNacosServerAndClientConfig()
	// a more graceful way to create naming client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func initSentry(sc *conf.ThirdParty_Sentry, env string) {
	if sc != nil && sc.Dsn != "" {
		err := sentryClint.Init(sentryClint.ClientOptions{
			Dsn:              sc.Dsn,
			AttachStacktrace: true, // recommended
			Environment:      env,
		})
		if err != nil {
			panic(err)
		}
	}
}

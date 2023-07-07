package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/go-kratos/kratos/v2/middleware/logging"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	gormopentracing "github.com/star-table/go-common/pkg/gorm/opentracing"
	middlewareMeta "github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewGormDBData,
	NewRedisCmd,
	NewLockRepo,
	NewTableRepo,
	NewRowRepo,
	NewTableCache,
	NewAppRepo,
	NewDatacenterRepo,
	NewUserCenterRepo,
	NewOrgColumnsRepo,
	NewOrgColumnsCache,
	NewPermissionRepo,
	NewGoPushRepo,
	NewFormRepo,
	NewProjectRepo,
)

// Data .
type Data struct {
	mysqlLcGo *gorm.DB
	postgres  map[string]*gorm.DB
	redisCli  redis.Cmdable
	snowFlake *snowflake.Node
}

type GormDBData struct {
	mysqlLcGo *gorm.DB
	mysqlLc   *gorm.DB
	postgres  map[string]*gorm.DB
}

// NewData .
func NewData(dbData *GormDBData, redisCli redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Errorf("snowflake new node failed, the error is '%v'", err)
		return nil, nil, err
	}

	d := &Data{
		mysqlLcGo: dbData.mysqlLcGo,
		postgres:  dbData.postgres,
		redisCli:  redisCli,
		snowFlake: node,
	}
	cleanup := func() {
		db, err := d.mysqlLcGo.DB()
		if err == nil {
			if err = db.Close(); err != nil {
				log.Errorf("Close db error:%v", err)
			}
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}

func NewGormDBData(c *conf.Data) *GormDBData {
	return &GormDBData{
		mysqlLcGo: NewMysqlLcGo(c),
		postgres:  NewPostgres(c),
	}
}

func NewMysqlLcGo(c *conf.Data) *gorm.DB {
	cc := c.Database["mysql_lc_go"]
	db := openDB(cc.Driver, cc.Dsn)
	//fmt.Println(db.AutoMigrate(&po.Table{}, &po.TableColumn{}, &po.OrgColumn{}))

	return db
}

func NewPostgres(c *conf.Data) map[string]*gorm.DB {
	dbs := make(map[string]*gorm.DB)
	for s, database := range c.Database {
		if strings.Contains(s, "postgres") {
			dbs[s] = openDB(database.Driver, database.Dsn).Debug()
		}
	}

	return dbs
}

func openDB(driver, dsn string) *gorm.DB {
	dialector := getDialector(driver, dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Errorf("Got error when connect database, the error is '%v'", err)
		return nil
	}

	err = db.Use(gormopentracing.New())
	if err != nil {
		log.Errorf("Got error when connect database, the error is '%v'", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("Got error when connect database, the error is '%v'", err)
		return nil
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(32)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func getDialector(driver, dsn string) gorm.Dialector {
	switch driver {
	case "mysql":
		return mysql.Open(dsn)
	case "postgres":
		return postgres.Open(dsn)
	default:
		panic("no support driver")
	}
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "data"))
	var client redis.Cmdable
	if conf.Redis.IsSentinel {
		client = newSentinelRedisCmd(conf)
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:         conf.Redis.Addr,
			Password:     conf.Redis.Password,
			DB:           int(conf.Redis.Db),
			ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
			WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
			DialTimeout:  time.Second * 3,
			PoolSize:     10,
		})
	}

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Errorf("redis connect error: %v", err)
		return nil
	}

	return client
}

func newSentinelRedisCmd(conf *conf.Data) redis.Cmdable {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    conf.Redis.MasterName,
		SentinelAddrs: []string{conf.Redis.Addr},
		Password:      conf.Redis.Password,
		DB:            int(conf.Redis.Db),
		ReadTimeout:   conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout:  conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:   time.Second * 3,
		PoolSize:      10,
	})
}

func NewDiscovery(sc []constant.ServerConfig, cc constant.ClientConfig) registry.Discovery {
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		log.Error(err)
		return nil
	}

	r := nacos.New(client)

	return r
}

func getHttpConn(serverName string, r registry.Discovery, logger log.Logger) (*http.Client, error) {
	options := []http.ClientOption{
		http.WithEndpoint(serverName),
		http.WithTimeout(15 * time.Second),
		http.WithMiddleware(
			recovery.Recovery(),
			middlewareMeta.Client(),
			tracing.Client(),
			logging.Client(logger),
		),
	}
	if strings.Contains(serverName, "discovery") {
		options = append(options, http.WithDiscovery(r))
	}

	return http.NewClient(
		context.Background(),
		options...,
	)
}

func convertToQueryParams(params map[string]interface{}) string {
	query := url.Values{}
	for key, value := range params {
		query.Set(key, cast.ToString(value))
	}
	return "?" + query.Encode()
}

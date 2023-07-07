module github.com/star-table/go-table

go 1.16

require (
	github.com/bsm/redislock v0.7.2
	github.com/bwmarrin/snowflake v0.3.0
	github.com/getsentry/sentry-go v0.13.0
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20220310144244-ac99a5c877c4
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20220531020131-5de1f081f636
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20220310144244-ac99a5c877c4
	github.com/go-kratos/kratos/v2 v2.3.0
	github.com/go-kratos/sentry v0.0.0-20211021071616-de3a2011c4e4
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/nacos-group/nacos-sdk-go v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/spf13/cast v1.5.0
	github.com/star-table/go-common v1.0.0
	github.com/star-table/interface v0.0.0-20230707032058-aa3d85d8a825
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/jaeger v1.4.1
	go.opentelemetry.io/otel/sdk v1.7.0
	go.uber.org/zap v1.21.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/datatypes v1.0.7 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.3.4
	gorm.io/gen v0.2.43
	gorm.io/gorm v1.23.6
)

replace gorm.io/gorm => github.com/jiangchangren/gorm v0.0.0-20230511075058-4d41a5b2acb6

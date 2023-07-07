//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/data"
	"github.com/star-table/go-table/internal/server"
	"github.com/star-table/go-table/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, []constant.ServerConfig, constant.ClientConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

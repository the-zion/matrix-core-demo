// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"cube-core/internal/biz"
	"cube-core/internal/conf"
	"cube-core/internal/data"
	"cube-core/internal/server"
	"cube-core/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "github.com/gorilla/handlers"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, jwt *conf.JWT, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	profileRepo := data.NewProfileRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, profileRepo, logger, jwt)
	realWorldService := service.NewRealWorldService(userUseCase)
	httpServer := server.NewHTTPServer(confServer, jwt, realWorldService, logger)
	grpcServer := server.NewGRPCServer(confServer, realWorldService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

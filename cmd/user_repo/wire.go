// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"kkako_video/internal/user_repo/handler"
	"kkako_video/internal/user_repo/repo"
	"kkako_video/internal/user_repo/server"
	"kkako_video/pkg/app"
)

func initApp(v *viper.Viper) (*app.App, error) {
	panic(wire.Build(
		newApp,
		server.Provider,
		handler.ProviderSet,
		repo.ProviderSet,
	),
	)
}

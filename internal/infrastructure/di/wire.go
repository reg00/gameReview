//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Reg00/gameReview/internal/infrastructure/appctx"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/Reg00/gameReview/internal/infrastructure/web"
	"github.com/Reg00/gameReview/internal/infrastructure/web/routes"
	"github.com/google/wire"
)

func InitWebServer() (*web.Server, error) {
	wire.Build(
		routes.Register,
		web.Register,
		appctx.Register,
		config.Register,
	)

	return &web.Server{}, nil
}

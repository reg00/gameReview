//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/infrastructure/appctx"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/Reg00/gameReview/internal/infrastructure/igdb"
	"github.com/Reg00/gameReview/internal/infrastructure/storage/gorm"
	"github.com/Reg00/gameReview/internal/infrastructure/storage/redis"
	"github.com/Reg00/gameReview/internal/infrastructure/web"
	"github.com/Reg00/gameReview/internal/infrastructure/web/routes"
	"github.com/google/wire"
)

func InitWebServer() (*web.Server, error) {
	wire.Build(
		wire.Bind(new(port.GameSearcher), new(*igdb.IgdbClient)),
		wire.Bind(new(port.Storager), new(*gorm.GormStorage)),
		wire.Bind(new(port.Cacher), new(*redis.RedisStorage)),

		routes.Register,
		web.Register,
		appctx.Register,
		config.Register,
		igdb.Register,
		gorm.Register,
		redis.Register,
	)

	return &web.Server{}, nil
}

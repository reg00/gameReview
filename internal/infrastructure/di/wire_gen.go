// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Reg00/gameReview/internal/infrastructure/appctx"
	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/Reg00/gameReview/internal/infrastructure/web"
	"github.com/Reg00/gameReview/internal/infrastructure/web/routes"
)

// Injectors from wire.go:

func InitWebServer() (*web.Server, error) {
	coreContext := appctx.Register()
	configuration := config.Register()
	engine := routes.Register()
	server := web.Register(coreContext, configuration, engine)
	return server, nil
}

package routes

import (
	"net/http"

	docs "github.com/Reg00/gameReview/docs"
	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/infrastructure/web/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @title Swagger Example API
// @version 1.0
// @description This is a sample game review server.
func Register(igdb port.GameSearcher, s port.Storager) (*gin.Engine, error) {
	gh, err := handler.NewGameHandler(igdb)
	if err != nil {
		return nil, err
	}

	rh, err := handler.NewReviewHandler(igdb, s)
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(handler.ErrorHandler)

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("api/v1")
	{
		g := v1.Group("games")
		{
			g.GET("/", gh.GetGamesByNameHandlerFunc)
			g.GET("/:id", gh.GetGameById)
		}
		r := v1.Group("reviews")
		{
			r.POST("/", rh.AddReview)
			r.GET("/:id", rh.GetReviewById)
		}
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, "PAGE_NOT_FOUND")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r, nil
}

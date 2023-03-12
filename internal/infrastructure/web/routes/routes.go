package routes

import (
	"net/http"

	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/infrastructure/web/handler"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func Register(igdb port.GameSearcher) (*gin.Engine, error) {
	h, err := handler.New(igdb)
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/games", h.GetGamesByNameHandlerFunc)
	r.GET("/games/:id", h.GetGameById)
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)

	//apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {

	// }

	return r, nil
}

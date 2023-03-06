package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)

	//apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {

	// }

	return r
}

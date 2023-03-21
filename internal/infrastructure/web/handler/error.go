package handler

import (
	"errors"
	"net/http"

	"github.com/Reg00/gameReview/internal/domain/dto/httperr"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch err.Err {
		case httperr.ErrNotFound:
			c.IndentedJSON(http.StatusNotFound, err.Err.Error())
			return
		}

		if errors.Is(err.Err, httperr.ErrNotFound) {
			c.IndentedJSON(http.StatusNotFound, err.Err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, "Unhandled exception: "+err.Error())
	}
}

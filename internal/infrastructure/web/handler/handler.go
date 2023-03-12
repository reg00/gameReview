package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	grs *service.GameReviewService
}

func New(igdb port.GameSearcher) (*Handler, error) {
	grs := service.NewGameReviewService(&igdb)
	h := &Handler{}
	h.grs = grs
	return h, nil
}

func (h *Handler) GetGameById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	game, err := h.grs.GetGameById(id)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, game)
}

func (h *Handler) GetGamesByNameHandlerFunc(c *gin.Context) {
	name := c.Query("name")
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	games, err := h.grs.GetGamesByName(offset, limit, name)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, games)
}

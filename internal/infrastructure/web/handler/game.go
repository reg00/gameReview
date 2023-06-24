package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gs *service.GameService
}

func NewGameHandler(
	igdb port.GameSearcher) (*GameHandler, error) {
	h := &GameHandler{}
	h.gs = service.NewGameService(&igdb)
	return h, nil
}

// @Summary get game by id
// @Schemes
// @Description getting a game by game id
// @Tags games
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Game
// @Router /games/{id} [get]
func (h *GameHandler) GetGameById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, err := h.gs.GetGameById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

// @Summary get list of games by name
// @Schemes
// @Description return a list of game, searched by name
// @Tags games
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param name query string false "name"
// @Success 200 {object} []models.Game
// @Router /games [get]
func (h *GameHandler) GetGamesByNameHandlerFunc(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	games, err := h.gs.GetGamesByName(offset, limit, name)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Reg00/gameReview/internal/domain/dto/storage"
	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	grs     *service.GameReviewService
	storage port.Storager
}

func New(
	igdb port.GameSearcher,
	s port.Storager) (*Handler, error) {
	grs := service.NewGameReviewService(&igdb)
	h := &Handler{}
	h.grs = grs
	h.storage = s
	return h, nil
}

// @Summary get game by id
// @Schemes
// @Description getting a game by game id
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.Game
// @Router /games/{id} [get]
func (h *Handler) GetGameById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, err := h.grs.GetGameById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, game)
}

// @Summary get list of games by name
// @Schemes
// @Description return a list of game, searched by name
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Param name query string false "name"
// @Success 200 {object} []dto.Game
// @Router /games [get]
func (h *Handler) GetGamesByNameHandlerFunc(c *gin.Context) {
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

	games, err := h.grs.GetGamesByName(offset, limit, name)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

// @Summary add game review
// @Schemes
// @Description  add game review
// @Accept json
// @Produce json
// @Param review body storage.Review true "review info"
// @Success 200 {object} []dto.Game
// @Router /games [get]
func (h *Handler) AddReview(c *gin.Context) {
	var review storage.Review

	if err := c.BindJSON(&review); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	r, err := h.storage.AddReview(&review)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, r)
}

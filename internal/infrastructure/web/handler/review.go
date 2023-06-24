package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Reg00/gameReview/internal/domain/models"
	"github.com/Reg00/gameReview/internal/domain/port"
	"github.com/Reg00/gameReview/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	gs      *service.GameService
	storage *service.ReviewService
}

func NewReviewHandler(
	igdb port.GameSearcher,
	s port.Storager) (*ReviewHandler, error) {
	h := &ReviewHandler{}
	h.gs = service.NewGameService(&igdb)
	h.storage = service.NewReviewService(&s, &igdb)
	return h, nil
}

// @Summary add game review
// @Schemes
// @Description  add game review
// @Accept json
// @Produce json
// @Param review body models.AddReview true "review info"
// @Success 200 {object} models.GetReview
// @Router /reviews [post]
func (h *ReviewHandler) AddReview(c *gin.Context) {

	var addReview models.AddReview

	if err := c.BindJSON(&addReview); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	review, err := h.storage.AddReview(&addReview)
	if err != nil {
		fmt.Println(err.Error())
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, review)
}

// @Summary get game review by id
// @Schemes
// @Description  get game review by id
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.GetReview
// @Router /reviews/{id} [get]
func (h *ReviewHandler) GetReviewById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	review, err := h.storage.GetReviewById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, review)
}

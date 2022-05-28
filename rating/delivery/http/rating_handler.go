package http

import (
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	ratingUsecase domain.RatingUsecase
}

func NewRatingHandler(r *gin.RouterGroup, ratingUsecase domain.RatingUsecase, jwtMiddleware gin.HandlerFunc) {
	handler := &RatingHandler{ratingUsecase: ratingUsecase}
	api := r.Group("/rating")
	{
		api.POST("/:id", jwtMiddleware, handler.Create)
	}
}

func (h *RatingHandler) Create(c *gin.Context) {
	var input domain.CreateRating
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	var uri domain.BindingUriSpeakerId
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("userLoggedin").(uint)
	input.SpeakerId = uri.Id
	input.UserId = userId
	if err := h.ratingUsecase.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully rated speaker", nil))
}

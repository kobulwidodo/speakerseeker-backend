package http

import (
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type SpeakerHandler struct {
	speakerUsecase domain.SpeakerUsecase
}

func NewSpeakerHandler(r *gin.RouterGroup, su domain.SpeakerUsecase) {
	handler := &SpeakerHandler{speakerUsecase: su}
	api := r.Group("/speakers")
	{
		api.GET("/", handler.GetAll)
		api.GET("/:id", handler.GetById)
	}
}

func (h *SpeakerHandler) GetAll(c *gin.Context) {
	query := c.DefaultQuery("name", "")
	speakers, err := h.speakerUsecase.GetAll(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch all data", speakers))
}

func (h *SpeakerHandler) GetById(c *gin.Context) {
	uri := new(domain.UriById)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	speaker, err := h.speakerUsecase.GetById(uri.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", speaker))
}

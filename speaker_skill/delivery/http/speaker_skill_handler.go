package http

import (
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type SpeakerSkillHandler struct {
	speakerSkillUsecase domain.SpeakerSkillUsecase
}

func NewSpeakerSkillHandler(r *gin.RouterGroup, ssu domain.SpeakerSkillUsecase) {
	handler := &SpeakerSkillHandler{speakerSkillUsecase: ssu}
	api := r.Group("/speaker-skills")
	{
		api.GET("/:id", handler.GetBySpeakerId)
	}
}

func (h *SpeakerSkillHandler) GetBySpeakerId(c *gin.Context) {
	uri := new(domain.SkillByIdUri)
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	skill, err := h.speakerSkillUsecase.GetBySpeakerId(uri.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", skill))
}

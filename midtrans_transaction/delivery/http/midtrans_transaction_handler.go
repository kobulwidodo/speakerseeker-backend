package http

import (
	"encoding/json"
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type MidtransTransactionHandler struct {
	midtransTransactionUsecase domain.MidtransTransactionUsecase
}

func NewMidtransTransactionHandler(r *gin.RouterGroup, mtu domain.MidtransTransactionUsecase) {
	handler := &MidtransTransactionHandler{midtransTransactionUsecase: mtu}
	api := r.Group("/midtrans_transaction")
	{
		api.POST("/handler", handler.Handler)
	}
}

func (h *MidtransTransactionHandler) Handler(c *gin.Context) {
	var notifPayload map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&notifPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	orderId, exist := notifPayload["order_id"].(string)
	if !exist {
		c.JSON(http.StatusNotFound, utils.NewFailResponse(err.Error()))
		return
	}
	if err := h.midtransTransactionUsecase.Handler(orderId); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully update data", nil))
}

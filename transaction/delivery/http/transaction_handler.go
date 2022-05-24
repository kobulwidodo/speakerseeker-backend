package http

import (
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase domain.TransactionUseCase
}

func NewTransactionHandler(r *gin.RouterGroup, transactionUsecase domain.TransactionUseCase, jwtMiddleware gin.HandlerFunc) {
	handler := &TransactionHandler{transactionUsecase: transactionUsecase}
	api := r.Group("/transaction")
	{
		api.POST("/:speaker_id", jwtMiddleware, handler.Order)
		api.GET("/:id", jwtMiddleware, handler.FindOne)
		api.GET("/", jwtMiddleware, handler.FindByUserId)
	}
}

func (h *TransactionHandler) Order(c *gin.Context) {
	var input domain.CreateTransaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	var uri domain.TransactionSpeakerIdUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("userLoggedin").(uint)
	id, err := h.transactionUsecase.Order(uri.SpeakerId, userId, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("successfully created order", gin.H{"id": id}))
}

func (h *TransactionHandler) FindOne(c *gin.Context) {
	var uri domain.TransactionIdUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	userId := c.MustGet("userLoggedin").(uint)
	transaction, err := h.transactionUsecase.FindOne(uri.Id, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", transaction))
}

func (h *TransactionHandler) FindByUserId(c *gin.Context) {
	userId := c.MustGet("userLoggedin").(uint)
	transactions, err := h.transactionUsecase.FindByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", transactions))
}

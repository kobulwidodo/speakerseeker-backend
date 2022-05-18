package http

import (
	"net/http"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.RouterGroup, uu domain.UserUseCase, jwtMiddleware gin.HandlerFunc) {
	handler := &UserHandler{UserUseCase: uu}
	api := r.Group("/auth")
	{
		api.POST("/signup", handler.SignUp)
		api.POST("/signin", handler.SignIn)
	}
	r.GET("/me", jwtMiddleware, handler.GetProfile)
}

func (u *UserHandler) SignUp(c *gin.Context) {
	input := new(domain.UserSignUp)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	if err := u.UserUseCase.SignUp(input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.NewSuccessResponse("user successfully registered", nil))
}

func (u *UserHandler) SignIn(c *gin.Context) {
	input := new(domain.UserSignIn)
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewFailResponse(err.Error()))
		return
	}
	token, err := u.UserUseCase.SignIn(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully login", map[string]string{"token": token}))
}

func (u *UserHandler) GetProfile(c *gin.Context) {
	userId := c.MustGet("userLoggedin").(uint)
	user, err := u.UserUseCase.GetMe(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewFailResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.NewSuccessResponse("successfully fetch data", user))
}

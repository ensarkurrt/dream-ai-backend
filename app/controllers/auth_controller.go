package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dto"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/services"
	"net/http"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	userService services.UserService
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constants.InvalidRequest)
	}

	userDto := controller.userService.Login(request)

	c.JSON(http.StatusOK, pkg.BuildResponse[dto.UserDto](constants.Success, userDto))
}

func (controller *AuthControllerImpl) Register(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constants.InvalidRequest)
	}

	userDto := controller.userService.Register(request)

	c.JSON(http.StatusOK, pkg.BuildResponse[dto.UserDto](constants.Success, userDto))
}

func NewAuthController(userService services.UserService) *AuthControllerImpl {
	return &AuthControllerImpl{
		userService,
	}
}

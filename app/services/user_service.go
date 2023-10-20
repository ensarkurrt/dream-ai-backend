package services

import (
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"
	"github.com/yazilimcigenclik/dream-ai-backend/app/domain/dto"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/repository"
	"github.com/yazilimcigenclik/dream-ai-backend/app/utils"
)

type UserService interface {
	Login(request dto.UserLoginRequest) dto.UserDto
	Register(request dto.UserRegisterRequest) dto.UserDto
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (service *UserServiceImpl) Login(request dto.UserLoginRequest) dto.UserDto {
	user, err := service.userRepository.FindByUsername(request.Username)
	if err != nil {
		pkg.PanicException(constants.Unauthorized)
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		pkg.PanicException(constants.Unauthorized)
	}

	accessToken, err := utils.GenerateJWTToken(user.Username, utils.AccessToken)
	if err != nil {
		pkg.PanicException(constants.Unauthorized)
	}

	refreshToken, err := utils.GenerateJWTToken(user.Username, utils.RefreshToken)
	if err != nil {
		pkg.PanicException(constants.Unauthorized)
	}

	var userDto dto.UserDto
	userDto.FromUser(user, accessToken, refreshToken)

	return userDto
}

func (service *UserServiceImpl) Register(request dto.UserRegisterRequest) dto.UserDto {
	user, err := service.userRepository.FindByUsername(request.Username)

	if err == nil || user.ID != 0 {
		pkg.PanicException(constants.InvalidRequest)
	}

	user = dao.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err = service.userRepository.Create(user)

	if err != nil {
		pkg.PanicException(constants.UnknownError)
	}

	accessToken, err := utils.GenerateJWTToken(user.Username, utils.AccessToken)
	if err != nil {
		pkg.PanicException(constants.Unauthorized)
	}

	refreshToken, err := utils.GenerateJWTToken(user.Username, utils.RefreshToken)
	if err != nil {
		pkg.PanicException(constants.Unauthorized)
	}

	var userDto dto.UserDto
	userDto.FromUser(user, accessToken, refreshToken)

	return userDto
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository,
	}
}

package dto

import "github.com/yazilimcigenclik/dream-ai-backend/app/domain/dao"

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDto struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (dto *UserDto) FromUser(user dao.User, accessToken string, refreshToken string) {
	dto.Username = user.Username
	dto.AccessToken = accessToken
	dto.RefreshToken = refreshToken
}

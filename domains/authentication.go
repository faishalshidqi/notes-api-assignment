package domains

import (
	"context"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthenticationUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	ValidateToken(token, secret string) (string, error)
	CheckPasswordHash(password, hash string) error
}

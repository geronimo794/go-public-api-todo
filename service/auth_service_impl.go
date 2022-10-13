package service

import (
	"context"
	"time"

	config "github.com/geronimo794/go-public-api-todo/config"
	helper "github.com/geronimo794/go-public-api-todo/helper"
	"github.com/geronimo794/go-public-api-todo/model/claim"
	"github.com/geronimo794/go-public-api-todo/model/web"

	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}
func (service *AuthServiceImpl) GenerateToken(ctx context.Context, request web.RequestAuth) web.ResponseToken {

	// Create claim or data for the JWT
	claims := &claim.JwtCustomClaims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Add claim data with JWT Header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Add sign the to the JWT Token
	t, err := token.SignedString([]byte(config.JWTKEY))
	helper.PanicIfError(err)

	return web.ResponseToken{
		Token: t,
	}
}
func (service *AuthServiceImpl) Login(ctx context.Context, request web.RequestAuth) bool {
	if request.Username == "admin" && request.Password == "admin" {
		return true
	}
	return false
}

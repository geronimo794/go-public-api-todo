package service

import (
	"context"

	"github.com/geronimo794/go-public-api-todo/model/web"
)

type AuthService interface {
	GenerateToken(ctx context.Context, request web.RequestAuth) web.ResponseToken
	Login(ctx context.Context, request web.RequestAuth) bool
}

package service

import (
	"context"
	"users/internal/model"
)

type Service interface {
	UserRegister(ctx context.Context, email string, password string, username string, bio string, avatar string) (*model.User, error)
	UserLogin(ctx context.Context, email string, password string) (*model.User, error)
	UserInfo(ctx context.Context, userId int64) (*model.User, error)
}

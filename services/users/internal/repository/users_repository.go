package repository

import (
	"context"
	"users/internal/model"
)

type Repository interface {
	RepoLogin(ctx context.Context, id int64) (*model.User, error)
	RepoCreate(ctx context.Context, email string, password string, username string, bio string, avatar string) (*model.User, error)
	RepoUserInfo(ctx context.Context, id int64) (*model.User, error)
}

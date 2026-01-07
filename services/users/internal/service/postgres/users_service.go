package postgres

import (
	"context"
	"errors"
	"users/internal/model"
	"users/internal/repository"
	service2 "users/internal/service"
)

var ErrInvalidInput = errors.New("invalid input")
var ErrPasswordTooShort = errors.New("password too short")

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) service2.Service {
	return &service{repo: repo}
}

func (s *service) UserRegister(ctx context.Context, email string, password string, username string, bio string, avatar string) (*model.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, ErrPasswordTooShort
	}
	if len(password) < 8 {
		return nil, ErrPasswordTooShort
	}
	return s.repo.RepoCreate(ctx, email, password, username, bio, avatar)
}
func (s *service) UserLogin(ctx context.Context, email string, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, ErrInvalidInput
	}
	return s.repo.RepoLogin(ctx, email, password)
}
func (s *service) UserInfo(ctx context.Context, id int64) (*model.User, error) {
	if id == 0 {
		return nil, ErrInvalidInput
	}
	return s.repo.RepoUserInfo(ctx, id)
}

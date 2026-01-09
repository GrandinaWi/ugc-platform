package postgres

import (
	"context"
	"errors"
	"posts/internal/model"
	"posts/internal/repository"
)

var ErrInvalidInput = errors.New("invalid input")

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) PostCreate(ctx context.Context, author_id int64, title, content string) (int64, error) {
	if author_id <= 0 || title == "" || content == "" {
		return 0, ErrInvalidInput
	}
	return svc.repo.RepoCreate(ctx, author_id, title, content)
}
func (svc *Service) PostDelete(ctx context.Context, post_id int64) error {
	if post_id <= 0 {
		return ErrInvalidInput
	}
	return svc.repo.RepoDelete(ctx, post_id)
}
func (svc *Service) PostGet(ctx context.Context, post_id int64) (*model.Post, error) {
	if post_id <= 0 {
		return nil, ErrInvalidInput
	}
	return svc.repo.RepoGetPost(ctx, post_id)
}
func (svc *Service) PostsGet(ctx context.Context) ([]model.Post, error) {
	return svc.repo.RepoGetPosts(ctx)
}

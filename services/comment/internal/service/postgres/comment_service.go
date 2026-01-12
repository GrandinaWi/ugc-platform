package postgres

import (
	"comment/internal/model"
	"comment/internal/repository"
	"context"
	"errors"
)

var ErrInvalidInput = errors.New("invalid input")
var ErrForbidden = errors.New("forbidden")

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

// 1. проверить, что post существует
// 2. если parent_id != nil →
//   - комментарий существует
//   - parent.post_id == comment.post_id
func (svc *Service) Create(ctx context.Context, comment *model.Comment) (int64, error) {
	if comment.UserID == 0 || comment.PostID == 0 || comment.Content == "" {
		return 0, ErrInvalidInput
	}
	return svc.repo.Create(ctx, comment)
}
func (svc *Service) GetByID(ctx context.Context, id int64) (*model.Comment, error) {
	if id == 0 {
		return nil, ErrInvalidInput
	}
	return svc.repo.GetByID(ctx, id)
}
func (svc *Service) GetByAll(ctx context.Context, postID int64) ([]model.Comment, error) {
	if postID == 0 {
		return nil, ErrInvalidInput
	}
	return svc.repo.GetByAll(ctx, postID)
}
func (svc *Service) DeleteByID(ctx context.Context, id int64, userID int64) error {
	if id == 0 {
		return ErrInvalidInput
	}
	comment, err := svc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return ErrForbidden
	}
	return svc.repo.DeleteByID(ctx, id)
}

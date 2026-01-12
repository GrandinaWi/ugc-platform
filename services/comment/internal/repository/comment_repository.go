package repository

import (
	"comment/internal/model"
	"context"
)

type Repository interface {
	Create(ctx context.Context, comment *model.Comment) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.Comment, error)
	DeleteByID(ctx context.Context, id int64, userId int64) error
	GetByAll(ctx context.Context, postId int64) ([]model.Comment, error)
}

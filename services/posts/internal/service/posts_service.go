package service

import (
	"context"
	"posts/internal/model"
)

type Service interface {
	PostCreate(ctx context.Context, author_id int64, title, content string) (int64, error)
	PostDelete(ctx context.Context, post_id int64) error
	PostGet(ctx context.Context, post_id int64) (*model.Post, error)
	PostsGet(ctx context.Context) ([]model.Post, error)
}

package repository

import (
	"context"
	"posts/internal/model"
)

type Repository interface {
	RepoCreate(ctx context.Context, author_id int64, title, content string) (int64, error)
	RepoGetPosts(ctx context.Context) ([]model.Post, error)
	RepoGetPost(ctx context.Context, id int64) (*model.Post, error)
	RepoDelete(ctx context.Context, id int64) error
}

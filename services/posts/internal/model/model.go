package model

import (
	"errors"
	"time"
)

type Post struct {
	ID           string    `json:"id"`
	AuthorID     string    `json:"author_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"created_at"`
}

var PostNotExist = errors.New("post does not exist")

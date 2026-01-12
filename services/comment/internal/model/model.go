package model

import (
	"errors"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	ParentID  *int64    `json:"parent_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var CommentNotExist = errors.New("Comment Not Exist")

package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	Username  string
	Bio       string
	Avatar    string
	CreatedAt time.Time
}

var ErrUserAlreadyExists = errors.New("user already exists")

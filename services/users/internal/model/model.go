package model

import "time"

type UserAuth struct {
	ID        int64     `json:"id"`
	Email     string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

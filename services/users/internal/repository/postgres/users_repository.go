package postgres

import (
	"context"
	"database/sql"
	"strings"
	"users/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) RepoCreate(ctx context.Context, email string, password string, username string, bio string, avatar string) (*model.User, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	var u model.User

	err = tx.QueryRowContext(ctx, "INSERT INTO users_auth (email,password) VALUES ($1,$2) RETURNING id,email", email, string(hash)).Scan(&u.ID, &u.Email)
	if err != nil {
		return nil, err
	}
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, model.ErrUserAlreadyExists
		}
		return nil, err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO users (id,username,bio,avatar_url) VALUES ($1,$2,$3,$4) RETURNING username,bio,avatar_url")
	if err != nil {
		return nil, err
	}
	return &u, nil
}

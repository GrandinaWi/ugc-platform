package postgres

import (
	"comment/internal/model"
	"context"
	"database/sql"
	"errors"
)

type Postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) Create(ctx context.Context, comment *model.Comment) (int64, error) {
	var id int64
	query := `INSERT INTO comments (post_id, author_id, parent_id, content)
		VALUES ($1, $2, $3, $4)
		RETURNING id;`
	err := p.db.QueryRowContext(ctx, query, comment.PostID, comment.UserID, comment.ParentID, comment.Content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (p *Postgres) DeleteByID(ctx context.Context, id int64) error {
	query := `DELETE FROM comments WHERE id = $1;`
	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return model.CommentNotExist // или ваша кастомная ошибка
	}
	return nil
}
func (p *Postgres) GetByAll(ctx context.Context, postId int64) ([]model.Comment, error) {
	var comments []model.Comment
	query := `SELECT id, post_id, author_id, parent_id, content, created_at FROM comments WHERE post_id = $1 ORDER BY created_at ASC;`
	rows, err := p.db.QueryContext(ctx, query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(
			&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
func (p *Postgres) GetByID(ctx context.Context, id int64) (*model.Comment, error) {
	var comment model.Comment
	query := `SELECT id, post_id, author_id, parent_id, content, created_at FROM comments WHERE id = $1;`
	err := p.db.QueryRowContext(ctx, query, id).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.ParentID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.CommentNotExist
		}
		return nil, err
	}
	return &comment, nil
}

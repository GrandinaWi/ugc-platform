package postgres

import (
	"context"
	"database/sql"
	"posts/internal/model"
)

type Postgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) RepoCreate(ctx context.Context, author_id int64, title, content string) (int64, error) {
	var post_id int64
	query := `INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3) RETURNING id`
	err := p.db.QueryRowContext(ctx, query, author_id, title, content).Scan(&post_id)
	if err != nil {
		return 0, err
	}
	return post_id, nil
}
func (p *Postgres) RepoDelete(ctx context.Context, post_id int64) error {
	query := `DELETE FROM posts WHERE id = $1`
	result, err := p.db.ExecContext(ctx, query, post_id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return model.PostNotExist // или ваша кастомная ошибка
	}
	return nil
}
func (p *Postgres) RepoGetPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	query := `SELECT id, author_id, title, content,like_count,comment_count,created_at FROM posts`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.LikeCount,
			&post.CommentCount,
			&post.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
func (p *Postgres) RepoGetPost(ctx context.Context, post_id int64) (*model.Post, error) {
	var post model.Post
	query := `SELECT id, author_id, title, content,like_count,comment_count,created_at FROM posts WHERE id = $1`
	err := p.db.QueryRowContext(ctx, query, post_id).Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.LikeCount, &post.CommentCount, &post.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, model.PostNotExist
	}
	if err != nil {
		return nil, err
	}
	return &post, nil
}

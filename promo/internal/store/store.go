package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"promo/internal/domain"
)

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) SaveClient(ctx context.Context, client domain.Client) error {
	_, err := s.pool.Exec(ctx,
		`INSERT INTO clients (id, registered_at) VALUES ($1, $2)`,
		client.ID, client.RegisteredAt,
	)
	return err
}

func (s *Store) SaveView(ctx context.Context, clientID, entityID string, viewedAt time.Time) (string, error) {
	id := uuid.NewString()
	_, err := s.pool.Exec(ctx,
		`INSERT INTO views (id, client_id, entity_id, viewed_at) VALUES ($1, $2, $3, $4)`,
		id, clientID, entityID, viewedAt,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Store) SaveClick(ctx context.Context, clientID, entityID string, clickedAt time.Time) (string, error) {
	id := uuid.NewString()
	_, err := s.pool.Exec(ctx,
		`INSERT INTO clicks (id, client_id, entity_id, clicked_at) VALUES ($1, $2, $3, $4)`,
		id, clientID, entityID, clickedAt,
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Store) SaveComment(ctx context.Context, clientID, entityID, message string, commentedAt time.Time) (string, error) {
	commentID := uuid.NewString()
	_, err := s.pool.Exec(ctx,
		`INSERT INTO comments (comment_id, client_id, entity_id, message, commented_at) VALUES ($1, $2, $3, $4, $5)`,
		commentID, clientID, entityID, message, commentedAt,
	)
	if err != nil {
		return "", err
	}
	return commentID, nil
}

func (s *Store) ListComments(ctx context.Context, entityID string, pg domain.Pagination) ([]domain.Comment, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT comment_id, client_id, entity_id, message, commented_at
		FROM comments
		WHERE entity_id = $1
		ORDER BY commented_at DESC
		LIMIT $2 OFFSET $3`,
		entityID, pg.Limit, pg.Offset,
	)
	if err != nil {
		return nil, fmt.Errorf("query comments: %w", err)
	}
	defer rows.Close()

	var comments []domain.Comment
	for rows.Next() {
		var c domain.Comment
		if err := rows.Scan(&c.CommentID, &c.ClientID, &c.EntityID, &c.Message, &c.CommentedAt); err != nil {
			return nil, fmt.Errorf("scan comment: %w", err)
		}
		comments = append(comments, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate comments: %w", err)
	}
	return comments, nil
}

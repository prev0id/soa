package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int64
	Login     string
	Password  string
	Email     string
	FirstName string
	LastName  string
	BirthDate time.Time
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
	query := `INSERT INTO users (login, password, email, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, user.Login, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*User, error) {
	user := &User{}
	query := `
	SELECT 
		id,
		login,
		password,
		email,
		first_name,
		last_name,
		birth_date,
		phone,
		created_at,
		updated_at 
    FROM users
	WHERE login = $1`
	err := r.db.
		QueryRow(ctx, query, login).
		Scan(&user.ID, &user.Login, &user.Password, &user.Email, &user.FirstName,
			&user.LastName, &user.BirthDate, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *User) error {
	query := `
	UPDATE users
	SET first_name = $1,
		last_name = $2,
		birth_date = $3,
		phone = $4,
		updated_at = $5,
		email = $6
    WHERE id = $7`

	args := []any{user.FirstName, user.LastName, user.BirthDate, user.Phone, user.UpdatedAt, user.Email, user.ID}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	query := `
	SELECT 
		id,
		login,
		password,
		email,
		first_name,
		last_name,
		birth_date,
		phone,
		created_at,
		updated_at 
	FROM users
	WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Login, &user.Password, &user.Email, &user.FirstName,
		&user.LastName, &user.BirthDate, &user.Phone, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

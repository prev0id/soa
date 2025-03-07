package service

import (
	"context"
	"errors"
	"time"

	"user_service/internal/db"
	"user_service/internal/domain"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("bla bla bla")

type UserService struct {
	repo *db.UserRepository
}

func NewUserService(repo *db.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password, login string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()

	user := &db.User{
		Password:  string(hashedPassword),
		Email:     email,
		Login:     login,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, login, password string) (string, error) {
	dbUser, err := s.repo.GetUserByLogin(ctx, login)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", domain.ErrUserNotFound
	}
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
		return "", domain.ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (s *UserService) ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrInvalidCredentials
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uid, ok := claims["user_id"].(float64); ok {
			return int64(uid), nil
		}
	}
	return 0, domain.ErrInvalidCredentials
}

func (s *UserService) UpdateProfile(ctx context.Context, user *domain.User) error {
	dbUser := &db.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
		Phone:     user.Phone,
		UpdatedAt: time.Now(),
	}

	return s.repo.UpdateUser(ctx, dbUser)
}

func (s *UserService) GetProfile(ctx context.Context, userID int64) (*domain.User, error) {
	dbUser, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, domain.ErrUserNotFound
	}
	return convertDBUserToDomain(dbUser), nil
}

func convertDBUserToDomain(u *db.User) *domain.User {
	if u == nil {
		return nil
	}
	return &domain.User{
		ID:        u.ID,
		Login:     u.Login,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		BirthDate: u.BirthDate,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

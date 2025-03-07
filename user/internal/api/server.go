package api

import (
	"context"
	"errors"
	"log/slog"
	"user_service/internal/common"
	"user_service/internal/domain"
	api_desc "user_service/internal/pkg/api"
	"user_service/internal/service"
)

type Server struct {
	svc *service.UserService
}

func NewServer(svc *service.UserService) *Server {
	return &Server{svc: svc}
}

var _ (api_desc.Handler) = (*Server)(nil)

func (s *Server) GetUserProfile(ctx context.Context) (api_desc.GetUserProfileRes, error) {
	userID, ok := common.UserIDFromContext(ctx)
	if !ok {
		return &api_desc.GetUserProfileUnauthorized{}, nil
	}

	user, err := s.svc.GetProfile(ctx, userID)
	if err != nil {
		slog.Error(err.Error())
		return &api_desc.GetUserProfileInternalServerError{}, nil
	}

	return &api_desc.GetUserProfileOK{
		Login:     api_desc.NewOptString(user.Login),
		Email:     api_desc.NewOptString(user.Email),
		FirstName: api_desc.NewOptString(user.FirstName),
		LastName:  api_desc.NewOptString(user.LastName),
		BirthDate: api_desc.NewOptDate(user.BirthDate),
		Phone:     api_desc.NewOptString(user.Phone),
	}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *api_desc.LoginUserReq) (api_desc.LoginUserRes, error) {
	jwt, err := s.svc.Login(ctx, req.GetLogin(), req.GetPassword())
	if errors.Is(err, domain.ErrInvalidCredentials) {
		slog.Error(err.Error())
		return &api_desc.LoginUserUnauthorized{}, nil
	}
	if err != nil {
		slog.Error(err.Error())
		return &api_desc.LoginUserInternalServerError{}, nil
	}

	return &api_desc.LoginUserOK{
		Token: api_desc.NewOptString(jwt),
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, req *api_desc.RegisterUserReq) (api_desc.RegisterUserRes, error) {
	err := s.svc.RegisterUser(ctx, req.GetEmail(), req.GetPassword(), req.GetLogin())
	if err != nil {
		slog.Error(err.Error())
		return &api_desc.RegisterUserInternalServerError{}, nil
	}
	return &api_desc.RegisterUserOK{}, nil
}

func (s *Server) UpdateUserProfile(ctx context.Context, req *api_desc.UpdateUserProfileReq) (api_desc.UpdateUserProfileRes, error) {
	userID, ok := common.UserIDFromContext(ctx)
	if !ok {
		return &api_desc.UpdateUserProfileUnauthorized{}, nil
	}

	user := &domain.User{
		ID:        userID,
		Email:     req.GetEmail().Value,
		FirstName: req.GetFirstName().Value,
		LastName:  req.GetLastName().Value,
		BirthDate: req.GetBirthDate().Value,
		Phone:     req.GetPhone().Value,
	}

	err := s.svc.UpdateProfile(ctx, user)
	if err != nil {
		slog.Error(err.Error())
		return &api_desc.UpdateUserProfileInternalServerError{}, nil
	}

	return &api_desc.UpdateUserProfileOK{}, nil
}

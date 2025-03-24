package api

import (
	"context"
	"log/slog"
	"user_service/internal/common"
	api_desc "user_service/internal/pkg/api"
	"user_service/internal/service"
)

type SecurityHandler struct {
	svc *service.UserService
}

func NewSecurity(svc *service.UserService) *SecurityHandler {
	return &SecurityHandler{svc: svc}
}

var _ api_desc.SecurityHandler = (*SecurityHandler)(nil)

func (h *SecurityHandler) HandleBearerAuth(ctx context.Context, _ api_desc.OperationName, t api_desc.BearerAuth) (context.Context, error) {
	userID, err := h.svc.ValidateToken(t.GetToken())
	if err != nil {
		slog.Error(err.Error())
		return ctx, err
	}

	return common.WithUserID(ctx, userID), nil
}

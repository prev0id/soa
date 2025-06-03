package api

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"promo/internal/domain"
	"promo/internal/kafka"
	"promo/internal/store"
	eventspb "promo/pkg/events"
	loyaltypb "promo/pkg/loyalty"
)

type Service struct {
	*loyaltypb.UnimplementedLoyaltyCoreServer

	store    *store.Store
	producer *kafka.Producer
}

func NewService(store *store.Store, producer *kafka.Producer) *Service {
	return &Service{store: store, producer: producer}
}

func (s *Service) RegisterClient(ctx context.Context, req *loyaltypb.RegisterClientRequest) (*loyaltypb.RegisterClientResponse, error) {
	client := domain.Client{ID: req.ClientId, RegisteredAt: req.RegisteredAt.AsTime()}
	if err := s.store.SaveClient(ctx, client); err != nil {
		return nil, fmt.Errorf("save client: %w", err)
	}
	event := &eventspb.ClientRegistrationEvent{
		ClientId:     req.ClientId,
		RegisteredAt: req.RegisteredAt,
	}
	if err := s.producer.Send(ctx, "client-registrations", event); err != nil {
		return nil, fmt.Errorf("publish registration event: %w", err)
	}
	return &loyaltypb.RegisterClientResponse{Success: true}, nil
}

func (s *Service) ViewPromo(ctx context.Context, req *loyaltypb.ViewPromoRequest) (*loyaltypb.ViewPromoResponse, error) {
	viewID, err := s.store.SaveView(ctx, req.ClientId, req.EntityId, req.ViewedAt.AsTime())
	if err != nil {
		return nil, fmt.Errorf("save view: %w", err)
	}

	event := &eventspb.ViewEvent{
		EventId:  viewID,
		ClientId: req.ClientId,
		EntityId: req.EntityId,
		ViewedAt: req.ViewedAt,
	}

	if err := s.producer.Send(ctx, "views", event); err != nil {
		return nil, fmt.Errorf("publish view event: %w", err)
	}

	return &loyaltypb.ViewPromoResponse{Success: true}, nil
}

func (s *Service) ClickPromo(ctx context.Context, req *loyaltypb.ClickPromoRequest) (*loyaltypb.ClickPromoResponse, error) {
	clickID, err := s.store.SaveClick(ctx, req.ClientId, req.EntityId, req.ClickedAt.AsTime())
	if err != nil {
		return nil, fmt.Errorf("save click: %w", err)
	}

	event := &eventspb.ClickEvent{
		EventId:   clickID,
		ClientId:  req.ClientId,
		EntityId:  req.EntityId,
		ClickedAt: req.ClickedAt,
	}

	if err := s.producer.Send(ctx, "clicks", event); err != nil {
		return nil, fmt.Errorf("publish click event: %w", err)
	}
	return &loyaltypb.ClickPromoResponse{Success: true}, nil
}

func (s *Service) CommentPromo(ctx context.Context, req *loyaltypb.CommentPromoRequest) (*loyaltypb.CommentPromoResponse, error) {
	commentID, err := s.store.SaveComment(ctx, req.ClientId, req.EntityId, req.Message, req.CommentedAt.AsTime())
	if err != nil {
		return nil, fmt.Errorf("save comment: %w", err)
	}

	event := &eventspb.CommentEvent{
		CommentId:   commentID,
		ClientId:    req.ClientId,
		EntityId:    req.EntityId,
		Message:     req.Message,
		CommentedAt: req.CommentedAt,
	}

	if err := s.producer.Send(ctx, "comments", event); err != nil {
		return nil, fmt.Errorf("publish comment event: %w", err)
	}
	return &loyaltypb.CommentPromoResponse{CommentId: commentID}, nil
}

func (s *Service) ListComments(ctx context.Context, req *loyaltypb.ListCommentsRequest) (*loyaltypb.ListCommentsResponse, error) {
	pg := domain.Pagination{Limit: int(req.PageSize), Offset: 0}
	comments, err := s.store.ListComments(ctx, req.EntityId, pg)
	if err != nil {
		return nil, fmt.Errorf("list comments: %w", err)
	}

	resp := &loyaltypb.ListCommentsResponse{}
	for _, c := range comments {
		resp.Comments = append(resp.Comments, &loyaltypb.Comment{
			CommentId:   c.CommentID,
			ClientId:    c.ClientID,
			Message:     c.Message,
			CommentedAt: timestamppb.New(c.CommentedAt),
		})
	}

	return resp, nil
}

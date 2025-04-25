package main

import (
	"context"
	"log"
	"net/http"

	api_desc "api_service/pkg/api"

	loyaltypb "api_service/pkg/loyalty"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	client loyaltypb.LoyaltyCoreClient
}

var _ api_desc.Handler = (*Service)(nil)

func (s *Service) ClickPromo(ctx context.Context, req *api_desc.ClickPromoReq) (api_desc.ClickPromoRes, error) {
	request := &loyaltypb.ClickPromoRequest{
		ClientId:  req.GetClientID(),
		EntityId:  req.GetEntityID(),
		ClickedAt: timestamppb.Now(),
	}

	resp, err := s.client.ClickPromo(ctx, request)
	if err != nil {
		return &api_desc.ClickPromoInternalServerError{}, nil
	}

	return &api_desc.ClickPromoOK{
		Success:   api_desc.NewOptBool(resp.GetSuccess()),
		ClickedAt: api_desc.NewOptDateTime(request.GetClickedAt().AsTime()),
	}, nil
}

func (s *Service) CommentPromo(ctx context.Context, req *api_desc.CommentPromoReq) (api_desc.CommentPromoRes, error) {
	request := &loyaltypb.CommentPromoRequest{
		ClientId:    req.GetClientID(),
		EntityId:    req.GetEntityID(),
		Message:     req.GetMessage(),
		CommentedAt: timestamppb.Now(),
	}

	resp, err := s.client.CommentPromo(ctx, request)
	if err != nil {
		return &api_desc.CommentPromoInternalServerError{}, nil
	}

	return &api_desc.CommentPromoOK{
		CommentID:   api_desc.NewOptString(resp.GetCommentId()),
		CommentedAt: api_desc.NewOptDateTime(request.GetCommentedAt().AsTime()),
	}, nil
}

func (s *Service) ListComments(ctx context.Context, params api_desc.ListCommentsParams) (api_desc.ListCommentsRes, error) {
	request := &loyaltypb.ListCommentsRequest{
		EntityId:  params.EntityID,
		PageSize:  params.PageSize.Value,
		PageToken: params.PageToken.Value,
	}

	resp, err := s.client.ListComments(ctx, request)
	if err != nil {
		return &api_desc.ListCommentsInternalServerError{}, nil
	}

	comments := make([]api_desc.ListCommentsOKCommentsItem, 0, len(resp.GetComments()))
	for _, comment := range resp.GetComments() {
		comments = append(comments, api_desc.ListCommentsOKCommentsItem{
			CommentID:   api_desc.NewOptString(comment.GetCommentId()),
			ClientID:    api_desc.NewOptString(comment.GetClientId()),
			Message:     api_desc.NewOptString(comment.GetMessage()),
			CommentedAt: api_desc.NewOptDateTime(comment.GetCommentedAt().AsTime()),
		})
	}

	return &api_desc.ListCommentsOK{
		Comments:      comments,
		NextPageToken: api_desc.NewOptString(resp.GetNextPageToken()),
	}, nil
}

func (s *Service) RegisterClient(ctx context.Context, req *api_desc.RegisterClientReq) (api_desc.RegisterClientRes, error) {
	request := &loyaltypb.RegisterClientRequest{
		ClientId:     req.GetClientID(),
		RegisteredAt: timestamppb.Now(),
	}

	resp, err := s.client.RegisterClient(ctx, request)
	if err != nil {
		return &api_desc.RegisterClientInternalServerError{}, nil
	}

	return &api_desc.RegisterClientOK{
		Success:      api_desc.NewOptBool(resp.GetSuccess()),
		RegisteredAt: api_desc.NewOptDateTime(request.GetRegisteredAt().AsTime()),
	}, nil
}

func (s *Service) ViewPromo(ctx context.Context, req *api_desc.ViewPromoReq) (api_desc.ViewPromoRes, error) {
	request := &loyaltypb.ViewPromoRequest{
		ClientId: req.GetClientID(),
		EntityId: req.GetEntityID(),
		ViewedAt: timestamppb.Now(),
	}

	resp, err := s.client.ViewPromo(ctx, request)
	if err != nil {
		return &api_desc.ViewPromoInternalServerError{}, nil
	}

	return &api_desc.ViewPromoOK{
		Success:  api_desc.NewOptBool(resp.GetSuccess()),
		ViewedAt: api_desc.NewOptDateTime(request.GetViewedAt().AsTime()),
	}, nil
}

func main() {
	conn, err := grpc.NewClient("localhost:12301", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	loyaltyClient := loyaltypb.NewLoyaltyCoreClient(conn)

	service := &Service{
		client: loyaltyClient,
	}

	srv, err := api_desc.NewServer(service)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":12300", srv); err != nil {
		log.Fatal(err.Error())
	}
}

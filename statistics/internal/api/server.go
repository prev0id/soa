package api

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"statistics/internal/domain"
	statspb "statistics/pkg/stats"
)

type StatisticsService struct {
	db *sql.DB
}

func NewStatisticsService(db *sql.DB) *StatisticsService {
	return &StatisticsService{db: db}
}

func (s *StatisticsService) GetPromoStats(ctx context.Context, req *statspb.GetPromoStatsRequest) (*statspb.GetPromoStatsResponse, error) {
	var stats domain.PromoStats
	err := s.db.QueryRowContext(ctx, "SELECT sum(views), sum(likes), sum(comments) FROM promo_stats WHERE promo_id = ?", req.PromoId).Scan(&stats.Views, &stats.Likes, &stats.Comments)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get promo stats: %v", err)
	}
	return &statspb.GetPromoStatsResponse{
		Views:    stats.Views,
		Likes:    stats.Likes,
		Comments: stats.Comments,
	}, nil
}

func (s *StatisticsService) GetPromoDynamics(ctx context.Context, req *statspb.GetPromoDynamicsRequest) (*statspb.GetPromoDynamicsResponse, error) {
	metric := map[statspb.Metric]string{
		statspb.Metric_VIEWS:    "views",
		statspb.Metric_LIKES:    "likes",
		statspb.Metric_COMMENTS: "comments",
	}[req.Metric]
	if metric == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid metric")
	}

	startDate := req.StartDate.AsTime().Format("2006-01-02")
	endDate := req.EndDate.AsTime().Format("2006-01-02")

	rows, err := s.db.QueryContext(ctx, "SELECT date, sum("+metric+") FROM promo_stats WHERE promo_id = ? AND date >= ? AND date <= ? GROUP BY date ORDER BY date", req.PromoId, startDate, endDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get promo dynamics: %v", err)
	}
	defer rows.Close()

	var entries []*statspb.DynamicsEntry
	for rows.Next() {
		var dateStr string
		var count uint64
		if err := rows.Scan(&dateStr, &count); err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to scan row: %v", err)
		}
		date, _ := time.Parse("2006-01-02", dateStr)
		entries = append(entries, &statspb.DynamicsEntry{
			Date:  timestamppb.New(date),
			Count: count,
		})
	}
	return &statspb.GetPromoDynamicsResponse{Entries: entries}, nil
}

package growing_zone_history_service

import (
	"context"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
)

func (s *GrowingZoneHistoryService) GetHistoryByDateRange(ctx context.Context, req *growingZoneHistoryP.GetHistoryByDateRangeRequest) (*growingZoneHistoryP.GetHistoryByDateRangeResponse, error) {
	dateRangeReq := s.createGetHistoryByDateRangeReq(req)
	histories, err := s.growingZoneHistoryUsecase.GetHistoryByDateRange.Execute(ctx, dateRangeReq)
	if err != nil {
		return nil, err
	}
	return &growingZoneHistoryP.GetHistoryByDateRangeResponse{
		Success:   true,
		Message:   "History records retrieved successfully",
		Histories: s.createProtoHistories(histories),
	}, nil
}

func (s *GrowingZoneHistoryService) createGetHistoryByDateRangeReq(req *growingZoneHistoryP.GetHistoryByDateRangeRequest) *growing_zone_history.GetHistoryByDateRangeRequest {
	return &growing_zone_history.GetHistoryByDateRangeRequest{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
}

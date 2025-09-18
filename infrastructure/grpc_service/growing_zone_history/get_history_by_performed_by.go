package growing_zone_history_service

import (
	"context"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
)

func (s *GrowingZoneHistoryService) GetHistoryByPerformedBy(ctx context.Context, req *growingZoneHistoryP.GetHistoryByPerformedByRequest) (*growingZoneHistoryP.GetHistoryByPerformedByResponse, error) {
	performedBy := s.createGetHistoryByPerformedByReq(req)
	histories, err := s.growingZoneHistoryUsecase.GetHistoryByPerformedBy.Execute(ctx, performedBy)
	if err != nil {
		return nil, err
	}
	return &growingZoneHistoryP.GetHistoryByPerformedByResponse{
		Success:   true,
		Message:   "History records retrieved successfully",
		Histories: s.createProtoHistories(histories),
	}, nil
}

func (s *GrowingZoneHistoryService) createGetHistoryByPerformedByReq(req *growingZoneHistoryP.GetHistoryByPerformedByRequest) *growing_zone_history.GetHistoryByPerformedByRequest {
	return &growing_zone_history.GetHistoryByPerformedByRequest{
		PerformedBy: req.PerformedBy,
	}
}

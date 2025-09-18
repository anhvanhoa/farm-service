package growing_zone_history_service

import (
	"context"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
)

func (s *GrowingZoneHistoryService) GetHistoryByZone(ctx context.Context, req *growingZoneHistoryP.GetHistoryByZoneRequest) (*growingZoneHistoryP.GetHistoryByZoneResponse, error) {
	zone := s.createGetHistoryByZoneReq(req)
	histories, err := s.growingZoneHistoryUsecase.GetHistoryByZone.Execute(ctx, zone)
	if err != nil {
		return nil, err
	}
	return &growingZoneHistoryP.GetHistoryByZoneResponse{
		Success:   true,
		Message:   "History records retrieved successfully",
		Histories: s.createProtoHistories(histories),
	}, nil
}

func (s *GrowingZoneHistoryService) createGetHistoryByZoneReq(req *growingZoneHistoryP.GetHistoryByZoneRequest) *growing_zone_history.GetHistoryByZoneRequest {
	return &growing_zone_history.GetHistoryByZoneRequest{
		ZoneID: req.ZoneId,
	}
}

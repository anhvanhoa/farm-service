package growing_zone_history_service

import (
	"context"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
)

func (s *GrowingZoneHistoryService) GetHistoryByAction(ctx context.Context, req *growingZoneHistoryP.GetHistoryByActionRequest) (*growingZoneHistoryP.GetHistoryByActionResponse, error) {
	action := s.createGetHistoryByActionReq(req)
	histories, err := s.growingZoneHistoryUsecase.GetHistoryByAction.Execute(ctx, action)
	if err != nil {
		return nil, err
	}
	return &growingZoneHistoryP.GetHistoryByActionResponse{
		Success:   true,
		Message:   "History records retrieved successfully",
		Histories: s.createProtoHistories(histories),
	}, nil
}

func (s *GrowingZoneHistoryService) createGetHistoryByActionReq(req *growingZoneHistoryP.GetHistoryByActionRequest) *growing_zone_history.GetHistoryByActionRequest {
	return &growing_zone_history.GetHistoryByActionRequest{
		Action: req.Action.String(),
	}
}

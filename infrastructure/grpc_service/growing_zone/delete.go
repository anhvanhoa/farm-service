package growingzone_service

import (
	"context"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) DeleteGrowingZone(ctx context.Context, req *growing_zone.DeleteGrowingZoneRequest) (*growing_zone.DeleteGrowingZoneResponse, error) {
	err := s.growingZoneUsecase.DeleteGrowingZone.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &growing_zone.DeleteGrowingZoneResponse{
		Success: true,
		Message: "Growing zone deleted successfully",
	}, nil
}

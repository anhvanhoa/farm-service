package growingzone_service

import (
	"context"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) GetGrowingZone(ctx context.Context, req *growing_zone.GetGrowingZoneRequest) (*growing_zone.GetGrowingZoneResponse, error) {
	growingZone, err := s.growingZoneUsecase.GetGrowingZone.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &growing_zone.GetGrowingZoneResponse{
		Success:     true,
		Message:     "Growing zone retrieved successfully",
		GrowingZone: s.createProtoGrowingZone(growingZone),
	}, nil
}

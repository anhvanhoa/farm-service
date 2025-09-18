package growingzone_service

import (
	"context"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) CreateGrowingZone(ctx context.Context, req *growing_zone.CreateGrowingZoneRequest) (*growing_zone.CreateGrowingZoneResponse, error) {
	growingZone, err := s.growingZoneUsecase.CreateGrowingZone.Execute(ctx, s.createEntityGrowingZone(req))
	if err != nil {
		return nil, err
	}
	return &growing_zone.CreateGrowingZoneResponse{
		Success:     true,
		Message:     "Growing zone created successfully",
		GrowingZone: s.createProtoGrowingZone(growingZone),
	}, nil
}

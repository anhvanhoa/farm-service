package growingzone_service

import (
	"context"
	"farm-service/domain/entity"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) ListGrowingZones(ctx context.Context, req *growing_zone.ListGrowingZonesRequest) (*growing_zone.ListGrowingZonesResponse, error) {
	growingZones, totalCount, err := s.growingZoneUsecase.ListGrowingZone.Execute(
		ctx,
		s.createEntityGrowingZoneFilter(req.Filter), int(req.Page), int(req.PageSize),
	)
	if err != nil {
		return nil, err
	}
	return &growing_zone.ListGrowingZonesResponse{
		Success:      true,
		Message:      "Growing zones retrieved successfully",
		GrowingZones: s.createProtoGrowingZones(growingZones),
		TotalCount:   totalCount,
		Page:         req.Page,
		PageSize:     req.PageSize,
	}, nil
}

func (s *GrowingZoneService) createEntityGrowingZoneFilter(filter *growing_zone.GrowingZoneFilter) *entity.GrowingZoneFilter {
	return &entity.GrowingZoneFilter{
		GreenhouseID:     filter.GreenhouseId,
		Status:           filter.Status,
		SoilType:         filter.SoilType,
		IrrigationSystem: filter.IrrigationSystem,
	}
}

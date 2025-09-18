package growingzone_service

import (
	"context"
	"farm-service/domain/entity"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) UpdateGrowingZone(ctx context.Context, req *growing_zone.UpdateGrowingZoneRequest) (*growing_zone.UpdateGrowingZoneResponse, error) {
	err := s.growingZoneUsecase.UpdateGrowingZone.Execute(ctx, req.Id, s.createEntityUpdateGrowingZone(req))
	if err != nil {
		return nil, err
	}
	return &growing_zone.UpdateGrowingZoneResponse{
		Success: true,
		Message: "Growing zone updated successfully",
	}, nil
}

func (s *GrowingZoneService) createEntityUpdateGrowingZone(req *growing_zone.UpdateGrowingZoneRequest) *entity.UpdateGrowingZoneRequest {
	return &entity.UpdateGrowingZoneRequest{
		ZoneName:         req.ZoneName,
		ZoneCode:         req.ZoneCode,
		AreaM2:           &req.AreaM2,
		MaxPlants:        &req.MaxPlants,
		SoilType:         req.SoilType,
		IrrigationSystem: req.IrrigationSystem,
		Status:           req.Status,
	}
}

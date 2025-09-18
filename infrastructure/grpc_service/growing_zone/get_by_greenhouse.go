package growingzone_service

import (
	"context"
	"farm-service/domain/entity"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
)

func (s *GrowingZoneService) GetZonesByGreenhouse(
	ctx context.Context,
	req *growing_zone.GetZonesByGreenhouseRequest,
) (*growing_zone.GetZonesByGreenhouseResponse, error) {
	growingZones, err := s.growingZoneUsecase.GetZonesByGreenhouse.Execute(ctx, req.GreenhouseId)
	if err != nil {
		return nil, err
	}
	return &growing_zone.GetZonesByGreenhouseResponse{
		Success:      true,
		Message:      "Growing zones retrieved successfully",
		GrowingZones: s.createProtoGrowingZones(growingZones),
	}, nil
}

func (s *GrowingZoneService) createProtoGrowingZones(growingZones []*entity.GrowingZone) (growingZonesProto []*growing_zone.GrowingZone) {
	for _, growingZone := range growingZones {
		growingZonesProto = append(growingZonesProto, s.createProtoGrowingZone(growingZone))
	}
	return growingZonesProto
}

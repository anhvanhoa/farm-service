package growingzone_service

import (
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	growingzone "farm-service/domain/usecase/growing_zone"
	"time"

	growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrowingZoneService struct {
	growing_zone.UnsafeGrowingZoneServiceServer
	growingZoneUsecase growingzone.GrowingZoneUsecase
}

func NewGrowingZoneService(growingZoneRepository repository.GrowingZoneRepository) growing_zone.GrowingZoneServiceServer {
	growingZoneUsecase := growingzone.NewGrowingZoneUsecase(growingZoneRepository)
	return &GrowingZoneService{
		growingZoneUsecase: *growingZoneUsecase,
	}
}

func (s *GrowingZoneService) createEntityGrowingZone(req *growing_zone.CreateGrowingZoneRequest) *entity.CreateGrowingZoneRequest {
	return &entity.CreateGrowingZoneRequest{
		GreenhouseID:     req.GreenhouseId,
		ZoneName:         req.ZoneName,
		ZoneCode:         req.ZoneCode,
		AreaM2:           req.AreaM2,
		MaxPlants:        req.MaxPlants,
		SoilType:         req.SoilType,
		IrrigationSystem: req.IrrigationSystem,
		CreatedBy:        req.CreatedBy,
	}
}

func (s *GrowingZoneService) createProtoGrowingZone(growingZone *entity.GrowingZone) *growing_zone.GrowingZone {
	response := &growing_zone.GrowingZone{
		Id:               growingZone.ID,
		GreenhouseId:     growingZone.GreenhouseID,
		ZoneName:         growingZone.ZoneName,
		ZoneCode:         growingZone.ZoneCode,
		AreaM2:           growingZone.AreaM2,
		MaxPlants:        growingZone.MaxPlants,
		SoilType:         growingZone.SoilType,
		IrrigationSystem: growingZone.IrrigationSystem,
		Status:           growingZone.Status,
		CreatedBy:        growingZone.CreatedBy,
	}

	if growingZone.CreatedAt != (time.Time{}) {
		response.CreatedAt = timestamppb.New(growingZone.CreatedAt)
	}
	if growingZone.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*growingZone.UpdatedAt)
	}
	return response
}

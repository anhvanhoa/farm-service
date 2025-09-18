package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"time"
)

// CreateGrowingZoneUsecase định nghĩa interface cho use case tạo khu vực trồng
type CreateGrowingZoneUsecase interface {
	Execute(ctx context.Context, req *entity.CreateGrowingZoneRequest) (*entity.GrowingZone, error)
}

type createGrowingZoneUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

// NewCreateGrowingZoneUsecase tạo instance mới của CreateGrowingZoneUsecase
func NewCreateGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) CreateGrowingZoneUsecase {
	return &createGrowingZoneUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *createGrowingZoneUsecase) Execute(ctx context.Context, req *entity.CreateGrowingZoneRequest) (*entity.GrowingZone, error) {
	// Kiểm tra zone code đã tồn tại chưa
	exists, err := u.growingZoneRepository.CheckZoneCodeExists(ctx, req.ZoneCode)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &entity.Error{Code: "ZONE_CODE_EXISTS", Message: "Zone code already exists"}
	}

	// Tạo entity GrowingZone từ request
	now := time.Now()
	zone := &entity.GrowingZone{
		GreenhouseID:     req.GreenhouseID,
		ZoneName:         req.ZoneName,
		ZoneCode:         req.ZoneCode,
		AreaM2:           req.AreaM2,
		MaxPlants:        req.MaxPlants,
		SoilType:         req.SoilType,
		IrrigationSystem: req.IrrigationSystem,
		Status:           entity.StatusActive,
		CreatedBy:        req.CreatedBy,
		CreatedAt:        now,
		UpdatedAt:        &now,
	}

	// Lưu vào database
	err = u.growingZoneRepository.Create(ctx, zone)
	if err != nil {
		return nil, err
	}

	return zone, nil
}

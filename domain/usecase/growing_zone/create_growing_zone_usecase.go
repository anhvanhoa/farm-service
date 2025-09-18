package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrZoneCodeExists = oops.New("Zone code already exists")
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
	exists, err := u.growingZoneRepository.CheckZoneCodeExists(ctx, req.ZoneCode)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrZoneCodeExists
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

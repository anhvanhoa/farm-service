package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetZonesByGreenhouseUsecase định nghĩa interface cho use case lấy tất cả khu vực trồng của một nhà lưới
type GetZonesByGreenhouseUsecase interface {
	Execute(ctx context.Context, greenhouseID string) ([]*entity.GrowingZone, error)
}

type getZonesByGreenhouseUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

// NewGetZonesByGreenhouseUsecase tạo instance mới của GetZonesByGreenhouseUsecase
func NewGetZonesByGreenhouseUsecase(growingZoneRepository repository.GrowingZoneRepository) GetZonesByGreenhouseUsecase {
	return &getZonesByGreenhouseUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *getZonesByGreenhouseUsecase) Execute(ctx context.Context, greenhouseID string) ([]*entity.GrowingZone, error) {
	return u.growingZoneRepository.GetByGreenhouseID(ctx, greenhouseID)
}

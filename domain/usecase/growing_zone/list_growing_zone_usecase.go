package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// ListGrowingZoneUsecase định nghĩa interface cho use case lấy danh sách khu vực trồng
type ListGrowingZoneUsecase interface {
	Execute(ctx context.Context, filter *entity.GrowingZoneFilter, page, pageSize int) ([]*entity.GrowingZone, int64, error)
}

type listGrowingZoneUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

// NewListGrowingZoneUsecase tạo instance mới của ListGrowingZoneUsecase
func NewListGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) ListGrowingZoneUsecase {
	return &listGrowingZoneUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *listGrowingZoneUsecase) Execute(ctx context.Context, filter *entity.GrowingZoneFilter, page, pageSize int) ([]*entity.GrowingZone, int64, error) {
	return u.growingZoneRepository.List(ctx, filter, page, pageSize)
}

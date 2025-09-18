package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/oops"
)

// GetGrowingZoneUsecase định nghĩa interface cho use case lấy thông tin khu vực trồng
type GetGrowingZoneUsecase interface {
	Execute(ctx context.Context, id string) (*entity.GrowingZone, error)
}

type getGrowingZoneUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

var (
	ErrNotFoundGrowingZone = oops.New("Growing zone not found")
)

// NewGetGrowingZoneUsecase tạo instance mới của GetGrowingZoneUsecase
func NewGetGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) GetGrowingZoneUsecase {
	return &getGrowingZoneUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *getGrowingZoneUsecase) Execute(ctx context.Context, id string) (*entity.GrowingZone, error) {
	growingZone, err := u.growingZoneRepository.GetByID(ctx, id)
	if growingZone == nil {
		return nil, ErrNotFoundGrowingZone
	}
	return growingZone, err
}

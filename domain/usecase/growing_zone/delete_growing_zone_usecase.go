package growingzone

import (
	"context"
	"farm-service/domain/repository"
)

// DeleteGrowingZoneUsecase định nghĩa interface cho use case xóa khu vực trồng
type DeleteGrowingZoneUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteGrowingZoneUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

// NewDeleteGrowingZoneUsecase tạo instance mới của DeleteGrowingZoneUsecase
func NewDeleteGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) DeleteGrowingZoneUsecase {
	return &deleteGrowingZoneUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *deleteGrowingZoneUsecase) Execute(ctx context.Context, id string) error {
	// Kiểm tra khu vực trồng có tồn tại không
	_, err := u.growingZoneRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Xóa khu vực trồng
	return u.growingZoneRepository.Delete(ctx, id)
}

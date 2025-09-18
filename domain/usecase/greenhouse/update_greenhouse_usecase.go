package greenhouse

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdateGreenhouseUsecase định nghĩa interface cho use case cập nhật nhà lưới
type UpdateGreenhouseUsecase interface {
	Execute(ctx context.Context, id string, req *entity.UpdateGreenhouseRequest) error
}

type updateGreenhouseUsecase struct {
	greenhouseRepository repository.GreenhouseRepository
}

// NewUpdateGreenhouseUsecase tạo instance mới của UpdateGreenhouseUsecase
func NewUpdateGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) UpdateGreenhouseUsecase {
	return &updateGreenhouseUsecase{
		greenhouseRepository: greenhouseRepository,
	}
}

func (u *updateGreenhouseUsecase) Execute(ctx context.Context, id string, req *entity.UpdateGreenhouseRequest) error {
	// Kiểm tra nhà lưới có tồn tại không
	_, err := u.greenhouseRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Cập nhật thông tin
	return u.greenhouseRepository.Update(ctx, id, req)
}

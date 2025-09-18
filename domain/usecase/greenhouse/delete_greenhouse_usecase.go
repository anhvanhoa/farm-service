package greenhouse

import (
	"context"
	"farm-service/domain/repository"
)

// DeleteGreenhouseUsecase định nghĩa interface cho use case xóa nhà lưới
type DeleteGreenhouseUsecase interface {
	Execute(ctx context.Context, id string) error
}

type deleteGreenhouseUsecase struct {
	greenhouseRepository repository.GreenhouseRepository
}

// NewDeleteGreenhouseUsecase tạo instance mới của DeleteGreenhouseUsecase
func NewDeleteGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) DeleteGreenhouseUsecase {
	return &deleteGreenhouseUsecase{
		greenhouseRepository: greenhouseRepository,
	}
}

func (u *deleteGreenhouseUsecase) Execute(ctx context.Context, id string) error {
	// Kiểm tra nhà lưới có tồn tại không
	_, err := u.greenhouseRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Xóa nhà lưới
	return u.greenhouseRepository.Delete(ctx, id)
}

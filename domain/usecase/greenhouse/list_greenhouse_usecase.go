package greenhouse

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// ListGreenhouseUsecase định nghĩa interface cho use case lấy danh sách nhà lưới
type ListGreenhouseUsecase interface {
	Execute(ctx context.Context, filter *entity.GreenhouseFilter, page, pageSize int) ([]*entity.Greenhouse, int64, error)
}

type listGreenhouseUsecase struct {
	greenhouseRepository repository.GreenhouseRepository
}

// NewListGreenhouseUsecase tạo instance mới của ListGreenhouseUsecase
func NewListGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) ListGreenhouseUsecase {
	return &listGreenhouseUsecase{
		greenhouseRepository: greenhouseRepository,
	}
}

func (u *listGreenhouseUsecase) Execute(ctx context.Context, filter *entity.GreenhouseFilter, page, pageSize int) ([]*entity.Greenhouse, int64, error) {
	return u.greenhouseRepository.List(ctx, filter, page, pageSize)
}

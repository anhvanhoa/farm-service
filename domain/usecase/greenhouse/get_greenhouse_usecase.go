package greenhouse

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/anhvanhoa/service-core/domain/oops"
)

// GetGreenhouseUsecase định nghĩa interface cho use case lấy thông tin nhà lưới
type GetGreenhouseUsecase interface {
	Execute(ctx context.Context, id string) (*entity.Greenhouse, error)
}

type getGreenhouseUsecase struct {
	greenhouseRepository repository.GreenhouseRepository
}

var (
	ErrNotFound = oops.New("Greenhouse not found")
)

// NewGetGreenhouseUsecase tạo instance mới của GetGreenhouseUsecase
func NewGetGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) GetGreenhouseUsecase {
	return &getGreenhouseUsecase{
		greenhouseRepository: greenhouseRepository,
	}
}

func (u *getGreenhouseUsecase) Execute(ctx context.Context, id string) (*entity.Greenhouse, error) {
	res, err := u.greenhouseRepository.GetByID(ctx, id)
	if res == nil {
		return nil, ErrNotFound
	}
	return res, err
}

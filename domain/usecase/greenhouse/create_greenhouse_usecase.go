package greenhouse

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"time"
)

type CreateGreenhouseUsecase interface {
	Execute(ctx context.Context, req *entity.CreateGreenhouseRequest) (*entity.Greenhouse, error)
}

type createGreenhouseUsecase struct {
	greenhouseRepository repository.GreenhouseRepository
}

func NewCreateGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) CreateGreenhouseUsecase {
	return &createGreenhouseUsecase{
		greenhouseRepository: greenhouseRepository,
	}
}

func (u *createGreenhouseUsecase) Execute(ctx context.Context, req *entity.CreateGreenhouseRequest) (*entity.Greenhouse, error) {
	greenhouse := &entity.Greenhouse{
		Name:             req.Name,
		Location:         req.Location,
		AreaM2:           req.AreaM2,
		Type:             req.Type,
		MaxCapacity:      req.MaxCapacity,
		InstallationDate: req.InstallationDate,
		Status:           entity.StatusActive,
		Description:      req.Description,
		CreatedBy:        req.CreatedBy,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := u.greenhouseRepository.Create(ctx, greenhouse)
	if err != nil {
		return nil, err
	}

	return greenhouse, nil
}

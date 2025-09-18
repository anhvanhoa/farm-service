package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"farm-service/domain/usecase/greenhouse"
)

type GetLogsByGreenhouseRequest struct {
	GreenhouseID string `json:"greenhouse_id" binding:"required"`
}

type GetLogsByGreenhouseUsecase interface {
	Execute(ctx context.Context, req *GetLogsByGreenhouseRequest) ([]*entity.GreenhouseInstallationLog, error)
}

type getLogsByGreenhouseUsecase struct {
	greenhouseRepo repository.GreenhouseRepository
	logRepo        repository.GreenhouseInstallationLogRepository
}

func NewGetLogsByGreenhouseUsecase(
	greenhouseRepo repository.GreenhouseRepository,
	logRepo repository.GreenhouseInstallationLogRepository,
) GetLogsByGreenhouseUsecase {
	return &getLogsByGreenhouseUsecase{
		greenhouseRepo: greenhouseRepo,
		logRepo:        logRepo,
	}
}

func (u *getLogsByGreenhouseUsecase) Execute(ctx context.Context, req *GetLogsByGreenhouseRequest) ([]*entity.GreenhouseInstallationLog, error) {
	// Kiểm tra greenhouse có tồn tại không
	gh, err := u.greenhouseRepo.GetByID(ctx, req.GreenhouseID)

	if gh == nil {
		return nil, greenhouse.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	// Lấy logs theo greenhouse ID
	logs, err := u.logRepo.GetByGreenhouseID(ctx, req.GreenhouseID)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

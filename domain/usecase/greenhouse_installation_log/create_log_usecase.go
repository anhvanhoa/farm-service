package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"farm-service/domain/usecase/greenhouse"
	"time"
)

type CreateLogRequest struct {
	GreenhouseID string
	Action       string
	ActionDate   time.Time
	Description  string
	PerformedBy  string
}

type CreateLogUsecase interface {
	Execute(ctx context.Context, req *CreateLogRequest) (*entity.GreenhouseInstallationLog, error)
}

type createLogUsecase struct {
	greenhouseRepo repository.GreenhouseRepository
	logRepo        repository.GreenhouseInstallationLogRepository
}

func NewCreateLogUsecase(
	greenhouseRepo repository.GreenhouseRepository,
	logRepo repository.GreenhouseInstallationLogRepository,
) CreateLogUsecase {
	return &createLogUsecase{
		greenhouseRepo: greenhouseRepo,
		logRepo:        logRepo,
	}
}

func (u *createLogUsecase) Execute(ctx context.Context, req *CreateLogRequest) (*entity.GreenhouseInstallationLog, error) {
	gh, err := u.greenhouseRepo.GetByID(ctx, req.GreenhouseID)

	if gh == nil {
		return nil, greenhouse.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	// Tạo log mới
	log := &entity.GreenhouseInstallationLog{
		GreenhouseID: req.GreenhouseID,
		Action:       req.Action,
		ActionDate:   req.ActionDate,
		Description:  req.Description,
		PerformedBy:  req.PerformedBy,
		CreatedAt:    time.Now(),
	}

	// Lưu vào database
	err = u.logRepo.Create(ctx, log)
	if err != nil {
		return nil, err
	}

	return log, nil
}

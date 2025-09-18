package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

type GetLogsByActionRequest struct {
	Action string `json:"action" binding:"required,oneof=install upgrade maintenance relocate dismantle"`
}

type GetLogsByActionUsecase interface {
	Execute(ctx context.Context, req *GetLogsByActionRequest) ([]*entity.GreenhouseInstallationLog, error)
}

type getLogsByActionUsecase struct {
	logRepo repository.GreenhouseInstallationLogRepository
}

func NewGetLogsByActionUsecase(logRepo repository.GreenhouseInstallationLogRepository) GetLogsByActionUsecase {
	return &getLogsByActionUsecase{
		logRepo: logRepo,
	}
}

func (u *getLogsByActionUsecase) Execute(ctx context.Context, req *GetLogsByActionRequest) ([]*entity.GreenhouseInstallationLog, error) {
	// Lấy logs theo action
	logs, err := u.logRepo.GetByAction(ctx, req.Action)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

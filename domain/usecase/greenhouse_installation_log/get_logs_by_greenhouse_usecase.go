package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
)

type GetLogsByGreenhouseRequest struct {
	GreenhouseID string `json:"greenhouse_id" binding:"required"`
}

type GetLogsByGreenhouseUseCase struct {
	BaseUseCase
}

func NewGetLogsByGreenhouseUseCase(baseUseCase *BaseUseCase) *GetLogsByGreenhouseUseCase {
	return &GetLogsByGreenhouseUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetLogsByGreenhouseUseCase) Execute(ctx context.Context, req *GetLogsByGreenhouseRequest) ([]*entity.GreenhouseInstallationLog, error) {
	// Kiểm tra greenhouse có tồn tại không
	greenhouse, err := u.GreenhouseRepo.GetByID(ctx, req.GreenhouseID)
	if err != nil {
		return nil, err
	}
	if greenhouse == nil {
		return nil, &entity.Error{
			Code:    "GREENHOUSE_NOT_FOUND",
			Message: "Greenhouse not found",
		}
	}

	// Lấy logs theo greenhouse ID
	logs, err := u.LogRepo.GetByGreenhouseID(ctx, req.GreenhouseID)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

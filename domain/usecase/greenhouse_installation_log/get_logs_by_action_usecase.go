package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
)

type GetLogsByActionRequest struct {
	Action string `json:"action" binding:"required,oneof=install upgrade maintenance relocate dismantle"`
}

type GetLogsByActionUseCase struct {
	BaseUseCase
}

func NewGetLogsByActionUseCase(baseUseCase *BaseUseCase) *GetLogsByActionUseCase {
	return &GetLogsByActionUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetLogsByActionUseCase) Execute(ctx context.Context, req *GetLogsByActionRequest) ([]*entity.GreenhouseInstallationLog, error) {
	// Lấy logs theo action
	logs, err := u.LogRepo.GetByAction(ctx, req.Action)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
)

type GetHistoryByActionRequest struct {
	Action string `json:"action" binding:"required,oneof=change_soil change_irrigation maintenance resize rename"`
}

type GetHistoryByActionUseCase struct {
	BaseUseCase
}

func NewGetHistoryByActionUseCase(baseUseCase *BaseUseCase) *GetHistoryByActionUseCase {
	return &GetHistoryByActionUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetHistoryByActionUseCase) Execute(ctx context.Context, req *GetHistoryByActionRequest) ([]*entity.GrowingZoneHistory, error) {
	// Lấy history theo action
	histories, err := u.HistoryRepo.GetByAction(ctx, req.Action)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

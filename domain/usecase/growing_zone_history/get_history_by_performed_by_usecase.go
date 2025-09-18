package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
)

type GetHistoryByPerformedByRequest struct {
	PerformedBy string `json:"performed_by" binding:"required"`
}

type GetHistoryByPerformedByUseCase struct {
	BaseUseCase
}

func NewGetHistoryByPerformedByUseCase(baseUseCase *BaseUseCase) *GetHistoryByPerformedByUseCase {
	return &GetHistoryByPerformedByUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetHistoryByPerformedByUseCase) Execute(ctx context.Context, req *GetHistoryByPerformedByRequest) ([]*entity.GrowingZoneHistory, error) {
	// Lấy history theo người thực hiện
	histories, err := u.HistoryRepo.GetByPerformedBy(ctx, req.PerformedBy)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

type GetHistoryByPerformedByRequest struct {
	PerformedBy string `json:"performed_by" binding:"required"`
}

type GetHistoryByPerformedByUsecase interface {
	Execute(ctx context.Context, req *GetHistoryByPerformedByRequest) ([]*entity.GrowingZoneHistory, error)
}

type getHistoryByPerformedByUsecase struct {
	historyRepo repository.GrowingZoneHistoryRepository
}

func NewGetHistoryByPerformedByUsecase(historyRepo repository.GrowingZoneHistoryRepository) GetHistoryByPerformedByUsecase {
	return &getHistoryByPerformedByUsecase{
		historyRepo: historyRepo,
	}
}

func (u *getHistoryByPerformedByUsecase) Execute(ctx context.Context, req *GetHistoryByPerformedByRequest) ([]*entity.GrowingZoneHistory, error) {
	// Lấy history theo người thực hiện
	histories, err := u.historyRepo.GetByPerformedBy(ctx, req.PerformedBy)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

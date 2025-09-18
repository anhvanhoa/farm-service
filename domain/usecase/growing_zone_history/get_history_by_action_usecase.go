package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

type GetHistoryByActionRequest struct {
	Action string
}

type GetHistoryByActionUsecase interface {
	Execute(ctx context.Context, req *GetHistoryByActionRequest) ([]*entity.GrowingZoneHistory, error)
}

type getHistoryByActionUsecase struct {
	historyRepo repository.GrowingZoneHistoryRepository
}

func NewGetHistoryByActionUsecase(historyRepo repository.GrowingZoneHistoryRepository) GetHistoryByActionUsecase {
	return &getHistoryByActionUsecase{
		historyRepo: historyRepo,
	}
}

func (u *getHistoryByActionUsecase) Execute(ctx context.Context, req *GetHistoryByActionRequest) ([]*entity.GrowingZoneHistory, error) {
	// Lấy history theo action
	histories, err := u.historyRepo.GetByAction(ctx, req.Action)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

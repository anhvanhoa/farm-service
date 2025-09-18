package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	growingzone "farm-service/domain/usecase/growing_zone"
)

type GetHistoryByZoneRequest struct {
	ZoneID string `json:"zone_id" binding:"required"`
}

type GetHistoryByZoneUsecase interface {
	Execute(ctx context.Context, req *GetHistoryByZoneRequest) ([]*entity.GrowingZoneHistory, error)
}

type getHistoryByZoneUsecase struct {
	growingZoneRepo repository.GrowingZoneRepository
	historyRepo     repository.GrowingZoneHistoryRepository
}

func NewGetHistoryByZoneUsecase(
	growingZoneRepo repository.GrowingZoneRepository,
	historyRepo repository.GrowingZoneHistoryRepository,
) GetHistoryByZoneUsecase {
	return &getHistoryByZoneUsecase{
		growingZoneRepo: growingZoneRepo,
		historyRepo:     historyRepo,
	}
}

func (u *getHistoryByZoneUsecase) Execute(ctx context.Context, req *GetHistoryByZoneRequest) ([]*entity.GrowingZoneHistory, error) {
	zone, err := u.growingZoneRepo.GetByID(ctx, req.ZoneID)
	if zone == nil {
		return nil, growingzone.ErrNotFoundGrowingZone
	}
	if err != nil {
		return nil, err
	}

	// Lấy history theo zone ID
	histories, err := u.historyRepo.GetByZoneID(ctx, req.ZoneID)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

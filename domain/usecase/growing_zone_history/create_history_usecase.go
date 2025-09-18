package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	growingzone "farm-service/domain/usecase/growing_zone"
	"time"
)

type CreateHistoryRequest struct {
	ZoneID      string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	PerformedBy string
	Notes       string
}

type CreateHistoryUsecase interface {
	Execute(ctx context.Context, req *CreateHistoryRequest) (*entity.GrowingZoneHistory, error)
}

type createHistoryUsecase struct {
	growingZoneRepo repository.GrowingZoneRepository
	historyRepo     repository.GrowingZoneHistoryRepository
}

func NewCreateHistoryUsecase(
	growingZoneRepo repository.GrowingZoneRepository,
	historyRepo repository.GrowingZoneHistoryRepository,
) CreateHistoryUsecase {
	return &createHistoryUsecase{
		growingZoneRepo: growingZoneRepo,
		historyRepo:     historyRepo,
	}
}

func (u *createHistoryUsecase) Execute(ctx context.Context, req *CreateHistoryRequest) (*entity.GrowingZoneHistory, error) {
	// Kiểm tra growing zone có tồn tại không
	zone, err := u.growingZoneRepo.GetByID(ctx, req.ZoneID)
	if zone == nil {
		return nil, growingzone.ErrNotFoundGrowingZone
	}
	if err != nil {
		return nil, err
	}

	// Tạo history record mới
	history := &entity.GrowingZoneHistory{
		ZoneID:      req.ZoneID,
		Action:      req.Action,
		OldValue:    req.OldValue,
		NewValue:    req.NewValue,
		ActionDate:  time.Now(),
		PerformedBy: req.PerformedBy,
		Notes:       req.Notes,
	}

	// Lưu vào database
	err = u.historyRepo.Create(ctx, history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

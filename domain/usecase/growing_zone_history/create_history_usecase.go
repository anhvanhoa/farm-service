package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"time"

	"github.com/google/uuid"
)

type CreateHistoryRequest struct {
	ZoneID      string                 `json:"zone_id" binding:"required"`
	Action      string                 `json:"action" binding:"required,oneof=change_soil change_irrigation maintenance resize rename"`
	OldValue    map[string]interface{} `json:"old_value"`
	NewValue    map[string]interface{} `json:"new_value"`
	PerformedBy string                 `json:"performed_by" binding:"required"`
	Notes       string                 `json:"notes"`
}

type CreateHistoryUseCase struct {
	BaseUseCase
}

func NewCreateHistoryUseCase(baseUseCase *BaseUseCase) *CreateHistoryUseCase {
	return &CreateHistoryUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *CreateHistoryUseCase) Execute(ctx context.Context, req *CreateHistoryRequest) (*entity.GrowingZoneHistory, error) {
	// Kiểm tra growing zone có tồn tại không
	zone, err := u.GrowingZoneRepo.GetByID(ctx, req.ZoneID)
	if err != nil {
		return nil, err
	}
	if zone == nil {
		return nil, &entity.Error{
			Code:    "GROWING_ZONE_NOT_FOUND",
			Message: "Growing zone not found",
		}
	}

	// Tạo history record mới
	history := &entity.GrowingZoneHistory{
		ID:          uuid.New().String(),
		ZoneID:      req.ZoneID,
		Action:      req.Action,
		OldValue:    req.OldValue,
		NewValue:    req.NewValue,
		ActionDate:  time.Now(),
		PerformedBy: req.PerformedBy,
		Notes:       req.Notes,
	}

	// Lưu vào database
	err = u.HistoryRepo.Create(ctx, history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

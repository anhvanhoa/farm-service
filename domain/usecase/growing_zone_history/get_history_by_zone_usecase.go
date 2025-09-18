package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
)

type GetHistoryByZoneRequest struct {
	ZoneID string `json:"zone_id" binding:"required"`
}

type GetHistoryByZoneUseCase struct {
	BaseUseCase
}

func NewGetHistoryByZoneUseCase(baseUseCase *BaseUseCase) *GetHistoryByZoneUseCase {
	return &GetHistoryByZoneUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetHistoryByZoneUseCase) Execute(ctx context.Context, req *GetHistoryByZoneRequest) ([]*entity.GrowingZoneHistory, error) {
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

	// Lấy history theo zone ID
	histories, err := u.HistoryRepo.GetByZoneID(ctx, req.ZoneID)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

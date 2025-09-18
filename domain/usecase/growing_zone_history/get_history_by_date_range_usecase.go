package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"time"
)

type GetHistoryByDateRangeRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type GetHistoryByDateRangeUseCase struct {
	BaseUseCase
}

func NewGetHistoryByDateRangeUseCase(baseUseCase *BaseUseCase) *GetHistoryByDateRangeUseCase {
	return &GetHistoryByDateRangeUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetHistoryByDateRangeUseCase) Execute(ctx context.Context, req *GetHistoryByDateRangeRequest) ([]*entity.GrowingZoneHistory, error) {
	// Validate date format
	_, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, &entity.Error{
			Code:    "INVALID_START_DATE_FORMAT",
			Message: "Invalid start date format. Expected YYYY-MM-DD",
		}
	}

	_, err = time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, &entity.Error{
			Code:    "INVALID_END_DATE_FORMAT",
			Message: "Invalid end date format. Expected YYYY-MM-DD",
		}
	}

	// Lấy history theo khoảng thời gian
	histories, err := u.HistoryRepo.GetByDateRange(ctx, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

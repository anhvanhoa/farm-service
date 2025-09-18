package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
	"time"
)

type GetLogsByDateRangeRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type GetLogsByDateRangeUseCase struct {
	BaseUseCase
}

func NewGetLogsByDateRangeUseCase(baseUseCase *BaseUseCase) *GetLogsByDateRangeUseCase {
	return &GetLogsByDateRangeUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *GetLogsByDateRangeUseCase) Execute(ctx context.Context, req *GetLogsByDateRangeRequest) ([]*entity.GreenhouseInstallationLog, error) {
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

	// Lấy logs theo khoảng thời gian
	logs, err := u.LogRepo.GetByDateRange(ctx, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

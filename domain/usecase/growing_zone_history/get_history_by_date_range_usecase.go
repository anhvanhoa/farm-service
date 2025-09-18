package growing_zone_history

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"time"

	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrInvalidStartDateFormat = oops.New("invalid start date format")
	ErrInvalidEndDateFormat   = oops.New("invalid end date format")
)

type GetHistoryByDateRangeRequest struct {
	StartDate string
	EndDate   string
}

type GetHistoryByDateRangeUsecase interface {
	Execute(ctx context.Context, req *GetHistoryByDateRangeRequest) ([]*entity.GrowingZoneHistory, error)
}

type getHistoryByDateRangeUsecase struct {
	historyRepo repository.GrowingZoneHistoryRepository
}

func NewGetHistoryByDateRangeUsecase(historyRepo repository.GrowingZoneHistoryRepository) GetHistoryByDateRangeUsecase {
	return &getHistoryByDateRangeUsecase{
		historyRepo: historyRepo,
	}
}

func (u *getHistoryByDateRangeUsecase) Execute(ctx context.Context, req *GetHistoryByDateRangeRequest) ([]*entity.GrowingZoneHistory, error) {
	// Validate date format
	_, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, ErrInvalidStartDateFormat
	}

	_, err = time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, ErrInvalidEndDateFormat
	}

	// Lấy history theo khoảng thời gian
	histories, err := u.historyRepo.GetByDateRange(ctx, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

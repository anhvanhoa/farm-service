package greenhouse_installation_log

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

type GetLogsByDateRangeRequest struct {
	StartDate string
	EndDate   string
}

type GetLogsByDateRangeUsecase interface {
	Execute(ctx context.Context, req *GetLogsByDateRangeRequest) ([]*entity.GreenhouseInstallationLog, error)
}

type getLogsByDateRangeUsecase struct {
	logRepo repository.GreenhouseInstallationLogRepository
}

func NewGetLogsByDateRangeUsecase(logRepo repository.GreenhouseInstallationLogRepository) GetLogsByDateRangeUsecase {
	return &getLogsByDateRangeUsecase{
		logRepo: logRepo,
	}
}

func (u *getLogsByDateRangeUsecase) Execute(ctx context.Context, req *GetLogsByDateRangeRequest) ([]*entity.GreenhouseInstallationLog, error) {
	// Validate date format
	_, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, ErrInvalidStartDateFormat
	}

	_, err = time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, ErrInvalidEndDateFormat
	}

	// Lấy logs theo khoảng thời gian
	logs, err := u.logRepo.GetByDateRange(ctx, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

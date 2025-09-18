package growing_zone_history

import (
	"farm-service/domain/repository"
)

type GrowingZoneHistoryUsecase struct {
	CreateHistory           CreateHistoryUsecase
	GetHistoryByZone        GetHistoryByZoneUsecase
	GetHistoryByAction      GetHistoryByActionUsecase
	GetHistoryByDateRange   GetHistoryByDateRangeUsecase
	GetHistoryByPerformedBy GetHistoryByPerformedByUsecase
}

func NewGrowingZoneHistoryUsecase(
	growingZoneRepo repository.GrowingZoneRepository,
	historyRepo repository.GrowingZoneHistoryRepository,
) *GrowingZoneHistoryUsecase {
	return &GrowingZoneHistoryUsecase{
		CreateHistory:           NewCreateHistoryUsecase(growingZoneRepo, historyRepo),
		GetHistoryByZone:        NewGetHistoryByZoneUsecase(growingZoneRepo, historyRepo),
		GetHistoryByAction:      NewGetHistoryByActionUsecase(historyRepo),
		GetHistoryByDateRange:   NewGetHistoryByDateRangeUsecase(historyRepo),
		GetHistoryByPerformedBy: NewGetHistoryByPerformedByUsecase(historyRepo),
	}
}

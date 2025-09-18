package growing_zone_history

import (
	"farm-service/domain/repository"
)

type BaseUseCase struct {
	GrowingZoneRepo repository.GrowingZoneRepository
	HistoryRepo     repository.GrowingZoneHistoryRepository
}

func NewBaseUseCase(
	growingZoneRepo repository.GrowingZoneRepository,
	historyRepo repository.GrowingZoneHistoryRepository,
) *BaseUseCase {
	return &BaseUseCase{
		GrowingZoneRepo: growingZoneRepo,
		HistoryRepo:     historyRepo,
	}
}

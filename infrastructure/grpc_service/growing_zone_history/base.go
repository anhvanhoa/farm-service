package growing_zone_history_service

import (
	"farm-service/domain/entity"
	"farm-service/domain/repository"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
)

type GrowingZoneHistoryService struct {
	growingZoneHistoryP.UnsafeGrowingZoneHistoryServiceServer
	growingZoneHistoryUsecase growing_zone_history.GrowingZoneHistoryUsecase
}

func NewGrowingZoneHistoryService(growingZoneRepository repository.GrowingZoneRepository, growingZoneHistoryRepository repository.GrowingZoneHistoryRepository) growingZoneHistoryP.GrowingZoneHistoryServiceServer {
	growingZoneHistoryUsecase := growing_zone_history.NewGrowingZoneHistoryUsecase(growingZoneRepository, growingZoneHistoryRepository)
	return &GrowingZoneHistoryService{
		growingZoneHistoryUsecase: *growingZoneHistoryUsecase,
	}
}

func (s *GrowingZoneHistoryService) createProtoHistories(histories []*entity.GrowingZoneHistory) []*growingZoneHistoryP.ZoneHistory {
	protoHistories := make([]*growingZoneHistoryP.ZoneHistory, len(histories))
	for i, history := range histories {
		protoHistories[i] = s.createProtoHistory(history)
	}
	return protoHistories
}

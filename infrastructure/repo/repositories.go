package repo

import (
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type Repositories struct {
	GreenhouseRepository                repository.GreenhouseRepository
	GreenhouseInstallationLogRepository repository.GreenhouseInstallationLogRepository
	GrowingZoneRepository               repository.GrowingZoneRepository
	GrowingZoneHistoryRepository        repository.GrowingZoneHistoryRepository
}

func NewRepositories(db *pg.DB) *Repositories {
	return &Repositories{
		GreenhouseRepository:                NewGreenhouseRepo(db),
		GreenhouseInstallationLogRepository: NewGreenhouseInstallationLogRepo(db),
		GrowingZoneRepository:               NewGrowingZoneRepo(db),
		GrowingZoneHistoryRepository:        NewGrowingZoneHistoryRepo(db),
	}
}

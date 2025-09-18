package greenhouse_installation_log

import (
	"farm-service/domain/repository"
)

type BaseUseCase struct {
	GreenhouseRepo repository.GreenhouseRepository
	LogRepo        repository.GreenhouseInstallationLogRepository
}

func NewBaseUseCase(
	greenhouseRepo repository.GreenhouseRepository,
	logRepo repository.GreenhouseInstallationLogRepository,
) *BaseUseCase {
	return &BaseUseCase{
		GreenhouseRepo: greenhouseRepo,
		LogRepo:        logRepo,
	}
}

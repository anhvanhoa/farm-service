package greenhouse_installation_log

import (
	"farm-service/domain/repository"
)

type GreenhouseInstallationLogUsecase struct {
	CreateLog           CreateLogUsecase
	GetLogsByGreenhouse GetLogsByGreenhouseUsecase
	GetLogsByAction     GetLogsByActionUsecase
	GetLogsByDateRange  GetLogsByDateRangeUsecase
}

func NewGreenhouseInstallationLogUsecase(
	greenhouseRepo repository.GreenhouseRepository,
	logRepo repository.GreenhouseInstallationLogRepository,
) *GreenhouseInstallationLogUsecase {
	return &GreenhouseInstallationLogUsecase{
		CreateLog:           NewCreateLogUsecase(greenhouseRepo, logRepo),
		GetLogsByGreenhouse: NewGetLogsByGreenhouseUsecase(greenhouseRepo, logRepo),
		GetLogsByAction:     NewGetLogsByActionUsecase(logRepo),
		GetLogsByDateRange:  NewGetLogsByDateRangeUsecase(logRepo),
	}
}

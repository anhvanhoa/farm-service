package greenhouse_installation_log_service

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/usecase/greenhouse_installation_log"

	greenhouseInstallationLogP "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
)

func (s *GreenhouseInstallationLogService) GetLogsByGreenhouse(ctx context.Context, req *greenhouseInstallationLogP.GetLogsByGreenhouseRequest) (*greenhouseInstallationLogP.GetLogsByGreenhouseResponse, error) {
	greenhouse := s.createGetLogsByGreenhouseReq(req)
	logs, err := s.greenhouseInstallationLogUsecase.GetLogsByGreenhouse.Execute(ctx, greenhouse)
	if err != nil {
		return nil, err
	}
	return &greenhouseInstallationLogP.GetLogsByGreenhouseResponse{
		Success: true,
		Message: "Logs retrieved successfully",
		Logs:    s.createProtoLogs(logs),
	}, nil
}

func (s *GreenhouseInstallationLogService) createProtoLogs(logs []*entity.GreenhouseInstallationLog) []*greenhouseInstallationLogP.InstallationLog {
	protoLogs := make([]*greenhouseInstallationLogP.InstallationLog, len(logs))
	for i, log := range logs {
		protoLogs[i] = s.createProtoLog(log)
	}
	return protoLogs
}

func (s *GreenhouseInstallationLogService) createGetLogsByGreenhouseReq(req *greenhouseInstallationLogP.GetLogsByGreenhouseRequest) *greenhouse_installation_log.GetLogsByGreenhouseRequest {
	return &greenhouse_installation_log.GetLogsByGreenhouseRequest{
		GreenhouseID: req.GreenhouseId,
	}
}

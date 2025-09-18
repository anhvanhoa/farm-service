package greenhouse_installation_log_service

import (
	"context"
	"farm-service/domain/usecase/greenhouse_installation_log"

	greenhouseInstallationLogP "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
)

func (s *GreenhouseInstallationLogService) GetLogsByDateRange(ctx context.Context, req *greenhouseInstallationLogP.GetLogsByDateRangeRequest) (*greenhouseInstallationLogP.GetLogsByDateRangeResponse, error) {
	dateRangeReq := &greenhouse_installation_log.GetLogsByDateRangeRequest{
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	logs, err := s.greenhouseInstallationLogUsecase.GetLogsByDateRange.Execute(ctx, dateRangeReq)
	if err != nil {
		return nil, err
	}
	return &greenhouseInstallationLogP.GetLogsByDateRangeResponse{
		Success: true,
		Message: "Logs retrieved successfully",
		Logs:    s.createProtoLogs(logs),
	}, nil
}

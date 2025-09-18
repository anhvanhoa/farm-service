package greenhouse_installation_log_service

import (
	"context"
	"farm-service/domain/usecase/greenhouse_installation_log"

	greenhouseInstallationLogP "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
)

func (s *GreenhouseInstallationLogService) GetLogsByAction(ctx context.Context, req *greenhouseInstallationLogP.GetLogsByActionRequest) (*greenhouseInstallationLogP.GetLogsByActionResponse, error) {
	action := s.createGetLogsByActionReq(req)
	logs, err := s.greenhouseInstallationLogUsecase.GetLogsByAction.Execute(ctx, action)
	if err != nil {
		return nil, err
	}
	return &greenhouseInstallationLogP.GetLogsByActionResponse{
		Success: true,
		Message: "Logs retrieved successfully",
		Logs:    s.createProtoLogs(logs),
	}, nil
}

func (s *GreenhouseInstallationLogService) createGetLogsByActionReq(req *greenhouseInstallationLogP.GetLogsByActionRequest) *greenhouse_installation_log.GetLogsByActionRequest {
	return &greenhouse_installation_log.GetLogsByActionRequest{
		Action: req.Action.String(),
	}
}

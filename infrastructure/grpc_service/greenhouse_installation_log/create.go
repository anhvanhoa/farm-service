package greenhouse_installation_log_service

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/usecase/greenhouse_installation_log"
	"time"

	greenhouseInstallationLogP "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GreenhouseInstallationLogService) CreateLog(ctx context.Context, req *greenhouseInstallationLogP.CreateLogRequest) (*greenhouseInstallationLogP.CreateLogResponse, error) {
	logReq, err := s.createEntityLogReq(req)
	if err != nil {
		return nil, err
	}
	log, err := s.greenhouseInstallationLogUsecase.CreateLog.Execute(ctx, logReq)
	if err != nil {
		return nil, err
	}
	return &greenhouseInstallationLogP.CreateLogResponse{
		Success: true,
		Message: "Installation log created successfully",
		Log:     s.createProtoLog(log),
	}, nil
}

func (s *GreenhouseInstallationLogService) createEntityLogReq(req *greenhouseInstallationLogP.CreateLogRequest) (*greenhouse_installation_log.CreateLogRequest, error) {
	log := &greenhouse_installation_log.CreateLogRequest{
		GreenhouseID: req.GreenhouseId,
		Action:       req.Action.String(),
		Description:  req.Description,
		PerformedBy:  req.PerformedBy,
	}

	actionDate, err := time.Parse("2006-01-02", req.ActionDate)
	if err != nil {
		return nil, err
	}
	log.ActionDate = actionDate

	return log, nil
}

func (s *GreenhouseInstallationLogService) createProtoLog(log *entity.GreenhouseInstallationLog) *greenhouseInstallationLogP.InstallationLog {
	// Convert action string to enum
	var action greenhouseInstallationLogP.InstallationAction
	switch log.Action {
	case "INSTALLATION_ACTION_INSTALL":
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_INSTALL
	case "INSTALLATION_ACTION_UPGRADE":
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_UPGRADE
	case "INSTALLATION_ACTION_MAINTENANCE":
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_MAINTENANCE
	case "INSTALLATION_ACTION_RELOCATE":
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_RELOCATE
	case "INSTALLATION_ACTION_DISMANTLE":
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_DISMANTLE
	default:
		action = greenhouseInstallationLogP.InstallationAction_INSTALLATION_ACTION_UNSPECIFIED
	}

	response := &greenhouseInstallationLogP.InstallationLog{
		Id:           log.ID,
		GreenhouseId: log.GreenhouseID,
		Action:       action,
		ActionDate:   timestamppb.New(log.ActionDate),
		Description:  log.Description,
		PerformedBy:  log.PerformedBy,
		CreatedAt:    timestamppb.New(log.CreatedAt),
	}

	return response
}

package grpc_service

import (
	"farm-service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	greenhouse_installation_log "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	"github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	growing_zone_history "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	greenhouseService greenhouse.GreenhouseServiceServer,
	growingZoneService growing_zone.GrowingZoneServiceServer,
	greenhouseInstallationLogService greenhouse_installation_log.GreenhouseInstallationLogServiceServer,
	growingZoneHistoryService growing_zone_history.GrowingZoneHistoryServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			greenhouse.RegisterGreenhouseServiceServer(server, greenhouseService)
			growing_zone.RegisterGrowingZoneServiceServer(server, growingZoneService)
			greenhouse_installation_log.RegisterGreenhouseInstallationLogServiceServer(server, greenhouseInstallationLogService)
			growing_zone_history.RegisterGrowingZoneHistoryServiceServer(server, growingZoneHistoryService)
		},
	)
}

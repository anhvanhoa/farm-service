package main

import (
	"context"
	"farm-service/bootstrap"
	"farm-service/infrastructure/grpc_service"
	greenhouse_service "farm-service/infrastructure/grpc_service/greenhouse"
	gil_service "farm-service/infrastructure/grpc_service/greenhouse_installation_log"
	growingzone_service "farm-service/infrastructure/grpc_service/growing_zone"
	gzh_service "farm-service/infrastructure/grpc_service/growing_zone_history"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	greenhouseService := greenhouse_service.NewGreenhouseService(app.Repos.GreenhouseRepository)
	growingZoneService := growingzone_service.NewGrowingZoneService(app.Repos.GrowingZoneRepository)
	greenhouseInstallationLogService := gil_service.NewGreenhouseInstallationLogService(
		app.Repos.GreenhouseRepository,
		app.Repos.GreenhouseInstallationLogRepository,
	)
	growingZoneHistoryService := gzh_service.NewGrowingZoneHistoryService(
		app.Repos.GrowingZoneRepository,
		app.Repos.GrowingZoneHistoryRepository,
	)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		greenhouseService,
		growingZoneService,
		greenhouseInstallationLogService,
		growingZoneHistoryService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}

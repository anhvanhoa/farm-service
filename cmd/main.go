package main

import (
	"context"
	"farm-service/bootstrap"
	"farm-service/infrastructure/grpc_client"
	"farm-service/infrastructure/grpc_service"
	greenhouse_service "farm-service/infrastructure/grpc_service/greenhouse"
	gil_service "farm-service/infrastructure/grpc_service/greenhouse_installation_log"
	growingzone_service "farm-service/infrastructure/grpc_service/growing_zone"
	gzh_service "farm-service/infrastructure/grpc_service/growing_zone_history"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

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
		env, log, app.Cache,
		greenhouseService,
		growingZoneService,
		greenhouseInstallationLogService,
		growingZoneHistoryService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}

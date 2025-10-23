package grpc_service

import (
	"farm-service/bootstrap"

	grpc_service "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/domain/user_context"
	"github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	greenhouse_installation_log "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	"github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	growing_zone_history "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	cache cache.CacheI,
	greenhouseService greenhouse.GreenhouseServiceServer,
	growingZoneService growing_zone.GrowingZoneServiceServer,
	greenhouseInstallationLogService greenhouse_installation_log.GreenhouseInstallationLogServiceServer,
	growingZoneHistoryService growing_zone_history.GrowingZoneHistoryServiceServer,
) *grpc_service.GRPCServer {
	config := &grpc_service.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	middleware := grpc_service.NewMiddleware()
	return grpc_service.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			greenhouse.RegisterGreenhouseServiceServer(server, greenhouseService)
			growing_zone.RegisterGrowingZoneServiceServer(server, growingZoneService)
			greenhouse_installation_log.RegisterGreenhouseInstallationLogServiceServer(server, greenhouseInstallationLogService)
			growing_zone_history.RegisterGrowingZoneHistoryServiceServer(server, growingZoneHistoryService)
		},
		middleware.AuthorizationInterceptor(
			env.SecretService,
			func(action string, resource string) bool {
				hasPermission, err := cache.Get(resource + "." + action)
				if err != nil {
					return false
				}
				return hasPermission != nil && string(hasPermission) == "true"
			},
			func(id string) *user_context.UserContext {
				userData, err := cache.Get(id)
				if err != nil || userData == nil {
					return nil
				}
				uCtx := user_context.NewUserContext()
				uCtx.FromBytes(userData)
				return uCtx
			},
		),
	)
}

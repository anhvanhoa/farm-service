package greenhouse_service

import (
	"farm-service/domain/repository"
	"farm-service/domain/usecase/greenhouse"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
)

type GreenhouseService struct {
	greenhouseP.UnsafeGreenhouseServiceServer
	greenhouseUsecase greenhouse.GreenhouseUsecase
}

func NewGreenhouseService(greenhouseRepository repository.GreenhouseRepository) greenhouseP.GreenhouseServiceServer {
	greenhouseUsecase := greenhouse.NewGreenhouseUsecase(greenhouseRepository)
	return &GreenhouseService{
		greenhouseUsecase: *greenhouseUsecase,
	}
}

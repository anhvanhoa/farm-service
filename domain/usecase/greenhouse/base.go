package greenhouse

import (
	"farm-service/domain/repository"
)

type GreenhouseUsecase struct {
	CreateGreenhouse CreateGreenhouseUsecase
	GetGreenhouse    GetGreenhouseUsecase
	UpdateGreenhouse UpdateGreenhouseUsecase
	DeleteGreenhouse DeleteGreenhouseUsecase
	ListGreenhouse   ListGreenhouseUsecase
}

func NewGreenhouseUsecase(greenhouseRepository repository.GreenhouseRepository) *GreenhouseUsecase {
	return &GreenhouseUsecase{
		CreateGreenhouse: NewCreateGreenhouseUsecase(greenhouseRepository),
		GetGreenhouse:    NewGetGreenhouseUsecase(greenhouseRepository),
		UpdateGreenhouse: NewUpdateGreenhouseUsecase(greenhouseRepository),
		DeleteGreenhouse: NewDeleteGreenhouseUsecase(greenhouseRepository),
		ListGreenhouse:   NewListGreenhouseUsecase(greenhouseRepository),
	}
}

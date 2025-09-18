package growingzone

import "farm-service/domain/repository"

type GrowingZoneUsecase struct {
	CreateGrowingZone    CreateGrowingZoneUsecase
	GetGrowingZone       GetGrowingZoneUsecase
	UpdateGrowingZone    UpdateGrowingZoneUsecase
	DeleteGrowingZone    DeleteGrowingZoneUsecase
	ListGrowingZone      ListGrowingZoneUsecase
	GetZonesByGreenhouse GetZonesByGreenhouseUsecase
}

func NewGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) *GrowingZoneUsecase {
	return &GrowingZoneUsecase{
		CreateGrowingZone:    NewCreateGrowingZoneUsecase(growingZoneRepository),
		GetGrowingZone:       NewGetGrowingZoneUsecase(growingZoneRepository),
		UpdateGrowingZone:    NewUpdateGrowingZoneUsecase(growingZoneRepository),
		DeleteGrowingZone:    NewDeleteGrowingZoneUsecase(growingZoneRepository),
		ListGrowingZone:      NewListGrowingZoneUsecase(growingZoneRepository),
		GetZonesByGreenhouse: NewGetZonesByGreenhouseUsecase(growingZoneRepository),
	}
}

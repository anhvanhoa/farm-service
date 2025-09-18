package repository

import (
	"context"
	"farm-service/domain/entity"
)

type GrowingZoneRepository interface {
	Create(ctx context.Context, zone *entity.GrowingZone) error

	GetByID(ctx context.Context, id string) (*entity.GrowingZone, error)

	GetByZoneCode(ctx context.Context, zoneCode string) (*entity.GrowingZone, error)

	Update(ctx context.Context, id string, zone *entity.UpdateGrowingZoneRequest) error

	Delete(ctx context.Context, id string) error

	List(ctx context.Context, filter *entity.GrowingZoneFilter, page, pageSize int) ([]*entity.GrowingZone, int64, error)

	GetByGreenhouseID(ctx context.Context, greenhouseID string) ([]*entity.GrowingZone, error)

	GetByStatus(ctx context.Context, status string) ([]*entity.GrowingZone, error)

	GetBySoilType(ctx context.Context, soilType string) ([]*entity.GrowingZone, error)

	GetByIrrigationSystem(ctx context.Context, irrigationSystem string) ([]*entity.GrowingZone, error)

	Count(ctx context.Context, filter *entity.GrowingZoneFilter) (int64, error)

	CheckZoneCodeExists(ctx context.Context, zoneCode string) (bool, error)
}

type GrowingZoneHistoryRepository interface {
	Create(ctx context.Context, history *entity.GrowingZoneHistory) error

	GetByZoneID(ctx context.Context, zoneID string) ([]*entity.GrowingZoneHistory, error)

	GetByAction(ctx context.Context, action string) ([]*entity.GrowingZoneHistory, error)

	GetByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.GrowingZoneHistory, error)

	GetByPerformedBy(ctx context.Context, performedBy string) ([]*entity.GrowingZoneHistory, error)
}

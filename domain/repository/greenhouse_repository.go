package repository

import (
	"context"
	"farm-service/domain/entity"
)

type GreenhouseRepository interface {
	Create(ctx context.Context, greenhouse *entity.Greenhouse) error

	GetByID(ctx context.Context, id string) (*entity.Greenhouse, error)

	GetByCode(ctx context.Context, code string) (*entity.Greenhouse, error)

	Update(ctx context.Context, id string, greenhouse *entity.UpdateGreenhouseRequest) error

	Delete(ctx context.Context, id string) error

	List(ctx context.Context, filter *entity.GreenhouseFilter, page, pageSize int) ([]*entity.Greenhouse, int64, error)

	GetByStatus(ctx context.Context, status string) ([]*entity.Greenhouse, error)

	GetByLocation(ctx context.Context, location string) ([]*entity.Greenhouse, error)

	Count(ctx context.Context, filter *entity.GreenhouseFilter) (int64, error)
}

type GreenhouseInstallationLogRepository interface {
	Create(ctx context.Context, log *entity.GreenhouseInstallationLog) error

	GetByGreenhouseID(ctx context.Context, greenhouseID string) ([]*entity.GreenhouseInstallationLog, error)

	GetByAction(ctx context.Context, action string) ([]*entity.GreenhouseInstallationLog, error)

	GetByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.GreenhouseInstallationLog, error)
}

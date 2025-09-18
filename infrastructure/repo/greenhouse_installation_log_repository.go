package repo

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type greenhouseInstallationLogRepository struct {
	db *pg.DB
}

func NewGreenhouseInstallationLogRepo(db *pg.DB) repository.GreenhouseInstallationLogRepository {
	return &greenhouseInstallationLogRepository{db: db}
}

func (r *greenhouseInstallationLogRepository) Create(ctx context.Context, log *entity.GreenhouseInstallationLog) error {
	_, err := r.db.ModelContext(ctx, log).Insert()
	return err
}

func (r *greenhouseInstallationLogRepository) GetByGreenhouseID(ctx context.Context, greenhouseID string) ([]*entity.GreenhouseInstallationLog, error) {
	var logs []*entity.GreenhouseInstallationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("greenhouse_id = ?", greenhouseID).
		Order("action_date DESC").
		Select()
	return logs, err
}

func (r *greenhouseInstallationLogRepository) GetByAction(ctx context.Context, action string) ([]*entity.GreenhouseInstallationLog, error) {
	var logs []*entity.GreenhouseInstallationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("action = ?", action).
		Order("action_date DESC").
		Select()
	return logs, err
}

func (r *greenhouseInstallationLogRepository) GetByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.GreenhouseInstallationLog, error) {
	var logs []*entity.GreenhouseInstallationLog
	err := r.db.ModelContext(ctx, &logs).
		Where("action_date >= ? AND action_date <= ?", startDate, endDate).
		Order("action_date DESC").
		Select()
	return logs, err
}

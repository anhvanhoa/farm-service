package repo

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type growingZoneHistoryRepository struct {
	db *pg.DB
}

func NewGrowingZoneHistoryRepo(db *pg.DB) repository.GrowingZoneHistoryRepository {
	return &growingZoneHistoryRepository{db: db}
}

func (r *growingZoneHistoryRepository) Create(ctx context.Context, history *entity.GrowingZoneHistory) error {
	_, err := r.db.ModelContext(ctx, history).Insert()
	return err
}

func (r *growingZoneHistoryRepository) GetByZoneID(ctx context.Context, zoneID string) ([]*entity.GrowingZoneHistory, error) {
	var histories []*entity.GrowingZoneHistory
	err := r.db.ModelContext(ctx, &histories).
		Where("zone_id = ?", zoneID).
		Order("action_date DESC").
		Select()
	return histories, err
}

func (r *growingZoneHistoryRepository) GetByAction(ctx context.Context, action string) ([]*entity.GrowingZoneHistory, error) {
	var histories []*entity.GrowingZoneHistory
	err := r.db.ModelContext(ctx, &histories).
		Where("action = ?", action).
		Order("action_date DESC").
		Select()
	return histories, err
}

func (r *growingZoneHistoryRepository) GetByDateRange(ctx context.Context, startDate, endDate string) ([]*entity.GrowingZoneHistory, error) {
	var histories []*entity.GrowingZoneHistory
	err := r.db.ModelContext(ctx, &histories).
		Where("action_date >= ? AND action_date <= ?", startDate, endDate).
		Order("action_date DESC").
		Select()
	return histories, err
}

func (r *growingZoneHistoryRepository) GetByPerformedBy(ctx context.Context, performedBy string) ([]*entity.GrowingZoneHistory, error) {
	var histories []*entity.GrowingZoneHistory
	err := r.db.ModelContext(ctx, &histories).
		Where("performed_by = ?", performedBy).
		Order("action_date DESC").
		Select()
	return histories, err
}

package repo

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type growingZoneRepository struct {
	db *pg.DB
}

func NewGrowingZoneRepo(db *pg.DB) repository.GrowingZoneRepository {
	return &growingZoneRepository{db: db}
}

func (r *growingZoneRepository) Create(ctx context.Context, zone *entity.GrowingZone) error {
	_, err := r.db.ModelContext(ctx, zone).Insert()
	return err
}

func (r *growingZoneRepository) GetByID(ctx context.Context, id string) (*entity.GrowingZone, error) {
	zone := &entity.GrowingZone{}
	err := r.db.ModelContext(ctx, zone).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return zone, nil
}

func (r *growingZoneRepository) GetByZoneCode(ctx context.Context, zoneCode string) (*entity.GrowingZone, error) {
	zone := &entity.GrowingZone{}
	err := r.db.ModelContext(ctx, zone).Where("zone_code = ?", zoneCode).Select()
	if err != nil {
		return nil, err
	}
	return zone, nil
}

func (r *growingZoneRepository) Update(ctx context.Context, id string, updateReq *entity.UpdateGrowingZoneRequest) error {
	zone := &entity.GrowingZone{
		ZoneName:         updateReq.ZoneName,
		ZoneCode:         updateReq.ZoneCode,
		AreaM2:           *updateReq.AreaM2,
		MaxPlants:        *updateReq.MaxPlants,
		SoilType:         updateReq.SoilType,
		IrrigationSystem: updateReq.IrrigationSystem,
		Status:           updateReq.Status,
	}
	_, err := r.db.ModelContext(ctx, zone).Where("id = ?", id).UpdateNotZero()
	return err
}

func (r *growingZoneRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.GrowingZone)(nil)).Where("id = ?", id).Delete()
	return err
}

func (r *growingZoneRepository) List(ctx context.Context, filter *entity.GrowingZoneFilter, page, pageSize int) ([]*entity.GrowingZone, int64, error) {
	var zones []*entity.GrowingZone
	query := r.db.ModelContext(ctx, &zones)

	if filter != nil {
		if filter.GreenhouseID != "" {
			query = query.Where("greenhouse_id = ?", filter.GreenhouseID)
		}
		if filter.Status != "" {
			query = query.Where("status = ?", filter.Status)
		}
		if filter.SoilType != "" {
			query = query.Where("soil_type = ?", filter.SoilType)
		}
		if filter.IrrigationSystem != "" {
			query = query.Where("irrigation_system = ?", filter.IrrigationSystem)
		}
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Limit(pageSize).Offset(offset).Order("created_at DESC").Select()
	if err != nil {
		return nil, 0, err
	}

	return zones, int64(total), nil
}

func (r *growingZoneRepository) GetByGreenhouseID(ctx context.Context, greenhouseID string) ([]*entity.GrowingZone, error) {
	var zones []*entity.GrowingZone
	err := r.db.ModelContext(ctx, &zones).
		Where("greenhouse_id = ?", greenhouseID).
		Order("created_at DESC").
		Select()
	return zones, err
}

func (r *growingZoneRepository) GetByStatus(ctx context.Context, status string) ([]*entity.GrowingZone, error) {
	var zones []*entity.GrowingZone
	err := r.db.ModelContext(ctx, &zones).
		Where("status = ?", status).
		Order("created_at DESC").
		Select()
	return zones, err
}

func (r *growingZoneRepository) GetBySoilType(ctx context.Context, soilType string) ([]*entity.GrowingZone, error) {
	var zones []*entity.GrowingZone
	err := r.db.ModelContext(ctx, &zones).
		Where("soil_type = ?", soilType).
		Order("created_at DESC").
		Select()
	return zones, err
}

func (r *growingZoneRepository) GetByIrrigationSystem(ctx context.Context, irrigationSystem string) ([]*entity.GrowingZone, error) {
	var zones []*entity.GrowingZone
	err := r.db.ModelContext(ctx, &zones).
		Where("irrigation_system = ?", irrigationSystem).
		Order("created_at DESC").
		Select()
	return zones, err
}

func (r *growingZoneRepository) Count(ctx context.Context, filter *entity.GrowingZoneFilter) (int64, error) {
	query := r.db.ModelContext(ctx, (*entity.GrowingZone)(nil))

	if filter != nil {
		if filter.GreenhouseID != "" {
			query = query.Where("greenhouse_id = ?", filter.GreenhouseID)
		}
		if filter.Status != "" {
			query = query.Where("status = ?", filter.Status)
		}
		if filter.SoilType != "" {
			query = query.Where("soil_type = ?", filter.SoilType)
		}
		if filter.IrrigationSystem != "" {
			query = query.Where("irrigation_system = ?", filter.IrrigationSystem)
		}
	}

	count, err := query.Count()
	return int64(count), err
}

func (r *growingZoneRepository) CheckZoneCodeExists(ctx context.Context, zoneCode string) (bool, error) {
	count, err := r.db.ModelContext(ctx, (*entity.GrowingZone)(nil)).
		Where("zone_code = ?", zoneCode).
		Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

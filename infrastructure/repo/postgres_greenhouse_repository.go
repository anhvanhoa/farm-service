package repo

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type greenhouseRepository struct {
	db *pg.DB
}

func NewGreenhouseRepo(db *pg.DB) repository.GreenhouseRepository {
	return &greenhouseRepository{db: db}
}

func (r *greenhouseRepository) Create(ctx context.Context, greenhouse *entity.Greenhouse) error {
	_, err := r.db.ModelContext(ctx, greenhouse).Insert()
	return err
}

func (r *greenhouseRepository) GetByID(ctx context.Context, id string) (*entity.Greenhouse, error) {
	greenhouse := &entity.Greenhouse{}
	err := r.db.ModelContext(ctx, greenhouse).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return greenhouse, nil
}

func (r *greenhouseRepository) GetByCode(ctx context.Context, code string) (*entity.Greenhouse, error) {
	greenhouse := &entity.Greenhouse{}
	err := r.db.ModelContext(ctx, greenhouse).Where("code = ?", code).Select()
	if err != nil {
		return nil, err
	}
	return greenhouse, nil
}

func (r *greenhouseRepository) Update(ctx context.Context, id string, updateReq *entity.UpdateGreenhouseRequest) error {
	greenhouse := &entity.Greenhouse{
		Name:             updateReq.Name,
		Location:         updateReq.Location,
		AreaM2:           *updateReq.AreaM2,
		Type:             updateReq.Type,
		MaxCapacity:      *updateReq.MaxCapacity,
		InstallationDate: updateReq.InstallationDate,
		Status:           updateReq.Status,
		Description:      updateReq.Description,
	}

	if _, err := r.db.ModelContext(ctx, greenhouse).Where("id = ?", id).UpdateNotZero(); err != nil {
		return err
	}
	return nil
}

func (r *greenhouseRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ModelContext(ctx, (*entity.Greenhouse)(nil)).Where("id = ?", id).Delete()
	return err
}

func (r *greenhouseRepository) List(ctx context.Context, filter *entity.GreenhouseFilter, page, pageSize int) ([]*entity.Greenhouse, int64, error) {
	var greenhouses []*entity.Greenhouse
	query := r.db.ModelContext(ctx, &greenhouses)

	if filter != nil {
		if filter.Status != "" {
			query = query.Where("status = ?", filter.Status)
		}
		if filter.Type != "" {
			query = query.Where("type = ?", filter.Type)
		}
		if filter.Location != "" {
			query = query.Where("location ILIKE ?", "%"+filter.Location+"%")
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

	return greenhouses, int64(total), nil
}

func (r *greenhouseRepository) GetByStatus(ctx context.Context, status string) ([]*entity.Greenhouse, error) {
	var greenhouses []*entity.Greenhouse
	err := r.db.ModelContext(ctx, &greenhouses).Where("status = ?", status).Select()
	return greenhouses, err
}

func (r *greenhouseRepository) GetByLocation(ctx context.Context, location string) ([]*entity.Greenhouse, error) {
	var greenhouses []*entity.Greenhouse
	err := r.db.ModelContext(ctx, &greenhouses).Where("location ILIKE ?", "%"+location+"%").Select()
	return greenhouses, err
}

func (r *greenhouseRepository) Count(ctx context.Context, filter *entity.GreenhouseFilter) (int64, error) {
	query := r.db.ModelContext(ctx, (*entity.Greenhouse)(nil))

	if filter != nil {
		if filter.Status != "" {
			query = query.Where("status = ?", filter.Status)
		}
		if filter.Type != "" {
			query = query.Where("type = ?", filter.Type)
		}
		if filter.Location != "" {
			query = query.Where("location ILIKE ?", "%"+filter.Location+"%")
		}
	}

	count, err := query.Count()
	return int64(count), err
}

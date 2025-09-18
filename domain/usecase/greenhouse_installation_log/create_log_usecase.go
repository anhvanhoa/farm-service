package greenhouse_installation_log

import (
	"context"
	"farm-service/domain/entity"
	"time"

	"github.com/google/uuid"
)

type CreateLogRequest struct {
	GreenhouseID string `json:"greenhouse_id" binding:"required"`
	Action       string `json:"action" binding:"required,oneof=install upgrade maintenance relocate dismantle"`
	ActionDate   string `json:"action_date" binding:"required"`
	Description  string `json:"description"`
	PerformedBy  string `json:"performed_by" binding:"required"`
}

type CreateLogUseCase struct {
	BaseUseCase
}

func NewCreateLogUseCase(baseUseCase *BaseUseCase) *CreateLogUseCase {
	return &CreateLogUseCase{
		BaseUseCase: *baseUseCase,
	}
}

func (u *CreateLogUseCase) Execute(ctx context.Context, req *CreateLogRequest) (*entity.GreenhouseInstallationLog, error) {
	// Kiểm tra greenhouse có tồn tại không
	greenhouse, err := u.GreenhouseRepo.GetByID(ctx, req.GreenhouseID)
	if err != nil {
		return nil, err
	}
	if greenhouse == nil {
		return nil, &entity.Error{
			Code:    "GREENHOUSE_NOT_FOUND",
			Message: "Greenhouse not found",
		}
	}

	// Parse action date
	actionDate, err := time.Parse("2006-01-02", req.ActionDate)
	if err != nil {
		return nil, &entity.Error{
			Code:    "INVALID_DATE_FORMAT",
			Message: "Invalid date format. Expected YYYY-MM-DD",
		}
	}

	// Tạo log mới
	log := &entity.GreenhouseInstallationLog{
		ID:           uuid.New().String(),
		GreenhouseID: req.GreenhouseID,
		Action:       req.Action,
		ActionDate:   actionDate,
		Description:  req.Description,
		PerformedBy:  req.PerformedBy,
		CreatedAt:    time.Now(),
	}

	// Lưu vào database
	err = u.LogRepo.Create(ctx, log)
	if err != nil {
		return nil, err
	}

	return log, nil
}

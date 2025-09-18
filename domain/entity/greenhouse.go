package entity

import (
	"time"
)

// Error định nghĩa struct lỗi tùy chỉnh
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

type Greenhouse struct {
	ID               string     `json:"id" db:"id"`
	Name             string     `json:"name" db:"name"`
	Location         string     `json:"location" db:"location"`
	AreaM2           float64    `json:"area_m2" db:"area_m2"`
	Type             string     `json:"type" db:"type"`
	MaxCapacity      int32      `json:"max_capacity" db:"max_capacity"`
	InstallationDate *time.Time `json:"installation_date" db:"installation_date"`
	Status           string     `json:"status" db:"status"`
	Description      string     `json:"description" db:"description"`
	CreatedBy        string     `json:"created_by" db:"created_by"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
}

type GreenhouseInstallationLog struct {
	ID           string    `json:"id" db:"id"`
	GreenhouseID string    `json:"greenhouse_id" db:"greenhouse_id"`
	Action       string    `json:"action" db:"action"`
	ActionDate   time.Time `json:"action_date" db:"action_date"`
	Description  string    `json:"description" db:"description"`
	PerformedBy  string    `json:"performed_by" db:"performed_by"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type CreateGreenhouseRequest struct {
	Name             string     `json:"name" binding:"required"`
	Location         string     `json:"location"`
	AreaM2           float64    `json:"area_m2" binding:"min=0"`
	Type             string     `json:"type" binding:"required"`
	MaxCapacity      int32      `json:"max_capacity" binding:"min=1"`
	InstallationDate *time.Time `json:"installation_date"`
	Description      string     `json:"description"`
	CreatedBy        string     `json:"created_by" binding:"required"`
}

type UpdateGreenhouseRequest struct {
	Name             string     `json:"name"`
	Location         string     `json:"location"`
	AreaM2           *float64   `json:"area_m2"`
	Type             string     `json:"type"`
	MaxCapacity      *int32     `json:"max_capacity"`
	InstallationDate *time.Time `json:"installation_date"`
	Status           string     `json:"status"`
	Description      string     `json:"description"`
}

type GreenhouseFilter struct {
	Status   string `json:"status" form:"status"`
	Type     string `json:"type" form:"type"`
	Location string `json:"location" form:"location"`
}

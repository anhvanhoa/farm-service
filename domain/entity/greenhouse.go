package entity

import (
	"time"
)

type Greenhouse struct {
	tableName        struct{} `pg:"greenhouses"`
	ID               string
	Name             string
	Location         string
	AreaM2           float64
	Type             string
	MaxCapacity      int32
	InstallationDate *time.Time
	Status           string
	Description      string
	CreatedBy        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (g *Greenhouse) GetTableName() any {
	return g.tableName
}

type GreenhouseInstallationLog struct {
	tableName    struct{} `pg:"greenhouse_installation_logs"`
	ID           string   `pg:"id,pk"`
	GreenhouseID string
	Action       string
	ActionDate   time.Time
	Description  string
	PerformedBy  string
	CreatedAt    time.Time
}

func (g *GreenhouseInstallationLog) GetTableName() any {
	return g.tableName
}

type CreateGreenhouseRequest struct {
	Name             string
	Location         string
	AreaM2           float64
	Type             string
	MaxCapacity      int32
	InstallationDate *time.Time
	Description      string
	CreatedBy        string
}

type UpdateGreenhouseRequest struct {
	Name             string
	Location         string
	AreaM2           *float64
	Type             string
	MaxCapacity      *int32
	InstallationDate *time.Time
	Status           string
	Description      string
}

type GreenhouseFilter struct {
	Status   string
	Type     string
	Location string
}

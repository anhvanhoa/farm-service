package entity

import (
	"time"
)

type GrowingZone struct {
	tableName        struct{} `pg:"growing_zones"`
	ID               string
	GreenhouseID     string `json:"greenhouse_id" db:"greenhouse_id"`
	ZoneName         string
	ZoneCode         string
	AreaM2           float64
	MaxPlants        int32
	SoilType         string
	IrrigationSystem string
	Status           string
	CreatedBy        string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

func (g *GrowingZone) GetTableName() any {
	return g.tableName
}

type GrowingZoneHistory struct {
	tableName   struct{} `pg:"growing_zone_history"`
	ID          string
	ZoneID      string
	Action      string
	OldValue    map[string]any
	NewValue    map[string]any
	ActionDate  time.Time
	PerformedBy string
	Notes       string
}

func (g *GrowingZoneHistory) GetTableName() any {
	return g.tableName
}

type CreateGrowingZoneRequest struct {
	GreenhouseID     string
	ZoneName         string
	ZoneCode         string
	AreaM2           float64
	MaxPlants        int32
	SoilType         string
	IrrigationSystem string
	CreatedBy        string
}

type UpdateGrowingZoneRequest struct {
	ZoneName         string
	ZoneCode         string
	AreaM2           *float64
	MaxPlants        *int32
	SoilType         string
	IrrigationSystem string
	Status           string
}

type GrowingZoneFilter struct {
	GreenhouseID     string
	Status           string
	SoilType         string
	IrrigationSystem string
}

const (
	StatusActive      = "active"
	StatusInactive    = "inactive"
	StatusMaintenance = "maintenance"
)

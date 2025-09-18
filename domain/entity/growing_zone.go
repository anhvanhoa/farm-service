package entity

import (
	"time"
)

type GrowingZone struct {
	ID               string     `json:"id" db:"id"`
	GreenhouseID     string     `json:"greenhouse_id" db:"greenhouse_id"`
	ZoneName         string     `json:"zone_name" db:"zone_name"`
	ZoneCode         string     `json:"zone_code" db:"zone_code"`
	AreaM2           float64    `json:"area_m2" db:"area_m2"`
	MaxPlants        int32      `json:"max_plants" db:"max_plants"`
	SoilType         string     `json:"soil_type" db:"soil_type"`
	IrrigationSystem string     `json:"irrigation_system" db:"irrigation_system"`
	Status           string     `json:"status" db:"status"`
	CreatedBy        string     `json:"created_by" db:"created_by"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at" db:"updated_at"`
}

type GrowingZoneHistory struct {
	ID          string         `json:"id" db:"id"`
	ZoneID      string         `json:"zone_id" db:"zone_id"`
	Action      string         `json:"action" db:"action"`
	OldValue    map[string]any `json:"old_value" db:"old_value"`
	NewValue    map[string]any `json:"new_value" db:"new_value"`
	ActionDate  time.Time      `json:"action_date" db:"action_date"`
	PerformedBy string         `json:"performed_by" db:"performed_by"`
	Notes       string         `json:"notes" db:"notes"`
}

type CreateGrowingZoneRequest struct {
	GreenhouseID     string  `json:"greenhouse_id" binding:"required"`
	ZoneName         string  `json:"zone_name" binding:"required"`
	ZoneCode         string  `json:"zone_code" binding:"required"`
	AreaM2           float64 `json:"area_m2" binding:"min=0"`
	MaxPlants        int32   `json:"max_plants" binding:"min=0"`
	SoilType         string  `json:"soil_type"`
	IrrigationSystem string  `json:"irrigation_system"`
	CreatedBy        string  `json:"created_by" binding:"required"`
}

type UpdateGrowingZoneRequest struct {
	ZoneName         string   `json:"zone_name"`
	ZoneCode         string   `json:"zone_code"`
	AreaM2           *float64 `json:"area_m2"`
	MaxPlants        *int32   `json:"max_plants"`
	SoilType         string   `json:"soil_type"`
	IrrigationSystem string   `json:"irrigation_system"`
	Status           string   `json:"status"`
}

type GrowingZoneFilter struct {
	GreenhouseID     string `json:"greenhouse_id" form:"greenhouse_id"`
	Status           string `json:"status" form:"status"`
	SoilType         string `json:"soil_type" form:"soil_type"`
	IrrigationSystem string `json:"irrigation_system" form:"irrigation_system"`
}

const (
	StatusActive      = "active"
	StatusInactive    = "inactive"
	StatusMaintenance = "maintenance"
)

-- Migration: Create growing_zones table
-- Description: Tạo bảng khu vực trồng

CREATE TABLE growing_zones (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    greenhouse_id UUID NOT NULL,
    zone_name VARCHAR(255) NOT NULL,
    zone_code VARCHAR(50) UNIQUE,
    area_m2 DECIMAL(10,2),
    max_plants INTEGER,
    soil_type VARCHAR(100), -- sandy, clay, loam, hydroponic
    irrigation_system VARCHAR(100), -- drip, spray, flood, manual
    status VARCHAR(50) DEFAULT 'active', -- active, inactive, maintenance
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (greenhouse_id) REFERENCES greenhouses(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_growing_zones_greenhouse ON growing_zones (greenhouse_id);
CREATE INDEX idx_growing_zones_status ON growing_zones (status);
CREATE INDEX idx_growing_zones_soil_type ON growing_zones (soil_type);
CREATE INDEX idx_growing_zones_irrigation ON growing_zones (irrigation_system);
CREATE INDEX idx_growing_zones_created_by ON growing_zones (created_by);

-- Create unique constraint
ALTER TABLE growing_zones ADD CONSTRAINT uk_zone_code UNIQUE (zone_code);

-- Create trigger for updated_at
CREATE TRIGGER update_growing_zones_updated_at BEFORE UPDATE ON growing_zones
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

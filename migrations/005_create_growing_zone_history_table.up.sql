-- Migration: Create growing_zone_history table
-- Description: Tạo bảng lịch sử thay đổi khu vực trồng

CREATE TABLE growing_zone_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    zone_id UUID NOT NULL,
    action VARCHAR(50), -- change_soil, change_irrigation, maintenance, resize, rename
    old_value JSONB,
    new_value JSONB,
    action_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    performed_by UUID,
    notes TEXT,
    
    FOREIGN KEY (zone_id) REFERENCES growing_zones(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_zone_history ON growing_zone_history (zone_id, action_date);
CREATE INDEX idx_zone_history_action ON growing_zone_history (action);
CREATE INDEX idx_zone_history_performed_by ON growing_zone_history (performed_by);

-- Create GIN indexes for JSONB fields for better performance
CREATE INDEX idx_zone_history_old_value ON growing_zone_history USING GIN (old_value);
CREATE INDEX idx_zone_history_new_value ON growing_zone_history USING GIN (new_value);

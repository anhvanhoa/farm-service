-- Migration: Create greenhouses table
-- Description: Tạo bảng quản lý nhà lưới

CREATE TABLE IF NOT EXISTS greenhouses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    location VARCHAR(500),
    area_m2 DECIMAL(10,2),
    type VARCHAR(100), -- glass, plastic, tunnel, etc.
    max_capacity INTEGER, -- Số lượng cây tối đa
    installation_date DATE,
    status VARCHAR(50) DEFAULT 'active', -- active, inactive, maintenance
    description TEXT,
    created_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_greenhouses_status ON greenhouses (status);
CREATE INDEX idx_greenhouses_location ON greenhouses (location);
CREATE INDEX idx_greenhouses_type ON greenhouses (type);
CREATE INDEX idx_greenhouses_created_by ON greenhouses (created_by);

-- Create trigger for updated_at
CREATE TRIGGER update_greenhouses_updated_at BEFORE UPDATE ON greenhouses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

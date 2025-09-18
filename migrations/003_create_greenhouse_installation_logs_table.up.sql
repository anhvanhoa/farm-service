-- Migration: Create greenhouse_installation_logs table
-- Description: Tạo bảng log cài đặt nhà lưới

CREATE TABLE greenhouse_installation_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    greenhouse_id UUID NOT NULL,
    action VARCHAR(50), -- install, upgrade, maintenance, relocate, dismantle
    action_date DATE NOT NULL,
    description TEXT,
    performed_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (greenhouse_id) REFERENCES greenhouses(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_greenhouse_logs ON greenhouse_installation_logs (greenhouse_id, action_date);
CREATE INDEX idx_greenhouse_logs_action ON greenhouse_installation_logs (action);
CREATE INDEX idx_greenhouse_logs_performed_by ON greenhouse_installation_logs (performed_by);

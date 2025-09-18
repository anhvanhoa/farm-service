-- 1. BẢNG QUẢN LÝ NHÀ LƯỚI
CREATE TABLE greenhouses (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    location VARCHAR(500),
    area_m2 DECIMAL(10,2),
    type VARCHAR(100) COMMENT 'glass, plastic, tunnel, etc.',
    max_capacity INTEGER COMMENT 'Số lượng cây tối đa',
    installation_date DATE,
    status VARCHAR(50) DEFAULT 'active' COMMENT 'active, inactive, maintenance',
    description TEXT,
    created_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (created_by) REFERENCES users(id),
    INDEX idx_greenhouses_status (status),
    INDEX idx_greenhouses_location (location(100))
);

CREATE TABLE greenhouse_installation_logs (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    greenhouse_id VARCHAR(36) NOT NULL,
    action VARCHAR(50) COMMENT 'install, upgrade, maintenance, relocate, dismantle',
    action_date DATE NOT NULL,
    description TEXT,
    performed_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (greenhouse_id) REFERENCES greenhouses(id) ON DELETE CASCADE,
    FOREIGN KEY (performed_by) REFERENCES users(id),
    INDEX idx_greenhouse_logs (greenhouse_id, action_date)
);

-- 2. BẢNG KHU VỰC TRỒNG
CREATE TABLE growing_zones (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    greenhouse_id VARCHAR(36) NOT NULL,
    zone_name VARCHAR(255) NOT NULL,
    zone_code VARCHAR(50) UNIQUE,
    area_m2 DECIMAL(10,2),
    max_plants INTEGER,
    soil_type VARCHAR(100) COMMENT 'sandy, clay, loam, hydroponic',
    irrigation_system VARCHAR(100) COMMENT 'drip, spray, flood, manual',
    status VARCHAR(50) DEFAULT 'active',
    created_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (greenhouse_id) REFERENCES greenhouses(id) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id),
    INDEX idx_growing_zones_greenhouse (greenhouse_id),
    INDEX idx_growing_zones_status (status),
    UNIQUE KEY uk_zone_code (zone_code)
);

CREATE TABLE growing_zone_history (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    zone_id VARCHAR(36) NOT NULL,
    action VARCHAR(50) COMMENT 'change_soil, change_irrigation, maintenance, resize, rename',
    old_value JSON,
    new_value JSON,
    action_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    performed_by VARCHAR(36),
    notes TEXT,
    
    FOREIGN KEY (zone_id) REFERENCES growing_zones(id) ON DELETE CASCADE,
    FOREIGN KEY (performed_by) REFERENCES users(id),
    INDEX idx_zone_history (zone_id, action_date)
);
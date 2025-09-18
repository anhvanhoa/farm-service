-- Migration: Insert sample data
-- Description: Thêm dữ liệu mẫu để test

-- Insert sample greenhouses
INSERT INTO greenhouses (id, name, location, area_m2, type, max_capacity, installation_date, status, description, created_by) VALUES
('660e8400-e29b-41d4-a716-446655440001', 'Nhà lưới A1', 'Hà Nội, Việt Nam', 500.00, 'plastic', 1000, '2024-01-15', 'active', 'Nhà lưới chính cho trồng rau xanh', '550e8400-e29b-41d4-a716-446655440001'),
('660e8400-e29b-41d4-a716-446655440002', 'Nhà lưới B1', 'TP.HCM, Việt Nam', 750.50, 'glass', 1500, '2024-02-20', 'active', 'Nhà lưới kính cho trồng hoa', '550e8400-e29b-41d4-a716-446655440002'),
('660e8400-e29b-41d4-a716-446655440003', 'Nhà lưới C1', 'Đà Nẵng, Việt Nam', 300.25, 'tunnel', 600, '2024-03-10', 'maintenance', 'Nhà lưới tunnel đang bảo trì', '550e8400-e29b-41d4-a716-446655440001');

-- Insert sample growing zones
INSERT INTO growing_zones (id, greenhouse_id, zone_name, zone_code, area_m2, max_plants, soil_type, irrigation_system, status, created_by) VALUES
('770e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', 'Khu A1-1', 'A1-1', 100.00, 200, 'loam', 'drip', 'active', '550e8400-e29b-41d4-a716-446655440002'),
('770e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440001', 'Khu A1-2', 'A1-2', 150.00, 300, 'clay', 'spray', 'active', '550e8400-e29b-41d4-a716-446655440002'),
('770e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440002', 'Khu B1-1', 'B1-1', 200.00, 400, 'sandy', 'drip', 'active', '550e8400-e29b-41d4-a716-446655440002'),
('770e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440002', 'Khu B1-2', 'B1-2', 180.50, 360, 'hydroponic', 'flood', 'active', '550e8400-e29b-41d4-a716-446655440002'),
('770e8400-e29b-41d4-a716-446655440005', '660e8400-e29b-41d4-a716-446655440003', 'Khu C1-1', 'C1-1', 120.25, 240, 'loam', 'manual', 'inactive', '550e8400-e29b-41d4-a716-446655440003');

-- Insert sample greenhouse installation logs
INSERT INTO greenhouse_installation_logs (id, greenhouse_id, action, action_date, description, performed_by) VALUES
('880e8400-e29b-41d4-a716-446655440001', '660e8400-e29b-41d4-a716-446655440001', 'install', '2024-01-15', 'Cài đặt nhà lưới A1 hoàn tất', '550e8400-e29b-41d4-a716-446655440001'),
('880e8400-e29b-41d4-a716-446655440002', '660e8400-e29b-41d4-a716-446655440002', 'install', '2024-02-20', 'Cài đặt nhà lưới B1 hoàn tất', '550e8400-e29b-41d4-a716-446655440001'),
('880e8400-e29b-41d4-a716-446655440003', '660e8400-e29b-41d4-a716-446655440003', 'install', '2024-03-10', 'Cài đặt nhà lưới C1 hoàn tất', '550e8400-e29b-41d4-a716-446655440001'),
('880e8400-e29b-41d4-a716-446655440004', '660e8400-e29b-41d4-a716-446655440003', 'maintenance', '2024-03-25', 'Bảo trì hệ thống tưới', '550e8400-e29b-41d4-a716-446655440002');

-- Insert sample growing zone history
INSERT INTO growing_zone_history (id, zone_id, action, old_value, new_value, action_date, performed_by, notes) VALUES
('990e8400-e29b-41d4-a716-446655440001', '770e8400-e29b-41d4-a716-446655440001', 'change_soil', '{"soil_type": "clay"}', '{"soil_type": "loam"}', '2024-01-20 10:30:00', '550e8400-e29b-41d4-a716-446655440002', 'Thay đổi loại đất từ clay sang loam'),
('990e8400-e29b-41d4-a716-446655440002', '770e8400-e29b-41d4-a716-446655440002', 'change_irrigation', '{"irrigation_system": "manual"}', '{"irrigation_system": "spray"}', '2024-02-15 14:20:00', '550e8400-e29b-41d4-a716-446655440002', 'Nâng cấp hệ thống tưới từ manual sang spray'),
('990e8400-e29b-41d4-a716-446655440003', '770e8400-e29b-41d4-a716-446655440003', 'resize', '{"area_m2": 150.00, "max_plants": 300}', '{"area_m2": 200.00, "max_plants": 400}', '2024-03-05 09:15:00', '550e8400-e29b-41d4-a716-446655440003', 'Mở rộng khu vực trồng');

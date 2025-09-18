-- Migration: Remove sample data
-- Description: Xóa dữ liệu mẫu

-- Delete sample data in reverse order (respecting foreign key constraints)
DELETE FROM growing_zone_history WHERE id IN (
    '990e8400-e29b-41d4-a716-446655440001',
    '990e8400-e29b-41d4-a716-446655440002',
    '990e8400-e29b-41d4-a716-446655440003'
);

DELETE FROM greenhouse_installation_logs WHERE id IN (
    '880e8400-e29b-41d4-a716-446655440001',
    '880e8400-e29b-41d4-a716-446655440002',
    '880e8400-e29b-41d4-a716-446655440003',
    '880e8400-e29b-41d4-a716-446655440004'
);

DELETE FROM growing_zones WHERE id IN (
    '770e8400-e29b-41d4-a716-446655440001',
    '770e8400-e29b-41d4-a716-446655440002',
    '770e8400-e29b-41d4-a716-446655440003',
    '770e8400-e29b-41d4-a716-446655440004',
    '770e8400-e29b-41d4-a716-446655440005'
);

DELETE FROM greenhouses WHERE id IN (
    '660e8400-e29b-41d4-a716-446655440001',
    '660e8400-e29b-41d4-a716-446655440002',
    '660e8400-e29b-41d4-a716-446655440003'
);
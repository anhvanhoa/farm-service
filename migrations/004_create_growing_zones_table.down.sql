-- Migration: Drop growing_zones table
-- Description: Xóa bảng growing_zones

DROP TRIGGER IF EXISTS update_growing_zones_updated_at ON growing_zones;
DROP TABLE IF EXISTS growing_zones;

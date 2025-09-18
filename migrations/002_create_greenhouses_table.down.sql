-- Migration: Drop greenhouses table
-- Description: Xóa bảng greenhouses

DROP TRIGGER IF EXISTS update_greenhouses_updated_at ON greenhouses;
DROP TABLE IF EXISTS greenhouses;

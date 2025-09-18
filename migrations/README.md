# Database Migrations

## Cấu trúc thư mục

```
migrations/
├── 001_create_users_table.up.sql
├── 001_create_users_table.down.sql
├── 002_create_greenhouses_table.up.sql
├── 002_create_greenhouses_table.down.sql
├── 003_create_greenhouse_installation_logs_table.up.sql
├── 003_create_greenhouse_installation_logs_table.down.sql
├── 004_create_growing_zones_table.up.sql
├── 004_create_growing_zones_table.down.sql
├── 005_create_growing_zone_history_table.up.sql
├── 005_create_growing_zone_history_table.down.sql
└── README.md
```

## Mô tả các migration

### 001 - Users Table
- **Mục đích**: Tạo bảng users để quản lý người dùng
- **Lý do**: Cần thiết vì các bảng khác có foreign key reference đến users.id
- **Các trường chính**:
  - `id`: UUID primary key
  - `username`: Tên đăng nhập (unique)
  - `email`: Email (unique)
  - `full_name`: Tên đầy đủ
  - `role`: Vai trò (admin, manager, user)
  - `status`: Trạng thái (active, inactive, suspended)

### 002 - Greenhouses Table
- **Mục đích**: Tạo bảng quản lý nhà lưới
- **Các trường chính**:
  - `id`: UUID primary key
  - `name`: Tên nhà lưới
  - `location`: Vị trí
  - `area_m2`: Diện tích (m²)
  - `type`: Loại nhà lưới (glass, plastic, tunnel, etc.)
  - `max_capacity`: Số lượng cây tối đa
  - `installation_date`: Ngày cài đặt
  - `status`: Trạng thái (active, inactive, maintenance)
  - `description`: Mô tả
  - `created_by`: Người tạo (FK to users)

### 003 - Greenhouse Installation Logs Table
- **Mục đích**: Tạo bảng log cài đặt nhà lưới
- **Các trường chính**:
  - `id`: UUID primary key
  - `greenhouse_id`: ID nhà lưới (FK to greenhouses)
  - `action`: Hành động (install, upgrade, maintenance, relocate, dismantle)
  - `action_date`: Ngày thực hiện
  - `description`: Mô tả
  - `performed_by`: Người thực hiện (FK to users)

### 004 - Growing Zones Table
- **Mục đích**: Tạo bảng khu vực trồng
- **Các trường chính**:
  - `id`: UUID primary key
  - `greenhouse_id`: ID nhà lưới (FK to greenhouses)
  - `zone_name`: Tên khu vực
  - `zone_code`: Mã khu vực (unique)
  - `area_m2`: Diện tích (m²)
  - `max_plants`: Số cây tối đa
  - `soil_type`: Loại đất (sandy, clay, loam, hydroponic)
  - `irrigation_system`: Hệ thống tưới (drip, spray, flood, manual)
  - `status`: Trạng thái
  - `created_by`: Người tạo (FK to users)

### 005 - Growing Zone History Table
- **Mục đích**: Tạo bảng lịch sử thay đổi khu vực trồng
- **Các trường chính**:
  - `id`: UUID primary key
  - `zone_id`: ID khu vực (FK to growing_zones)
  - `action`: Hành động (change_soil, change_irrigation, maintenance, resize, rename)
  - `old_value`: Giá trị cũ (JSON)
  - `new_value`: Giá trị mới (JSON)
  - `action_date`: Ngày thực hiện
  - `performed_by`: Người thực hiện (FK to users)
  - `notes`: Ghi chú

## Cách sử dụng

### Chạy migration lên (up)
```bash
# Chạy tất cả migration
migrate -path migrations -database "postgres://postgres:password@localhost:5432/farm_service?sslmode=disable" up

# Chạy migration cụ thể
migrate -path migrations -database "postgres://postgres:password@localhost:5432/farm_service?sslmode=disable" goto 005
```

### Rollback migration (down)
```bash
# Rollback 1 bước
migrate -path migrations -database "postgres://postgres:password@localhost:5432/farm_service?sslmode=disable" down 1

# Rollback tất cả
migrate -path migrations -database "postgres://postgres:password@localhost:5432/farm_service?sslmode=disable" down
```

### Kiểm tra trạng thái migration
```bash
migrate -path migrations -database "postgres://postgres:password@localhost:5432/farm_service?sslmode=disable" version
```

## Lưu ý

1. **Thứ tự migration**: Phải chạy theo thứ tự số để đảm bảo foreign key constraints
2. **Foreign Key Dependencies**:
   - `greenhouses.created_by` → `users.id`
   - `greenhouse_installation_logs.greenhouse_id` → `greenhouses.id`
   - `greenhouse_installation_logs.performed_by` → `users.id`
   - `growing_zones.greenhouse_id` → `greenhouses.id`
   - `growing_zones.created_by` → `users.id`
   - `growing_zone_history.zone_id` → `growing_zones.id`
   - `growing_zone_history.performed_by` → `users.id`

3. **Indexes**: Đã tạo các index cần thiết cho performance
4. **Cascade Delete**: Các bảng con sẽ tự động xóa khi bảng cha bị xóa
5. **UUID**: Sử dụng UUID cho primary key để tránh conflict khi scale
6. **JSONB Fields**: Sử dụng JSONB cho old_value và new_value trong history table (PostgreSQL)
7. **Triggers**: Tự động cập nhật updated_at timestamp
8. **GIN Indexes**: Tối ưu cho JSONB fields

## Database Schema Diagram

```
users (1) ──→ (N) greenhouses
users (1) ──→ (N) greenhouse_installation_logs
users (1) ──→ (N) growing_zones
users (1) ──→ (N) growing_zone_history

greenhouses (1) ──→ (N) greenhouse_installation_logs
greenhouses (1) ──→ (N) growing_zones

growing_zones (1) ──→ (N) growing_zone_history
```

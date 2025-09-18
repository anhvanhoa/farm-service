# Farm Service

Microservice quản lý nhà lưới và khu vực trồng trọt trong hệ thống nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   ├── repository/   # Giao diện truy cập dữ liệu
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Nhà lưới
- ✅ Tạo, đọc, cập nhật, xóa nhà lưới
- ✅ Liệt kê nhà lưới với bộ lọc (trạng thái, loại, vị trí)
- ✅ Hỗ trợ phân trang
- ✅ Ghi log cài đặt
- ✅ Xác thực dữ liệu

### Quản lý Khu vực Trồng
- ✅ Tạo, đọc, cập nhật, xóa khu vực trồng
- ✅ Liệt kê khu vực với bộ lọc (greenhouse_id, trạng thái, loại đất, hệ thống tưới)
- ✅ Hỗ trợ phân trang
- ✅ Xác thực mã khu vực duy nhất
- ✅ Lấy tất cả khu vực theo nhà lưới
- ✅ Theo dõi lịch sử thay đổi khu vực

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.24.6
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc

## 📋 Yêu cầu hệ thống

- Go 1.24.6 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd farm-service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://postgres:123456@localhost:5432/farm_service_db?sslmode=disable"
name_service: "FarmService"
port_grpc: 50054
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force version=1
```

## 📁 Cấu trúc Dự án

```
farm-service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   ├── greenhouse.go   # Entity nhà lưới và DTOs
│   │   └── growing_zone.go # Entity khu vực trồng và DTOs
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   ├── greenhouse_repository.go
│   │   └── growing_zone_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       ├── greenhouse/     # Use cases nhà lưới
│       └── growing_zone/   # Use cases khu vực trồng
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── greenhouse/     # gRPC handlers nhà lưới
│   │   ├── growing_zone/   # gRPC handlers khu vực trồng
│   │   └── server.go       # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── postgres_greenhouse_repository.go
│       ├── postgres_growing_zone_repository.go
│       └── repositories.go
├── migrations/              # Database migrations
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Nhà lưới (Greenhouse)
- **ID**: Định danh duy nhất
- **Name**: Tên nhà lưới
- **Location**: Vị trí vật lý
- **AreaM2**: Diện tích tính bằng mét vuông
- **Type**: Loại nhà lưới
- **MaxCapacity**: Sức chứa tối đa
- **InstallationDate**: Ngày cài đặt
- **Status**: Trạng thái hiện tại (active, inactive, maintenance)
- **Description**: Mô tả bổ sung
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

### Khu vực Trồng (Growing Zone)
- **ID**: Định danh duy nhất
- **GreenhouseID**: Tham chiếu nhà lưới cha
- **ZoneName**: Tên khu vực
- **ZoneCode**: Mã khu vực duy nhất
- **AreaM2**: Diện tích khu vực tính bằng mét vuông
- **MaxPlants**: Sức chứa cây trồng tối đa
- **SoilType**: Loại đất
- **IrrigationSystem**: Loại hệ thống tưới
- **Status**: Trạng thái hiện tại
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Greenhouse Service
- `CreateGreenhouse` - Tạo nhà lưới mới
- `GetGreenhouse` - Lấy thông tin nhà lưới theo ID
- `UpdateGreenhouse` - Cập nhật thông tin nhà lưới
- `DeleteGreenhouse` - Xóa nhà lưới
- `ListGreenhouses` - Liệt kê nhà lưới với bộ lọc

### Growing Zone Service
- `CreateGrowingZone` - Tạo khu vực trồng mới
- `GetGrowingZone` - Lấy thông tin khu vực trồng theo ID
- `UpdateGrowingZone` - Cập nhật thông tin khu vực trồng
- `DeleteGrowingZone` - Xóa khu vực trồng
- `ListGrowingZones` - Liệt kê khu vực trồng với bộ lọc
- `GetZonesByGreenhouse` - Lấy tất cả khu vực của một nhà lưới

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để là một phần của hệ thống quản lý nông nghiệp lớn hơn và tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng.

# Domain Layer - Farm Service

## Cấu trúc thư mục

```
domain/
├── common/           # Các struct và interface chung
│   └── response.go   # BaseResponse, PaginationRequest, PaginationResponse
├── entity/           # Các thực thể (entities)
│   ├── greenhouse.go # Entity và request/response cho Greenhouse
│   └── growing_zone.go # Entity và request/response cho GrowingZone
├── repository/       # Interface repository
│   ├── greenhouse_repository.go
│   └── growing_zone_repository.go
└── usecase/          # Interface use case
    ├── greenhouse/   # Use cases cho Greenhouse
    │   ├── create_greenhouse_usecase.go
    │   ├── get_greenhouse_usecase.go
    │   ├── update_greenhouse_usecase.go
    │   ├── delete_greenhouse_usecase.go
    │   └── list_greenhouse_usecase.go
    └── growing_zone/ # Use cases cho GrowingZone
        ├── create_growing_zone_usecase.go
        ├── get_growing_zone_usecase.go
        ├── update_growing_zone_usecase.go
        ├── delete_growing_zone_usecase.go
        ├── list_growing_zone_usecase.go
        └── get_zones_by_greenhouse_usecase.go
```

## Mô tả các thành phần

### 1. Common
- **BaseResponse**: Struct chung cho tất cả response API
- **PaginationRequest**: Struct cho request phân trang
- **PaginationResponse**: Struct cho response phân trang
- **Status constants**: Các hằng số trạng thái chung

### 2. Entity
#### Greenhouse
- **Greenhouse**: Entity chính cho nhà lưới
- **GreenhouseInstallationLog**: Log cài đặt nhà lưới
- **CreateGreenhouseRequest**: Request tạo nhà lưới mới
- **UpdateGreenhouseRequest**: Request cập nhật nhà lưới
- **GreenhouseFilter**: Filter tìm kiếm nhà lưới

#### GrowingZone
- **GrowingZone**: Entity chính cho khu vực trồng
- **GrowingZoneHistory**: Lịch sử thay đổi khu vực trồng
- **CreateGrowingZoneRequest**: Request tạo khu vực trồng mới
- **UpdateGrowingZoneRequest**: Request cập nhật khu vực trồng
- **GrowingZoneFilter**: Filter tìm kiếm khu vực trồng

### 3. Repository Interface
#### GreenhouseRepository
- `Create()`: Tạo nhà lưới mới
- `GetByID()`: Lấy nhà lưới theo ID
- `Update()`: Cập nhật nhà lưới
- `Delete()`: Xóa nhà lưới
- `List()`: Lấy danh sách với filter và phân trang
- `GetByStatus()`: Lấy theo trạng thái
- `GetByLocation()`: Lấy theo vị trí
- `Count()`: Đếm số lượng

#### GrowingZoneRepository
- `Create()`: Tạo khu vực trồng mới
- `GetByID()`: Lấy khu vực trồng theo ID
- `GetByZoneCode()`: Lấy theo mã zone code
- `Update()`: Cập nhật khu vực trồng
- `Delete()`: Xóa khu vực trồng
- `List()`: Lấy danh sách với filter và phân trang
- `GetByGreenhouseID()`: Lấy tất cả khu vực của một nhà lưới
- `CheckZoneCodeExists()`: Kiểm tra zone code đã tồn tại

### 4. Use Case
#### Greenhouse Use Cases
- **CreateGreenhouseUsecase**: Tạo nhà lưới mới
- **GetGreenhouseUsecase**: Lấy thông tin nhà lưới
- **UpdateGreenhouseUsecase**: Cập nhật nhà lưới
- **DeleteGreenhouseUsecase**: Xóa nhà lưới
- **ListGreenhouseUsecase**: Lấy danh sách nhà lưới

#### GrowingZone Use Cases
- **CreateGrowingZoneUsecase**: Tạo khu vực trồng mới
- **GetGrowingZoneUsecase**: Lấy thông tin khu vực trồng
- **UpdateGrowingZoneUsecase**: Cập nhật khu vực trồng
- **DeleteGrowingZoneUsecase**: Xóa khu vực trồng
- **ListGrowingZoneUsecase**: Lấy danh sách khu vực trồng
- **GetZonesByGreenhouseUsecase**: Lấy tất cả khu vực của một nhà lưới

## Cách sử dụng

### 1. Tạo Use Case
```go
// Tạo repository (implement từ infrastructure layer)
greenhouseRepo := infrastructure.NewGreenhouseRepository(db)

// Tạo use case
createGreenhouseUsecase := greenhouse.NewCreateGreenhouseUsecase(greenhouseRepo)
```

### 2. Sử dụng Use Case
```go
// Tạo request
req := &entity.CreateGreenhouseRequest{
    Name:        "Nhà lưới A",
    Location:    "Hà Nội",
    AreaM2:      100.5,
    Type:        "plastic",
    MaxCapacity: 1000,
    CreatedBy:   "user123",
}

// Thực thi use case
greenhouse, err := createGreenhouseUsecase.Execute(ctx, req)
```

## Lưu ý
- Tất cả use case đều nhận `context.Context` làm tham số đầu tiên
- Các repository interface sẽ được implement trong infrastructure layer
- Các use case interface sẽ được implement trong infrastructure layer (gRPC service)
- Sử dụng Clean Architecture pattern để tách biệt business logic và infrastructure
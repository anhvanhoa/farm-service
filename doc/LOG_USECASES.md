# Log Usecases Documentation

## Tổng quan

Tài liệu này mô tả các usecase đã được tạo cho 2 bảng log trong hệ thống:
1. **Greenhouse Installation Logs** - Log cài đặt nhà lưới
2. **Growing Zone History** - Lịch sử thay đổi khu vực trồng

## 1. Greenhouse Installation Logs

### Cấu trúc dữ liệu
- **Bảng**: `greenhouse_installation_logs`
- **Entity**: `GreenhouseInstallationLog`
- **Repository**: `GreenhouseInstallationLogRepository`

### Các Usecase đã tạo

#### 1.1 CreateLogUseCase
- **Mục đích**: Tạo log mới cho hoạt động cài đặt nhà lưới
- **Input**: 
  - `greenhouse_id`: ID của nhà lưới
  - `action`: Loại hành động (install, upgrade, maintenance, relocate, dismantle)
  - `action_date`: Ngày thực hiện (YYYY-MM-DD)
  - `description`: Mô tả chi tiết
  - `performed_by`: ID người thực hiện
- **Validation**: 
  - Kiểm tra greenhouse tồn tại
  - Validate format ngày tháng
  - Validate action type
- **Output**: `GreenhouseInstallationLog` object

#### 1.2 GetLogsByGreenhouseUseCase
- **Mục đích**: Lấy tất cả logs của một nhà lưới cụ thể
- **Input**: `greenhouse_id`
- **Validation**: Kiểm tra greenhouse tồn tại
- **Output**: Danh sách `GreenhouseInstallationLog` được sắp xếp theo ngày giảm dần

#### 1.3 GetLogsByActionUseCase
- **Mục đích**: Lấy logs theo loại hành động
- **Input**: `action` (install, upgrade, maintenance, relocate, dismantle)
- **Output**: Danh sách `GreenhouseInstallationLog` được sắp xếp theo ngày giảm dần

#### 1.4 GetLogsByDateRangeUseCase
- **Mục đích**: Lấy logs trong khoảng thời gian
- **Input**: 
  - `start_date`: Ngày bắt đầu (YYYY-MM-DD)
  - `end_date`: Ngày kết thúc (YYYY-MM-DD)
- **Validation**: Validate format ngày tháng
- **Output**: Danh sách `GreenhouseInstallationLog` được sắp xếp theo ngày giảm dần

## 2. Growing Zone History

### Cấu trúc dữ liệu
- **Bảng**: `growing_zone_history`
- **Entity**: `GrowingZoneHistory`
- **Repository**: `GrowingZoneHistoryRepository`

### Các Usecase đã tạo

#### 2.1 CreateHistoryUseCase
- **Mục đích**: Tạo record lịch sử mới cho khu vực trồng
- **Input**:
  - `zone_id`: ID của khu vực trồng
  - `action`: Loại hành động (change_soil, change_irrigation, maintenance, resize, rename)
  - `old_value`: Giá trị cũ (JSON object)
  - `new_value`: Giá trị mới (JSON object)
  - `performed_by`: ID người thực hiện
  - `notes`: Ghi chú
- **Validation**: Kiểm tra growing zone tồn tại
- **Output**: `GrowingZoneHistory` object

#### 2.2 GetHistoryByZoneUseCase
- **Mục đích**: Lấy tất cả lịch sử của một khu vực trồng
- **Input**: `zone_id`
- **Validation**: Kiểm tra growing zone tồn tại
- **Output**: Danh sách `GrowingZoneHistory` được sắp xếp theo ngày giảm dần

#### 2.3 GetHistoryByActionUseCase
- **Mục đích**: Lấy lịch sử theo loại hành động
- **Input**: `action` (change_soil, change_irrigation, maintenance, resize, rename)
- **Output**: Danh sách `GrowingZoneHistory` được sắp xếp theo ngày giảm dần

#### 2.4 GetHistoryByDateRangeUseCase
- **Mục đích**: Lấy lịch sử trong khoảng thời gian
- **Input**:
  - `start_date`: Ngày bắt đầu (YYYY-MM-DD)
  - `end_date`: Ngày kết thúc (YYYY-MM-DD)
- **Validation**: Validate format ngày tháng
- **Output**: Danh sách `GrowingZoneHistory` được sắp xếp theo ngày giảm dần

#### 2.5 GetHistoryByPerformedByUseCase
- **Mục đích**: Lấy lịch sử theo người thực hiện
- **Input**: `performed_by` (ID người thực hiện)
- **Output**: Danh sách `GrowingZoneHistory` được sắp xếp theo ngày giảm dần

## 3. Cấu trúc thư mục

```
domain/usecase/
├── greenhouse_installation_log/
│   ├── base.go
│   ├── create_log_usecase.go
│   ├── get_logs_by_greenhouse_usecase.go
│   ├── get_logs_by_action_usecase.go
│   └── get_logs_by_date_range_usecase.go
└── growing_zone_history/
    ├── base.go
    ├── create_history_usecase.go
    ├── get_history_by_zone_usecase.go
    ├── get_history_by_action_usecase.go
    ├── get_history_by_date_range_usecase.go
    └── get_history_by_performed_by_usecase.go
```

## 4. Cách sử dụng

### Khởi tạo Usecase

```go
// Tạo base usecase cho greenhouse installation logs
baseLogUseCase := greenhouse_installation_log.NewBaseUseCase(
    greenhouseRepo,
    logRepo,
)

// Tạo các usecase cụ thể
createLogUseCase := greenhouse_installation_log.NewCreateLogUseCase(baseLogUseCase)
getLogsByGreenhouseUseCase := greenhouse_installation_log.NewGetLogsByGreenhouseUseCase(baseLogUseCase)

// Tạo base usecase cho growing zone history
baseHistoryUseCase := growing_zone_history.NewBaseUseCase(
    growingZoneRepo,
    historyRepo,
)

// Tạo các usecase cụ thể
createHistoryUseCase := growing_zone_history.NewCreateHistoryUseCase(baseHistoryUseCase)
getHistoryByZoneUseCase := growing_zone_history.NewGetHistoryByZoneUseCase(baseHistoryUseCase)
```

### Ví dụ sử dụng

```go
// Tạo log cài đặt nhà lưới
logReq := &greenhouse_installation_log.CreateLogRequest{
    GreenhouseID: "greenhouse-uuid",
    Action:       "install",
    ActionDate:   "2024-01-15",
    Description:  "Cài đặt nhà lưới mới",
    PerformedBy:  "user-uuid",
}

log, err := createLogUseCase.Execute(ctx, logReq)

// Tạo history cho khu vực trồng
historyReq := &growing_zone_history.CreateHistoryRequest{
    ZoneID:      "zone-uuid",
    Action:      "change_soil",
    OldValue:    map[string]interface{}{"soil_type": "clay"},
    NewValue:    map[string]interface{}{"soil_type": "sandy"},
    PerformedBy: "user-uuid",
    Notes:       "Thay đổi loại đất",
}

history, err := createHistoryUseCase.Execute(ctx, historyReq)
```

## 5. Lưu ý

- Tất cả các usecase đều có validation đầu vào
- Các lỗi được trả về dưới dạng `entity.Error` với code và message cụ thể
- Tất cả các query đều được sắp xếp theo thời gian giảm dần (mới nhất trước)
- Các repository đã được đăng ký trong `infrastructure/repo/repositories.go`
- Code đã được build thành công và không có lỗi syntax

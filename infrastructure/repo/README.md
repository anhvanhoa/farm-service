# PostgreSQL Repository Implementations

This package contains PostgreSQL implementations of the repository interfaces using `github.com/go-pg/pg/v10`.

## Features

- Full CRUD operations for all entities
- Pagination support
- Filtering capabilities
- Context-aware operations
- Error handling with proper nil checks
- Connection management

## Available Repositories

### 1. GreenhouseRepository
- `Create(ctx, greenhouse)` - Create a new greenhouse
- `GetByID(ctx, id)` - Get greenhouse by ID
- `GetByCode(ctx, code)` - Get greenhouse by code
- `Update(ctx, id, updateReq)` - Update greenhouse
- `Delete(ctx, id)` - Delete greenhouse
- `List(ctx, filter, page, pageSize)` - List greenhouses with pagination and filtering
- `GetByStatus(ctx, status)` - Get greenhouses by status
- `GetByLocation(ctx, location)` - Get greenhouses by location
- `Count(ctx, filter)` - Count greenhouses with filters

### 2. GreenhouseInstallationLogRepository
- `Create(ctx, log)` - Create installation log
- `GetByGreenhouseID(ctx, greenhouseID)` - Get logs by greenhouse ID
- `GetByAction(ctx, action)` - Get logs by action
- `GetByDateRange(ctx, startDate, endDate)` - Get logs by date range

### 3. GrowingZoneRepository
- `Create(ctx, zone)` - Create a new growing zone
- `GetByID(ctx, id)` - Get zone by ID
- `GetByZoneCode(ctx, zoneCode)` - Get zone by zone code
- `Update(ctx, id, updateReq)` - Update zone
- `Delete(ctx, id)` - Delete zone
- `List(ctx, filter, page, pageSize)` - List zones with pagination and filtering
- `GetByGreenhouseID(ctx, greenhouseID)` - Get zones by greenhouse ID
- `GetByStatus(ctx, status)` - Get zones by status
- `GetBySoilType(ctx, soilType)` - Get zones by soil type
- `GetByIrrigationSystem(ctx, irrigationSystem)` - Get zones by irrigation system
- `Count(ctx, filter)` - Count zones with filters
- `CheckZoneCodeExists(ctx, zoneCode)` - Check if zone code exists

### 4. GrowingZoneHistoryRepository
- `Create(ctx, history)` - Create history record
- `GetByZoneID(ctx, zoneID)` - Get history by zone ID
- `GetByAction(ctx, action)` - Get history by action
- `GetByDateRange(ctx, startDate, endDate)` - Get history by date range
- `GetByPerformedBy(ctx, performedBy)` - Get history by performer

## Usage

### 1. Database Connection

```go
import "farm-service/infrastructure/repo"

// Create database connection
db, err := repo.NewDatabaseConnection()
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}
defer db.Close()
```

### 2. Initialize Repositories

```go
// Create all repositories
repos := repo.NewRepositories(db)

// Access individual repositories
greenhouseRepo := repos.GreenhouseRepository
zoneRepo := repos.GrowingZoneRepository
```

### 3. Environment Variables

Set the following environment variables for database connection:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=farm_service
DB_SSLMODE=disable
```

### 4. Example Operations

```go
ctx := context.Background()

// Create a greenhouse
greenhouse := &entity.Greenhouse{
    ID:          uuid.New().String(),
    Name:        "Greenhouse A",
    Location:    "Farm 1",
    AreaM2:      100.5,
    Type:        "Glass",
    MaxCapacity: 500,
    Status:      entity.StatusActive,
    CreatedBy:   "admin",
}

err := repos.GreenhouseRepository.Create(ctx, greenhouse)

// Get greenhouse by ID
retrieved, err := repos.GreenhouseRepository.GetByID(ctx, greenhouse.ID)

// List with pagination and filtering
greenhouses, total, err := repos.GreenhouseRepository.List(ctx, &entity.GreenhouseFilter{
    Status: entity.StatusActive,
}, 1, 10)
```

## Database Schema Requirements

Make sure your PostgreSQL database has the following tables:
- `greenhouses`
- `greenhouse_installation_logs`
- `growing_zones`
- `growing_zone_history`

The table schemas should match the entity structs defined in the `domain/entity` package.

## Dependencies

- `github.com/go-pg/pg/v10` - PostgreSQL ORM
- `github.com/google/uuid` - UUID generation
- Standard Go packages: `context`, `fmt`, `os`

## Error Handling

All repository methods return errors that should be handled appropriately:
- `pg.ErrNoRows` is returned when no records are found (converted to `nil` for single record queries)
- Database connection errors
- Constraint violation errors
- Other database-related errors

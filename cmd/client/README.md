# Farm Service Test Client

This is a gRPC test client for the Farm Service, based on the exam.go template. It provides an interactive command-line interface to test all the gRPC services.

## Features

### Greenhouse Service
- Create Greenhouse
- Get Greenhouse by ID
- List Greenhouses (with pagination and filtering)
- Update Greenhouse
- Delete Greenhouse

### Growing Zone Service
- Create Growing Zone
- Get Growing Zone by ID
- List Growing Zones (with pagination and filtering)
- Get Zones by Greenhouse ID
- Update Growing Zone
- Delete Growing Zone

## Usage

### Prerequisites
1. Make sure the Farm Service gRPC server is running
2. Ensure the `dev.config.yml` file is properly configured

### Running the Test Client

```bash
# From the project root directory
cd cmd/client
go run main.go

# Or specify a custom server address
go run main.go localhost:50054
```

### Configuration

The client reads configuration from `dev.config.yml` in the project root:
- `host_grpc`: Server host (default: localhost)
- `port_grpc`: Server port (default: 50054)

## Example Usage

1. Start the Farm Service server
2. Run the test client
3. Select a service (Greenhouse or Growing Zone)
4. Choose an operation (Create, Get, List, Update, Delete)
5. Follow the prompts to enter required data
6. View the results

## Menu Structure

```
Main Menu
├── 1. Greenhouse Service
│   ├── 1. Create Greenhouse
│   ├── 2. Get Greenhouse
│   ├── 3. List Greenhouses
│   ├── 4. Update Greenhouse
│   └── 5. Delete Greenhouse
└── 2. Growing Zone Service
    ├── 1. Create Growing Zone
    ├── 2. Get Growing Zone
    ├── 3. List Growing Zones
    ├── 4. Get Zones By Greenhouse
    ├── 5. Update Growing Zone
    └── 6. Delete Growing Zone
```

## Notes

- All input is cleaned and validated
- Default values are provided for optional fields
- The client handles errors gracefully
- Timeout is set to 10 seconds for all requests
- The client automatically connects to the gRPC server on startup

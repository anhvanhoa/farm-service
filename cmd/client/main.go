package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_greenhouse "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	proto_greenhouse_installation_log "github.com/anhvanhoa/sf-proto/gen/greenhouse_installation_log/v1"
	proto_growing_zone "github.com/anhvanhoa/sf-proto/gen/growing_zone/v1"
	proto_growing_zone_history "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/structpb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

type FarmServiceClient struct {
	greenhouseClient                proto_greenhouse.GreenhouseServiceClient
	greenhouseInstallationLogClient proto_greenhouse_installation_log.GreenhouseInstallationLogServiceClient
	growingZoneClient               proto_growing_zone.GrowingZoneServiceClient
	growingZoneHistoryClient        proto_growing_zone_history.GrowingZoneHistoryServiceClient
	conn                            *grpc.ClientConn
}

func NewFarmServiceClient(address string) (*FarmServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &FarmServiceClient{
		greenhouseClient:                proto_greenhouse.NewGreenhouseServiceClient(conn),
		greenhouseInstallationLogClient: proto_greenhouse_installation_log.NewGreenhouseInstallationLogServiceClient(conn),
		growingZoneClient:               proto_growing_zone.NewGrowingZoneServiceClient(conn),
		growingZoneHistoryClient:        proto_growing_zone_history.NewGrowingZoneHistoryServiceClient(conn),
		conn:                            conn,
	}, nil
}

func (c *FarmServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Greenhouse Service Tests ==================

func (c *FarmServiceClient) TestCreateGreenhouse() {
	fmt.Println("\n=== Test Create Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter location: ")
	location, _ := reader.ReadString('\n')
	location = cleanInput(location)

	fmt.Print("Enter area (m²): ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float32(100.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 32); err == nil {
			area = float32(a)
		}
	}

	fmt.Print("Enter greenhouse type: ")
	greenhouseType, _ := reader.ReadString('\n')
	greenhouseType = cleanInput(greenhouseType)

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseClient.CreateGreenhouse(ctx, &proto_greenhouse.CreateGreenhouseRequest{
		Name:        name,
		Location:    location,
		AreaM2:      float64(area),
		Type:        greenhouseType,
		MaxCapacity: 100,
		Description: "Test greenhouse",
		CreatedBy:   createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Create Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.Greenhouse.Id)
	fmt.Printf("Name: %s\n", resp.Greenhouse.Name)
	fmt.Printf("Location: %s\n", resp.Greenhouse.Location)
	fmt.Printf("Area: %.2f m²\n", resp.Greenhouse.AreaM2)
	fmt.Printf("Type: %s\n", resp.Greenhouse.Type)
	fmt.Printf("Max Capacity: %d\n", resp.Greenhouse.MaxCapacity)
}

func (c *FarmServiceClient) TestGetGreenhouse() {
	fmt.Println("\n=== Test Get Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseClient.GetGreenhouse(ctx, &proto_greenhouse.GetGreenhouseRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Get Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.Greenhouse.Id)
	fmt.Printf("Name: %s\n", resp.Greenhouse.Name)
	fmt.Printf("Location: %s\n", resp.Greenhouse.Location)
	fmt.Printf("Area: %.2f m²\n", resp.Greenhouse.AreaM2)
	fmt.Printf("Type: %s\n", resp.Greenhouse.Type)
	fmt.Printf("Max Capacity: %d\n", resp.Greenhouse.MaxCapacity)
}

func (c *FarmServiceClient) TestListGreenhouses() {
	fmt.Println("\n=== Test List Greenhouses ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	fmt.Print("Enter search term (optional): ")
	search, _ := reader.ReadString('\n')
	search = cleanInput(search)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseClient.ListGreenhouses(ctx, &proto_greenhouse.ListGreenhousesRequest{
		Page:     page,
		PageSize: limit,
		Filter: &proto_greenhouse.GreenhouseFilter{
			Status:   "",
			Type:     "",
			Location: search,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListGreenhouses: %v\n", err)
		return
	}

	fmt.Printf("List Greenhouses result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.TotalCount)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Page Size: %d\n", resp.PageSize)
	fmt.Printf("Greenhouses:\n")
	for i, greenhouse := range resp.Greenhouses {
		fmt.Printf("  [%d] ID: %s, Name: %s, Location: %s, Area: %.2f m², Type: %s\n",
			i+1, greenhouse.Id, greenhouse.Name, greenhouse.Location, greenhouse.AreaM2, greenhouse.Type)
	}
}

func (c *FarmServiceClient) TestUpdateGreenhouse() {
	fmt.Println("\n=== Test Update Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new location: ")
	location, _ := reader.ReadString('\n')
	location = cleanInput(location)

	fmt.Print("Enter new area (m²): ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float32(100.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 32); err == nil {
			area = float32(a)
		}
	}

	fmt.Print("Enter new greenhouse type: ")
	greenhouseType, _ := reader.ReadString('\n')
	greenhouseType = cleanInput(greenhouseType)

	// fmt.Print("Enter new status: ")
	// status, _ := reader.ReadString('\n')
	// status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseClient.UpdateGreenhouse(ctx, &proto_greenhouse.UpdateGreenhouseRequest{
		Id:          id,
		Name:        name,
		Location:    location,
		AreaM2:      float64(area),
		Type:        greenhouseType,
		MaxCapacity: 100,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Update Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
}

func (c *FarmServiceClient) TestDeleteGreenhouse() {
	fmt.Println("\n=== Test Delete Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseClient.DeleteGreenhouse(ctx, &proto_greenhouse.DeleteGreenhouseRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Delete Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Growing Zone Service Tests ==================

func (c *FarmServiceClient) TestCreateGrowingZone() {
	fmt.Println("\n=== Test Create Growing Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	fmt.Print("Enter zone name: ")
	zoneName, _ := reader.ReadString('\n')
	zoneName = cleanInput(zoneName)

	fmt.Print("Enter zone code: ")
	zoneCode, _ := reader.ReadString('\n')
	zoneCode = cleanInput(zoneCode)

	fmt.Print("Enter area (m²): ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float32(50.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 32); err == nil {
			area = float32(a)
		}
	}

	fmt.Print("Enter max plants: ")
	maxPlantsStr, _ := reader.ReadString('\n')
	maxPlantsStr = cleanInput(maxPlantsStr)
	maxPlants := int32(100)
	if maxPlantsStr != "" {
		if mp, err := strconv.Atoi(maxPlantsStr); err == nil {
			maxPlants = int32(mp)
		}
	}

	fmt.Print("Enter soil type: ")
	soilType, _ := reader.ReadString('\n')
	soilType = cleanInput(soilType)

	fmt.Print("Enter irrigation system: ")
	irrigationSystem, _ := reader.ReadString('\n')
	irrigationSystem = cleanInput(irrigationSystem)

	fmt.Print("Enter created by: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.CreateGrowingZone(ctx, &proto_growing_zone.CreateGrowingZoneRequest{
		GreenhouseId:     greenhouseId,
		ZoneName:         zoneName,
		ZoneCode:         zoneCode,
		AreaM2:           float64(area),
		MaxPlants:        maxPlants,
		SoilType:         soilType,
		IrrigationSystem: irrigationSystem,
		CreatedBy:        createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateGrowingZone: %v\n", err)
		return
	}

	fmt.Printf("Create Growing Zone result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("ID: %s\n", resp.GrowingZone.Id)
	fmt.Printf("Greenhouse ID: %s\n", resp.GrowingZone.GreenhouseId)
	fmt.Printf("Zone Name: %s\n", resp.GrowingZone.ZoneName)
	fmt.Printf("Zone Code: %s\n", resp.GrowingZone.ZoneCode)
	fmt.Printf("Area: %.2f m²\n", resp.GrowingZone.AreaM2)
	fmt.Printf("Max Plants: %d\n", resp.GrowingZone.MaxPlants)
	fmt.Printf("Soil Type: %s\n", resp.GrowingZone.SoilType)
	fmt.Printf("Irrigation System: %s\n", resp.GrowingZone.IrrigationSystem)
	fmt.Printf("Status: %s\n", resp.GrowingZone.Status)
	fmt.Printf("Created By: %s\n", resp.GrowingZone.CreatedBy)
}

func (c *FarmServiceClient) TestGetGrowingZone() {
	fmt.Println("\n=== Test Get Growing Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter growing zone ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.GetGrowingZone(ctx, &proto_growing_zone.GetGrowingZoneRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetGrowingZone: %v\n", err)
		return
	}

	fmt.Printf("Get Growing Zone result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	if resp.GrowingZone != nil {
		fmt.Printf("ID: %s\n", resp.GrowingZone.Id)
		fmt.Printf("Greenhouse ID: %s\n", resp.GrowingZone.GreenhouseId)
		fmt.Printf("Zone Name: %s\n", resp.GrowingZone.ZoneName)
		fmt.Printf("Zone Code: %s\n", resp.GrowingZone.ZoneCode)
		fmt.Printf("Area: %.2f m²\n", resp.GrowingZone.AreaM2)
		fmt.Printf("Max Plants: %d\n", resp.GrowingZone.MaxPlants)
		fmt.Printf("Soil Type: %s\n", resp.GrowingZone.SoilType)
		fmt.Printf("Irrigation System: %s\n", resp.GrowingZone.IrrigationSystem)
		fmt.Printf("Status: %s\n", resp.GrowingZone.Status)
		fmt.Printf("Created By: %s\n", resp.GrowingZone.CreatedBy)
	}
}

func (c *FarmServiceClient) TestListGrowingZones() {
	fmt.Println("\n=== Test List Growing Zones ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	// fmt.Print("Enter search term (optional): ")
	// search, _ := reader.ReadString('\n')
	// search = cleanInput(search)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.ListGrowingZones(ctx, &proto_growing_zone.ListGrowingZonesRequest{
		Page:     page,
		PageSize: limit,
		Filter: &proto_growing_zone.GrowingZoneFilter{
			GreenhouseId:     "",
			Status:           "",
			SoilType:         "",
			IrrigationSystem: "",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListGrowingZones: %v\n", err)
		return
	}

	fmt.Printf("List Growing Zones result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Total: %d\n", resp.TotalCount)
	fmt.Printf("Page: %d\n", resp.Page)
	fmt.Printf("Page Size: %d\n", resp.PageSize)
	fmt.Printf("Growing Zones:\n")
	for i, zone := range resp.GrowingZones {
		fmt.Printf("  [%d] ID: %s, Zone Name: %s, Zone Code: %s, Area: %.2f m², Max Plants: %d\n",
			i+1, zone.Id, zone.ZoneName, zone.ZoneCode, zone.AreaM2, zone.MaxPlants)
	}
}

func (c *FarmServiceClient) TestGetZonesByGreenhouse() {
	fmt.Println("\n=== Test Get Zones By Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.GetZonesByGreenhouse(ctx, &proto_growing_zone.GetZonesByGreenhouseRequest{
		GreenhouseId: greenhouseId,
	})
	if err != nil {
		fmt.Printf("Error calling GetZonesByGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Get Zones By Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Growing Zones:\n")
	for i, zone := range resp.GrowingZones {
		fmt.Printf("  [%d] ID: %s, Zone Name: %s, Zone Code: %s, Area: %.2f m², Max Plants: %d\n",
			i+1, zone.Id, zone.ZoneName, zone.ZoneCode, zone.AreaM2, zone.MaxPlants)
	}
}

func (c *FarmServiceClient) TestUpdateGrowingZone() {
	fmt.Println("\n=== Test Update Growing Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter growing zone ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new zone name: ")
	zoneName, _ := reader.ReadString('\n')
	zoneName = cleanInput(zoneName)

	fmt.Print("Enter new zone code: ")
	zoneCode, _ := reader.ReadString('\n')
	zoneCode = cleanInput(zoneCode)

	fmt.Print("Enter new area (m²): ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float32(50.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 32); err == nil {
			area = float32(a)
		}
	}

	fmt.Print("Enter new max plants: ")
	maxPlantsStr, _ := reader.ReadString('\n')
	maxPlantsStr = cleanInput(maxPlantsStr)
	maxPlants := int32(100)
	if maxPlantsStr != "" {
		if mp, err := strconv.Atoi(maxPlantsStr); err == nil {
			maxPlants = int32(mp)
		}
	}

	fmt.Print("Enter new soil type: ")
	soilType, _ := reader.ReadString('\n')
	soilType = cleanInput(soilType)

	fmt.Print("Enter new irrigation system: ")
	irrigationSystem, _ := reader.ReadString('\n')
	irrigationSystem = cleanInput(irrigationSystem)

	fmt.Print("Enter new status: ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.UpdateGrowingZone(ctx, &proto_growing_zone.UpdateGrowingZoneRequest{
		Id:               id,
		ZoneName:         zoneName,
		ZoneCode:         zoneCode,
		AreaM2:           float64(area),
		MaxPlants:        maxPlants,
		SoilType:         soilType,
		IrrigationSystem: irrigationSystem,
		Status:           status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateGrowingZone: %v\n", err)
		return
	}

	fmt.Printf("Update Growing Zone result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
}

func (c *FarmServiceClient) TestDeleteGrowingZone() {
	fmt.Println("\n=== Test Delete Growing Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter growing zone ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneClient.DeleteGrowingZone(ctx, &proto_growing_zone.DeleteGrowingZoneRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteGrowingZone: %v\n", err)
		return
	}

	fmt.Printf("Delete Growing Zone result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
}

// ================== Greenhouse Installation Log Service Tests ==================

func (c *FarmServiceClient) TestCreateInstallationLog() {
	fmt.Println("\n=== Test Create Installation Log ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	fmt.Print("Enter action (INSTALL/UPGRADE/MAINTENANCE/RELOCATE/DISMANTLE): ")
	actionStr, _ := reader.ReadString('\n')
	actionStr = cleanInput(actionStr)

	var action proto_greenhouse_installation_log.InstallationAction
	switch strings.ToUpper(actionStr) {
	case "INSTALL":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_INSTALL
	case "UPGRADE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_UPGRADE
	case "MAINTENANCE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_MAINTENANCE
	case "RELOCATE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_RELOCATE
	case "DISMANTLE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_DISMANTLE
	default:
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_INSTALL
	}

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter performed by: ")
	performedBy, _ := reader.ReadString('\n')
	performedBy = cleanInput(performedBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseInstallationLogClient.CreateLog(ctx, &proto_greenhouse_installation_log.CreateLogRequest{
		GreenhouseId: greenhouseId,
		Action:       action,
		ActionDate:   time.Now().Format("2006-01-02"),
		Description:  description,
		PerformedBy:  performedBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateLog: %v\n", err)
		return
	}

	fmt.Printf("Create Installation Log result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Log ID: %s\n", resp.Log.Id)
	fmt.Printf("Greenhouse ID: %s\n", resp.Log.GreenhouseId)
	fmt.Printf("Action: %s\n", resp.Log.Action.String())
	fmt.Printf("Description: %s\n", resp.Log.Description)
	fmt.Printf("Performed By: %s\n", resp.Log.PerformedBy)
}

func (c *FarmServiceClient) TestGetLogsByGreenhouse() {
	fmt.Println("\n=== Test Get Logs By Greenhouse ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter greenhouse ID: ")
	greenhouseId, _ := reader.ReadString('\n')
	greenhouseId = cleanInput(greenhouseId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseInstallationLogClient.GetLogsByGreenhouse(ctx, &proto_greenhouse_installation_log.GetLogsByGreenhouseRequest{
		GreenhouseId: greenhouseId,
	})
	if err != nil {
		fmt.Printf("Error calling GetLogsByGreenhouse: %v\n", err)
		return
	}

	fmt.Printf("Get Logs By Greenhouse result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Logs:\n")
	for i, log := range resp.Logs {
		fmt.Printf("  [%d] ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, log.Id, log.Action.String(), log.ActionDate.AsTime().Format("2006-01-02 15:04:05"), log.PerformedBy)
		fmt.Printf("      Description: %s\n", log.Description)
	}
}

func (c *FarmServiceClient) TestGetLogsByAction() {
	fmt.Println("\n=== Test Get Logs By Action ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter action (INSTALL/UPGRADE/MAINTENANCE/RELOCATE/DISMANTLE): ")
	actionStr, _ := reader.ReadString('\n')
	actionStr = cleanInput(actionStr)

	var action proto_greenhouse_installation_log.InstallationAction
	switch strings.ToUpper(actionStr) {
	case "INSTALL":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_INSTALL
	case "UPGRADE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_UPGRADE
	case "MAINTENANCE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_MAINTENANCE
	case "RELOCATE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_RELOCATE
	case "DISMANTLE":
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_DISMANTLE
	default:
		action = proto_greenhouse_installation_log.InstallationAction_INSTALLATION_ACTION_INSTALL
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseInstallationLogClient.GetLogsByAction(ctx, &proto_greenhouse_installation_log.GetLogsByActionRequest{
		Action: action,
	})
	if err != nil {
		fmt.Printf("Error calling GetLogsByAction: %v\n", err)
		return
	}

	fmt.Printf("Get Logs By Action result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Logs:\n")
	for i, log := range resp.Logs {
		fmt.Printf("  [%d] ID: %s, Greenhouse ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, log.Id, log.GreenhouseId, log.Action.String(), log.ActionDate.AsTime().Format("2006-01-02 15:04:05"), log.PerformedBy)
		fmt.Printf("      Description: %s\n", log.Description)
	}
}

func (c *FarmServiceClient) TestGetLogsByDateRange() {
	fmt.Println("\n=== Test Get Logs By Date Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start date (YYYY-MM-DD): ")
	startDateStr, _ := reader.ReadString('\n')
	startDateStr = cleanInput(startDateStr)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		startDate = time.Now().AddDate(0, 0, -30) // 30 days ago
	}

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	endDateStr, _ := reader.ReadString('\n')
	endDateStr = cleanInput(endDateStr)
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		endDate = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.greenhouseInstallationLogClient.GetLogsByDateRange(ctx, &proto_greenhouse_installation_log.GetLogsByDateRangeRequest{
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
	})
	if err != nil {
		fmt.Printf("Error calling GetLogsByDateRange: %v\n", err)
		return
	}

	fmt.Printf("Get Logs By Date Range result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Logs:\n")
	for i, log := range resp.Logs {
		fmt.Printf("  [%d] ID: %s, Greenhouse ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, log.Id, log.GreenhouseId, log.Action.String(), log.ActionDate.AsTime().Format("2006-01-02 15:04:05"), log.PerformedBy)
		fmt.Printf("      Description: %s\n", log.Description)
	}
}

// ================== Growing Zone History Service Tests ==================

func (c *FarmServiceClient) TestCreateZoneHistory() {
	fmt.Println("\n=== Test Create Zone History ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter zone ID: ")
	zoneId, _ := reader.ReadString('\n')
	zoneId = cleanInput(zoneId)

	fmt.Print("Enter action (CHANGE_SOIL/CHANGE_IRRIGATION/MAINTENANCE/RESIZE/RENAME): ")
	actionStr, _ := reader.ReadString('\n')
	actionStr = cleanInput(actionStr)

	var action proto_growing_zone_history.HistoryAction
	switch strings.ToUpper(actionStr) {
	case "CHANGE_SOIL":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_CHANGE_SOIL
	case "CHANGE_IRRIGATION":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_CHANGE_IRRIGATION
	case "MAINTENANCE":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_MAINTENANCE
	case "RESIZE":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_RESIZE
	case "RENAME":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_RENAME
	default:
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_MAINTENANCE
	}

	fmt.Print("Enter old value (JSON format, optional): ")
	oldValueStr, _ := reader.ReadString('\n')
	oldValueStr = cleanInput(oldValueStr)
	var oldValue *structpb.Struct
	if oldValueStr != "" {
		oldValue = &structpb.Struct{}
		// Simple JSON parsing - in real implementation, you'd want proper validation
		oldValue.Fields = map[string]*structpb.Value{
			"value": {Kind: &structpb.Value_StringValue{StringValue: oldValueStr}},
		}
	}

	fmt.Print("Enter new value (JSON format, optional): ")
	newValueStr, _ := reader.ReadString('\n')
	newValueStr = cleanInput(newValueStr)
	var newValue *structpb.Struct
	if newValueStr != "" {
		newValue = &structpb.Struct{}
		newValue.Fields = map[string]*structpb.Value{
			"value": {Kind: &structpb.Value_StringValue{StringValue: newValueStr}},
		}
	}

	fmt.Print("Enter performed by: ")
	performedBy, _ := reader.ReadString('\n')
	performedBy = cleanInput(performedBy)

	fmt.Print("Enter notes (optional): ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneHistoryClient.CreateHistory(ctx, &proto_growing_zone_history.CreateHistoryRequest{
		ZoneId:      zoneId,
		Action:      action,
		OldValue:    oldValue,
		NewValue:    newValue,
		PerformedBy: performedBy,
		Notes:       notes,
	})
	if err != nil {
		fmt.Printf("Error calling CreateHistory: %v\n", err)
		return
	}

	fmt.Printf("Create Zone History result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("History ID: %s\n", resp.History.Id)
	fmt.Printf("Zone ID: %s\n", resp.History.ZoneId)
	fmt.Printf("Action: %s\n", resp.History.Action.String())
	fmt.Printf("Performed By: %s\n", resp.History.PerformedBy)
	fmt.Printf("Notes: %s\n", resp.History.Notes)
}

func (c *FarmServiceClient) TestGetHistoryByZone() {
	fmt.Println("\n=== Test Get History By Zone ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter zone ID: ")
	zoneId, _ := reader.ReadString('\n')
	zoneId = cleanInput(zoneId)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneHistoryClient.GetHistoryByZone(ctx, &proto_growing_zone_history.GetHistoryByZoneRequest{
		ZoneId: zoneId,
	})
	if err != nil {
		fmt.Printf("Error calling GetHistoryByZone: %v\n", err)
		return
	}

	fmt.Printf("Get History By Zone result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Histories:\n")
	for i, history := range resp.Histories {
		fmt.Printf("  [%d] ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, history.Id, history.Action.String(), history.ActionDate.AsTime().Format("2006-01-02 15:04:05"), history.PerformedBy)
		fmt.Printf("      Notes: %s\n", history.Notes)
	}
}

func (c *FarmServiceClient) TestGetHistoryByAction() {
	fmt.Println("\n=== Test Get History By Action ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter action (CHANGE_SOIL/CHANGE_IRRIGATION/MAINTENANCE/RESIZE/RENAME): ")
	actionStr, _ := reader.ReadString('\n')
	actionStr = cleanInput(actionStr)

	var action proto_growing_zone_history.HistoryAction
	switch strings.ToUpper(actionStr) {
	case "CHANGE_SOIL":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_CHANGE_SOIL
	case "CHANGE_IRRIGATION":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_CHANGE_IRRIGATION
	case "MAINTENANCE":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_MAINTENANCE
	case "RESIZE":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_RESIZE
	case "RENAME":
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_RENAME
	default:
		action = proto_growing_zone_history.HistoryAction_HISTORY_ACTION_MAINTENANCE
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneHistoryClient.GetHistoryByAction(ctx, &proto_growing_zone_history.GetHistoryByActionRequest{
		Action: action,
	})
	if err != nil {
		fmt.Printf("Error calling GetHistoryByAction: %v\n", err)
		return
	}

	fmt.Printf("Get History By Action result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Histories:\n")
	for i, history := range resp.Histories {
		fmt.Printf("  [%d] ID: %s, Zone ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, history.Id, history.ZoneId, history.Action.String(), history.ActionDate.AsTime().Format("2006-01-02 15:04:05"), history.PerformedBy)
		fmt.Printf("      Notes: %s\n", history.Notes)
	}
}

func (c *FarmServiceClient) TestGetHistoryByDateRange() {
	fmt.Println("\n=== Test Get History By Date Range ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start date (YYYY-MM-DD): ")
	startDateStr, _ := reader.ReadString('\n')
	startDateStr = cleanInput(startDateStr)
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		startDate = time.Now().AddDate(0, 0, -30) // 30 days ago
	}

	fmt.Print("Enter end date (YYYY-MM-DD): ")
	endDateStr, _ := reader.ReadString('\n')
	endDateStr = cleanInput(endDateStr)
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		endDate = time.Now()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneHistoryClient.GetHistoryByDateRange(ctx, &proto_growing_zone_history.GetHistoryByDateRangeRequest{
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
	})
	if err != nil {
		fmt.Printf("Error calling GetHistoryByDateRange: %v\n", err)
		return
	}

	fmt.Printf("Get History By Date Range result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Histories:\n")
	for i, history := range resp.Histories {
		fmt.Printf("  [%d] ID: %s, Zone ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, history.Id, history.ZoneId, history.Action.String(), history.ActionDate.AsTime().Format("2006-01-02 15:04:05"), history.PerformedBy)
		fmt.Printf("      Notes: %s\n", history.Notes)
	}
}

func (c *FarmServiceClient) TestGetHistoryByPerformedBy() {
	fmt.Println("\n=== Test Get History By Performed By ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter performed by: ")
	performedBy, _ := reader.ReadString('\n')
	performedBy = cleanInput(performedBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.growingZoneHistoryClient.GetHistoryByPerformedBy(ctx, &proto_growing_zone_history.GetHistoryByPerformedByRequest{
		PerformedBy: performedBy,
	})
	if err != nil {
		fmt.Printf("Error calling GetHistoryByPerformedBy: %v\n", err)
		return
	}

	fmt.Printf("Get History By Performed By result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Histories:\n")
	for i, history := range resp.Histories {
		fmt.Printf("  [%d] ID: %s, Zone ID: %s, Action: %s, Date: %s, Performed By: %s\n",
			i+1, history.Id, history.ZoneId, history.Action.String(), history.ActionDate.AsTime().Format("2006-01-02 15:04:05"), history.PerformedBy)
		fmt.Printf("      Notes: %s\n", history.Notes)
	}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Farm Service Test Client ===")
	fmt.Println("1. Greenhouse Service")
	fmt.Println("2. Growing Zone Service")
	fmt.Println("3. Greenhouse Installation Log Service")
	fmt.Println("4. Growing Zone History Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printGreenhouseMenu() {
	fmt.Println("\n=== Greenhouse Service ===")
	fmt.Println("1. Create Greenhouse")
	fmt.Println("2. Get Greenhouse")
	fmt.Println("3. List Greenhouses")
	fmt.Println("4. Update Greenhouse")
	fmt.Println("5. Delete Greenhouse")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printGrowingZoneMenu() {
	fmt.Println("\n=== Growing Zone Service ===")
	fmt.Println("1. Create Growing Zone")
	fmt.Println("2. Get Growing Zone")
	fmt.Println("3. List Growing Zones")
	fmt.Println("4. Get Zones By Greenhouse")
	fmt.Println("5. Update Growing Zone")
	fmt.Println("6. Delete Growing Zone")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printGreenhouseInstallationLogMenu() {
	fmt.Println("\n=== Greenhouse Installation Log Service ===")
	fmt.Println("1. Create Installation Log")
	fmt.Println("2. Get Logs By Greenhouse")
	fmt.Println("3. Get Logs By Action")
	fmt.Println("4. Get Logs By Date Range")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printGrowingZoneHistoryMenu() {
	fmt.Println("\n=== Growing Zone History Service ===")
	fmt.Println("1. Create Zone History")
	fmt.Println("2. Get History By Zone")
	fmt.Println("3. Get History By Action")
	fmt.Println("4. Get History By Date Range")
	fmt.Println("5. Get History By Performed By")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewFarmServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Greenhouse Service
			for {
				printGreenhouseMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateGreenhouse()
				case "2":
					client.TestGetGreenhouse()
				case "3":
					client.TestListGreenhouses()
				case "4":
					client.TestUpdateGreenhouse()
				case "5":
					client.TestDeleteGreenhouse()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Growing Zone Service
			for {
				printGrowingZoneMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateGrowingZone()
				case "2":
					client.TestGetGrowingZone()
				case "3":
					client.TestListGrowingZones()
				case "4":
					client.TestGetZonesByGreenhouse()
				case "5":
					client.TestUpdateGrowingZone()
				case "6":
					client.TestDeleteGrowingZone()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "3":
			// Greenhouse Installation Log Service
			for {
				printGreenhouseInstallationLogMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateInstallationLog()
				case "2":
					client.TestGetLogsByGreenhouse()
				case "3":
					client.TestGetLogsByAction()
				case "4":
					client.TestGetLogsByDateRange()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "4":
			// Growing Zone History Service
			for {
				printGrowingZoneHistoryMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateZoneHistory()
				case "2":
					client.TestGetHistoryByZone()
				case "3":
					client.TestGetHistoryByAction()
				case "4":
					client.TestGetHistoryByDateRange()
				case "5":
					client.TestGetHistoryByPerformedBy()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

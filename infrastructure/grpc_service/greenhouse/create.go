package greenhouse_service

import (
	"context"
	"farm-service/domain/entity"
	"time"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GreenhouseService) CreateGreenhouse(ctx context.Context, req *greenhouseP.CreateGreenhouseRequest) (*greenhouseP.CreateGreenhouseResponse, error) {
	greenhouseReq, err := s.createEntityGreenhouseReq(req)
	if err != nil {
		return nil, err
	}
	greenhouse, err := s.greenhouseUsecase.CreateGreenhouse.Execute(ctx, greenhouseReq)
	if err != nil {
		return nil, err
	}
	return &greenhouseP.CreateGreenhouseResponse{
		Success:    true,
		Message:    "Greenhouse created successfully",
		Greenhouse: s.createProtoGreenhouse(greenhouse),
	}, nil
}

func (s *GreenhouseService) createEntityGreenhouseReq(req *greenhouseP.CreateGreenhouseRequest) (*entity.CreateGreenhouseRequest, error) {
	greenhouse := &entity.CreateGreenhouseRequest{
		Name:        req.Name,
		Location:    req.Location,
		AreaM2:      req.AreaM2,
		Type:        req.Type,
		MaxCapacity: req.MaxCapacity,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	}

	if req.InstallationDate != nil {
		installationDate, err := time.Parse(time.RFC3339, req.InstallationDate.String())
		if err != nil {
			return nil, err
		}
		greenhouse.InstallationDate = &installationDate
	}
	return greenhouse, nil
}

func (s *GreenhouseService) createProtoGreenhouse(greenhouse *entity.Greenhouse) *greenhouseP.Greenhouse {
	response := &greenhouseP.Greenhouse{
		Id:          greenhouse.ID,
		Name:        greenhouse.Name,
		Location:    greenhouse.Location,
		AreaM2:      greenhouse.AreaM2,
		Type:        greenhouse.Type,
		MaxCapacity: greenhouse.MaxCapacity,
	}

	if greenhouse.InstallationDate != nil {
		response.InstallationDate = timestamppb.New(*greenhouse.InstallationDate)
	}
	return response
}

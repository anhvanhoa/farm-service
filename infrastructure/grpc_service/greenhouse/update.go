package greenhouse_service

import (
	"context"
	"farm-service/domain/entity"
	"time"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
)

func (s *GreenhouseService) UpdateGreenhouse(ctx context.Context, req *greenhouseP.UpdateGreenhouseRequest) (*greenhouseP.UpdateGreenhouseResponse, error) {
	request, err := s.createEntityUpdateGreenhouse(req)
	if err != nil {
		return nil, err
	}
	err = s.greenhouseUsecase.UpdateGreenhouse.Execute(ctx, req.Id, request)
	if err != nil {
		return nil, err
	}
	return &greenhouseP.UpdateGreenhouseResponse{
		Success: true,
		Message: "Greenhouse updated successfully",
	}, nil
}

func (s *GreenhouseService) createEntityUpdateGreenhouse(req *greenhouseP.UpdateGreenhouseRequest) (*entity.UpdateGreenhouseRequest, error) {
	request := &entity.UpdateGreenhouseRequest{
		Name:        req.Name,
		Location:    req.Location,
		AreaM2:      &req.AreaM2,
		Type:        req.Type,
		MaxCapacity: &req.MaxCapacity,
	}

	if req.InstallationDate != nil {
		installationDate, err := time.Parse(time.RFC3339, req.InstallationDate.String())
		if err != nil {
			return nil, err
		}
		request.InstallationDate = &installationDate
	}
	return request, nil
}

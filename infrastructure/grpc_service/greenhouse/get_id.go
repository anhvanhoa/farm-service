package greenhouse_service

import (
	"context"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
)

func (s *GreenhouseService) GetGreenhouse(ctx context.Context, req *greenhouseP.GetGreenhouseRequest) (*greenhouseP.GetGreenhouseResponse, error) {
	greenhouse, err := s.greenhouseUsecase.GetGreenhouse.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &greenhouseP.GetGreenhouseResponse{
		Success:    true,
		Message:    "Greenhouse retrieved successfully",
		Greenhouse: s.createProtoGreenhouse(greenhouse),
	}, nil
}

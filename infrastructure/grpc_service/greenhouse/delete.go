package greenhouse_service

import (
	"context"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
)

func (s *GreenhouseService) DeleteGreenhouse(ctx context.Context, req *greenhouseP.DeleteGreenhouseRequest) (*greenhouseP.DeleteGreenhouseResponse, error) {
	err := s.greenhouseUsecase.DeleteGreenhouse.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &greenhouseP.DeleteGreenhouseResponse{
		Success: true,
		Message: "Greenhouse deleted successfully",
	}, nil
}

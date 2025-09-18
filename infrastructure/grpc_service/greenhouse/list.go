package greenhouse_service

import (
	"context"
	"farm-service/domain/entity"

	greenhouseP "github.com/anhvanhoa/sf-proto/gen/greenhouse/v1"
)

func (s *GreenhouseService) ListGreenhouses(ctx context.Context, req *greenhouseP.ListGreenhousesRequest) (*greenhouseP.ListGreenhousesResponse, error) {
	greenhouses, totalCount, err := s.greenhouseUsecase.ListGreenhouse.Execute(ctx, s.createEntityGreenhouseFilter(req.Filter), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	return &greenhouseP.ListGreenhousesResponse{
		Success:     true,
		Message:     "Greenhouses retrieved successfully",
		Greenhouses: s.createProtoGreenhouses(greenhouses),
		TotalCount:  totalCount,
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}

func (s *GreenhouseService) createEntityGreenhouseFilter(filter *greenhouseP.GreenhouseFilter) *entity.GreenhouseFilter {
	return &entity.GreenhouseFilter{
		Status:   filter.Status,
		Type:     filter.Type,
		Location: filter.Location,
	}
}

func (s *GreenhouseService) createProtoGreenhouses(greenhouses []*entity.Greenhouse) []*greenhouseP.Greenhouse {
	protoGreenhouses := make([]*greenhouseP.Greenhouse, len(greenhouses))
	for i, greenhouse := range greenhouses {
		protoGreenhouses[i] = s.createProtoGreenhouse(greenhouse)
	}
	return protoGreenhouses
}

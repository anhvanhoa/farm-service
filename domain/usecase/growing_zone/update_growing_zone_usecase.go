package growingzone

import (
	"context"
	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdateGrowingZoneUsecase định nghĩa interface cho use case cập nhật khu vực trồng
type UpdateGrowingZoneUsecase interface {
	Execute(ctx context.Context, id string, req *entity.UpdateGrowingZoneRequest) error
}

type updateGrowingZoneUsecase struct {
	growingZoneRepository repository.GrowingZoneRepository
}

// NewUpdateGrowingZoneUsecase tạo instance mới của UpdateGrowingZoneUsecase
func NewUpdateGrowingZoneUsecase(growingZoneRepository repository.GrowingZoneRepository) UpdateGrowingZoneUsecase {
	return &updateGrowingZoneUsecase{
		growingZoneRepository: growingZoneRepository,
	}
}

func (u *updateGrowingZoneUsecase) Execute(ctx context.Context, id string, req *entity.UpdateGrowingZoneRequest) error {
	// Kiểm tra khu vực trồng có tồn tại không
	_, err := u.growingZoneRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Nếu có thay đổi zone code, kiểm tra xem code mới có bị trùng không
	if req.ZoneCode != "" {
		exists, err := u.growingZoneRepository.CheckZoneCodeExists(ctx, req.ZoneCode)
		if err != nil {
			return err
		}
		if exists {
			return &entity.Error{Code: "ZONE_CODE_EXISTS", Message: "Zone code already exists"}
		}
	}

	// Cập nhật thông tin
	return u.growingZoneRepository.Update(ctx, id, req)
}

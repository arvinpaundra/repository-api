package request_access

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	requestAccess "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/requestAccess/request"
	"github.com/arvinpaundra/repository-api/models/web/requestAccess/response"
	"gorm.io/gorm"
)

type RequestAccessServiceImpl struct {
	requestAccessRepository requestAccess.RequestAccessRepository
	pemustakaRepository     pemustaka.PemustakaRepository

	tx *gorm.DB
}

func NewRequestAccessService(
	requestAccessRepository requestAccess.RequestAccessRepository,
	pemustakaRepository pemustaka.PemustakaRepository,
	tx *gorm.DB,
) RequestAccessService {
	return RequestAccessServiceImpl{
		requestAccessRepository: requestAccessRepository,
		pemustakaRepository:     pemustakaRepository,
		tx:                      tx,
	}
}

func (service RequestAccessServiceImpl) Update(ctx context.Context, requestAccessDTO request.UpdateRequestAccessRequest, requestAccessId string) error {
	tx := service.tx.Begin()

	requestAccess, err := service.requestAccessRepository.FindById(ctx, requestAccessId)

	if err != nil {
		return err
	}

	if _, err := service.pemustakaRepository.FindById(ctx, requestAccess.PemustakaId); err != nil {
		return err
	}

	requestAccessDomain := requestAccessDTO.ToDomainRequestAccess()

	if err := service.requestAccessRepository.Update(ctx, tx, requestAccessDomain, requestAccessId); err != nil {
		if errorRollback := tx.Rollback().Error; errorRollback != nil {
			return errorRollback
		}

		return err
	}

	if requestAccessDTO.Status == "accepted" {
		pemustakaDomain := domain.Pemustaka{
			IsActive: "1",
		}

		if err := service.pemustakaRepository.Update(ctx, tx, pemustakaDomain, requestAccess.PemustakaId); err != nil {
			if errorRollback := tx.Rollback().Error; errorRollback != nil {
				return errorRollback
			}

			return err
		}
	}

	if errorCommit := tx.Commit().Error; errorCommit != nil {
		return errorCommit
	}

	return nil
}

func (service RequestAccessServiceImpl) FindAll(ctx context.Context, keyword string, status string, limit int, offset int) ([]response.RequestAccessResponse, int, int, error) {
	requestAccesses, totalRows, err := service.requestAccessRepository.FindAll(ctx, keyword, status, limit, offset)

	if err != nil {
		return []response.RequestAccessResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToRequestAccessesResponse(requestAccesses), totalRows, int(totalPages), nil
}

func (service RequestAccessServiceImpl) FindById(ctx context.Context, requestAccessId string) (response.RequestAccessResponse, error) {
	requestAccess, err := service.requestAccessRepository.FindById(ctx, requestAccessId)

	if err != nil {
		return response.RequestAccessResponse{}, err
	}

	return response.ToRequestAccessResponse(requestAccess), nil
}

func (service RequestAccessServiceImpl) GetTotal(ctx context.Context, status string) (int, error) {
	total, err := service.requestAccessRepository.Total(ctx, status)

	if err != nil {
		return 0, err
	}

	return total, nil
}

package departement

import (
	"context"
	"math"

	"github.com/arvinpaundra/repository-api/drivers/mysql/departement"
	"github.com/arvinpaundra/repository-api/models/web/departement/request"
	"github.com/arvinpaundra/repository-api/models/web/departement/response"
	"github.com/google/uuid"
)

type DepartementServiceImpl struct {
	departementRepository departement.DepartementRepository
}

func NewDepartementService(departementRepository departement.DepartementRepository) DepartementService {
	return DepartementServiceImpl{
		departementRepository: departementRepository,
	}
}

func (service DepartementServiceImpl) Create(ctx context.Context, departement request.CreateDepartementRequest) error {
	departementDomain := departement.ToDomainDepartement()

	departementDomain.ID = uuid.NewString()

	err := service.departementRepository.Save(ctx, departementDomain)

	if err != nil {
		return err
	}

	return nil
}

func (service DepartementServiceImpl) Update(ctx context.Context, departement request.UpdateDepartementRequest, departementId string) error {
	if _, err := service.departementRepository.FindById(ctx, departementId); err != nil {
		return err
	}

	err := service.departementRepository.Update(ctx, departement.ToDomainDepartement(), departementId)

	if err != nil {
		return err
	}

	return nil
}

func (service DepartementServiceImpl) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.DepartementResponse, int, int, error) {
	departements, totalRows, err := service.departementRepository.FindAll(ctx, keyword, limit, offset)

	if err != nil {
		return []response.DepartementResponse{}, 0, 0, err
	}

	totalPages := math.Ceil(float64(totalRows) / float64(limit))

	return response.ToDepartementsResponse(departements), totalRows, int(totalPages), nil
}

func (service DepartementServiceImpl) FindById(ctx context.Context, departementId string) (response.DepartementResponse, error) {
	departement, err := service.departementRepository.FindById(ctx, departementId)

	if err != nil {
		return response.DepartementResponse{}, err
	}

	return response.ToDepartementResponse(departement), nil
}

func (service DepartementServiceImpl) Delete(ctx context.Context, departementId string) error {
	panic("not implemented")
}

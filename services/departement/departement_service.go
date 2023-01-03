package departement

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/departement/request"
	"github.com/arvinpaundra/repository-api/models/web/departement/response"
)

type DepartementService interface {
	Create(ctx context.Context, departement request.CreateDepartementRequest) error
	Update(ctx context.Context, departement request.UpdateDepartementRequest, departementId string) error
	Delete(ctx context.Context, departementId string) error
	FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.DepartementResponse, int, int, error)
	FindById(ctx context.Context, departementId string) (response.DepartementResponse, error)
	FindByProgramStudyId(ctx context.Context, studyProgramId string) ([]response.DepartementResponse, error)
}

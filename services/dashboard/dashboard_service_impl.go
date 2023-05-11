package dashboard

import (
	"context"

	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/repository"
	requestAccess "github.com/arvinpaundra/repository-api/drivers/mysql/requestAccess"
)

type DashboardServiceImpl struct {
	pemustakaRepository     pemustaka.PemustakaRepository
	repository              repository.Repository
	requestAccessRepository requestAccess.RequestAccessRepository
}

func NewDashboardService(
	pemustakaRepository pemustaka.PemustakaRepository,
	repository repository.Repository,
	requestAccessRepository requestAccess.RequestAccessRepository,
) DashboardService {
	return DashboardServiceImpl{
		pemustakaRepository:     pemustakaRepository,
		repository:              repository,
		requestAccessRepository: requestAccessRepository,
	}
}

func (service DashboardServiceImpl) Overview(ctx context.Context) (int, int, int, error) {
	totalPemustaka, err := service.pemustakaRepository.GetTotal(ctx)

	if err != nil {
		return 0, 0, 0, err
	}

	totalRepository, err := service.repository.GetTotal(ctx, "")

	if err != nil {
		return 0, 0, 0, err
	}

	totalRequestAccessPending, err := service.requestAccessRepository.Total(ctx, "pending")

	if err != nil {
		return 0, 0, 0, err
	}

	return totalPemustaka, totalRepository, totalRequestAccessPending, nil
}

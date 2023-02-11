package pemustaka

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/pemustaka/request"
	"github.com/arvinpaundra/repository-api/models/web/pemustaka/response"
)

type PemustakaService interface {
	Register(ctx context.Context, req request.RegisterPemustakaRequest) error
	Login(ctx context.Context, req request.LoginPemustakaRequest) (string, error)
	Update(ctx context.Context, req request.UpdatePemustakaRequest, pemustakaId string) error
	FindAll(ctx context.Context, query request.PemustakaRequestQuery, limit int, offset int) ([]response.PemustakaResponse, int, int, error)
	FindById(ctx context.Context, pemustakaId string) (response.PemustakaResponse, error)
}

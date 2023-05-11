package request_access

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/requestAccess/request"
	"github.com/arvinpaundra/repository-api/models/web/requestAccess/response"
)

type RequestAccessService interface {
	Update(ctx context.Context, requestAccessDTO request.UpdateRequestAccessRequest, requestAccessId string) error
	FindAll(ctx context.Context, keyword string, status string, limit int, offset int) ([]response.RequestAccessResponse, int, int, error)
	FindById(ctx context.Context, requestAccessId string) (response.RequestAccessResponse, error)
	GetTotal(ctx context.Context, status string) (int, error)
}

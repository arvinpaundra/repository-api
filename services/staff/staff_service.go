package staff

import (
	"context"
	"mime/multipart"

	"github.com/arvinpaundra/repository-api/models/web/staff/request"
	"github.com/arvinpaundra/repository-api/models/web/staff/response"
)

type StaffService interface {
	Register(ctx context.Context, req request.RegisterStaffRequest) error
	Login(ctx context.Context, req request.LoginStaffRequest) (string, error)
	Update(ctx context.Context, req request.UpdateStaffRequest, avatar *multipart.FileHeader, staffId string) error
	FindAll(ctx context.Context, query request.StaffRequestQuery, limit int, offset int) ([]response.StaffResponse, int, int, error)
	FindById(ctx context.Context, staffId string) (response.StaffResponse, error)
	UploadSignature(ctx context.Context, signature *multipart.FileHeader, staffId string) error
}

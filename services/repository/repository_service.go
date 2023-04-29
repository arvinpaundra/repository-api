package repository

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"github.com/arvinpaundra/repository-api/models/web/repository/response"
)

type RepositoryService interface {
	CreateFinalProjectReport(ctx context.Context, repositoryDTO request.CreateFinalProjectReportRequest, files request.RepositoryInputFiles) error
	CreateInternshipReport(ctx context.Context, repositoryDTO request.CreateInternshipReportRequest, files request.RepositoryInputFiles) error
	CreateResearchReport(ctx context.Context, repositoryDTO request.CreateResearchReportRequest, files request.RepositoryInputFiles) error
	Delete(ctx context.Context, repositoryId string) error
	FindAll(ctx context.Context, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
	FindById(ctx context.Context, repositoryId string) (response.DetailRepositoryResponse, error)
	FindByAuthorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
	FindByMentorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
	FindByExaminerId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
	FindByCollectionId(ctx context.Context, collectionId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
	FindByDepartementId(ctx context.Context, departementId string, query request.RepositoryRequestQuery, limit int, offset int) ([]response.RepositoryResponse, int, int, error)
}

package repository

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/repository/request"
	"gorm.io/gorm"
)

type Repository interface {
	Save(ctx context.Context, tx *gorm.DB, repository domain.Repository) error
	Update(ctx context.Context, tx *gorm.DB, repository domain.Repository, repositoryId string) error
	Delete(ctx context.Context, repositoryId string) error
	FindAll(ctx context.Context, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	FindById(ctx context.Context, repositoryId string) (domain.Repository, error)
	FindByAuthorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	FindByMentorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	FindByExaminerId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	FindByCollectionId(ctx context.Context, collectionId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	FindByDepartementId(ctx context.Context, departementId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error)
	GetTotal(ctx context.Context, status string) (int, error)
}

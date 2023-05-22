package contributor

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type ContributorRepository interface {
	Save(ctx context.Context, tx *gorm.DB, contributor domain.Contributor) error
	Update(ctx context.Context, tx *gorm.DB, contributorId string, contributor domain.Contributor) error
	Delete(ctx context.Context, repositoryId string, pemustakaId string) error
	FindByRepositoryId(ctx context.Context, repositoryId string) ([]domain.Contributor, error)
}

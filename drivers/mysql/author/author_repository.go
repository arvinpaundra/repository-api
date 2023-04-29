package author

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	Save(ctx context.Context, tx *gorm.DB, author []domain.Author) error
	Delete(ctx context.Context, repositoryId string, pemustakaId string) error
	FindByRepositoryId(ctx context.Context, repositoryId string) ([]domain.Author, error)
}

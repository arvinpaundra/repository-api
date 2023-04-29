package document

import (
	"context"

	"github.com/arvinpaundra/repository-api/models/domain"
	"gorm.io/gorm"
)

type DocumentRepository interface {
	Save(ctx context.Context, tx *gorm.DB, document domain.Document) error
	Update(ctx context.Context, tx *gorm.DB, document domain.Document, repositoryId string) error
	FindByRepositoryId(ctx context.Context, repositoryId string) (domain.Document, error)
}

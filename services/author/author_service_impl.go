package author

import (
	"context"

	"github.com/arvinpaundra/repository-api/drivers/mysql/author"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/drivers/mysql/repository"
)

type AuthorServiceImpl struct {
	authorRepository    author.AuthorRepository
	repository          repository.Repository
	pemustakaRepository pemustaka.PemustakaRepository
}

func NewAuthorService(
	authorRepository author.AuthorRepository,
	repository repository.Repository,
	pemustakaRepository pemustaka.PemustakaRepository,
) AuthorService {
	return AuthorServiceImpl{
		authorRepository:    authorRepository,
		repository:          repository,
		pemustakaRepository: pemustakaRepository,
	}
}

func (service AuthorServiceImpl) Delete(ctx context.Context, repositoryId string, pemustakaId string) error {
	if _, err := service.repository.FindById(ctx, repositoryId); err != nil {
		return err
	}

	if _, err := service.pemustakaRepository.FindById(ctx, pemustakaId); err != nil {
		return err
	}

	err := service.authorRepository.Delete(ctx, repositoryId, pemustakaId)

	if err != nil {
		return err
	}

	return nil
}

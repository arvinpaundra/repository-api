package author

import "context"

type AuthorService interface {
	Delete(ctx context.Context, repositoryId string, pemustakaId string) error
}

package identity_card

import "context"

type IdentityCardService interface {
	Generate(ctx context.Context, pemustakaId string) ([]byte, error)
}

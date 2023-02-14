package cloudinary

import (
	"context"
	"mime/multipart"
)

type Cloudinary interface {
	Upload(ctx context.Context, folder string, filename string, file *multipart.FileHeader) (string, error)
	Delete(ctx context.Context, filename string) error
}

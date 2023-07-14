package cloudinary

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryImpl struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinary() Cloudinary {
	cld, err := cloudinary.NewFromParams(configs.GetConfig("CLD_CLOUD"), configs.GetConfig("CLD_KEY"), configs.GetConfig("CLD_SECRET"))

	if err != nil {
		panic(err)
	}

	return CloudinaryImpl{
		cloudinary: cld,
	}
}

func (c CloudinaryImpl) Upload(ctx context.Context, folder string, filename string, file *multipart.FileHeader) (string, error) {
	timeout, cancel := context.WithTimeout(ctx, 30*time.Second)

	defer cancel()

	fileBuffer, err := file.Open()

	if err != nil {
		return "", err
	}

	defer fileBuffer.Close()

	result, err := c.cloudinary.Upload.Upload(timeout, fileBuffer, uploader.UploadParams{
		PublicID: filename,
		Folder:   folder,
	})

	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}

func (c CloudinaryImpl) Delete(ctx context.Context, filename string) error {
	timeout, cancel := context.WithTimeout(ctx, 30*time.Second)

	defer cancel()

	_, err := c.cloudinary.Upload.Destroy(timeout, uploader.DestroyParams{
		PublicID: filename,
	})

	if err != nil {
		return err
	}

	return nil
}

package identity_card

import (
	"bytes"
	"context"
	"html/template"

	"github.com/arvinpaundra/repository-api/configs"
	"github.com/arvinpaundra/repository-api/drivers/mysql/pemustaka"
	"github.com/arvinpaundra/repository-api/helper"
	"github.com/arvinpaundra/repository-api/templates"
	"github.com/ninetwentyfour/go-wkhtmltoimage"
)

type IdentityCardServiceImpl struct {
	pemustakaRepository pemustaka.PemustakaRepository
}

func NewIdentityCardService(pemustakaRepository pemustaka.PemustakaRepository) IdentityCardService {
	return IdentityCardServiceImpl{
		pemustakaRepository: pemustakaRepository,
	}
}

func (service IdentityCardServiceImpl) Generate(ctx context.Context, pemustakaId string) ([]byte, error) {
	pemustaka, err := service.pemustakaRepository.FindById(ctx, pemustakaId)

	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("").Parse(templates.IDCard)

	if err != nil {
		return nil, err
	}

	idCardData := map[string]string{
		"avatar":      pemustaka.Avatar,
		"fullname":    pemustaka.Fullname,
		"memberCode":  pemustaka.MemberCode,
		"createdAt":   helper.FormatDate(pemustaka.CreatedAt),
		"departement": pemustaka.Departement.Name,
		"address":     pemustaka.Address,
	}

	// parse and keep html string and the data into buff
	var buff bytes.Buffer

	if err := tmpl.Execute(&buff, idCardData); err != nil {
		return nil, err
	}

	imageOptions := wkhtmltoimage.ImageOptions{BinaryPath: configs.GetConfig("WKHTMLTOIMAGE_PATH"), Input: "-", Format: "png", Html: buff.String(), Width: 450}

	res, err := wkhtmltoimage.GenerateImage(&imageOptions)

	if err != nil {
		return nil, err
	}

	return res, nil
}

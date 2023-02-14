package request

import (
	"github.com/arvinpaundra/repository-api/models/domain"
)

type UpdatePemustakaRequest struct {
	StudyProgramId string `json:"study_program_id" form:"study_program_id" validate:"required"`
	DepartementId  string `json:"departement_id" form:"departement_id" validate:"required"`
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	YearGen        string `json:"year_gen" form:"year_gen"`
	Gender         string `json:"gender" form:"gender"`
	Telp           string `json:"telp" form:"telp"`
	BirthDate      string `json:"birth_date" form:"birth_date"`
	Address        string `json:"address" form:"address"`
}

func (req *UpdatePemustakaRequest) ToDomainPemustaka() domain.Pemustaka {
	return domain.Pemustaka{
		StudyProgramId: req.StudyProgramId,
		DepartementId:  req.DepartementId,
		Fullname:       req.Fullname,
		YearGen:        req.YearGen,
		Gender:         req.Gender,
		Telp:           req.Telp,
		BirthDate:      req.BirthDate,
		Address:        req.Address,
	}
}

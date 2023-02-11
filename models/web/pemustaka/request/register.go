package request

import (
	"github.com/arvinpaundra/repository-api/models/domain"
)

type RegisterPemustakaRequest struct {
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required"`
	StudyProgramId string `json:"study_program_id" form:"study_program_id" validate:"required"`
	DepartementId  string `json:"departement_id" form:"departement_id" validate:"required"`
	RoleId         string `json:"role_id" form:"role_id" validate:"required"`
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	IdentityNumber string `json:"identity_number" form:"identity_number" validate:"required"`
	YearGen        string `json:"year_gen" form:"year_gen"`
	// SuppFile       *multipart.FileHeader `json:"supp_file" form:"supp_file" validate:"required"`
}

func (req *RegisterPemustakaRequest) ToDomainPemustaka() domain.Pemustaka {
	return domain.Pemustaka{
		StudyProgramId: req.StudyProgramId,
		DepartementId:  req.DepartementId,
		RoleId:         req.RoleId,
		Fullname:       req.Fullname,
		IdentityNumber: req.IdentityNumber,
		YearGen:        req.YearGen,
	}
}

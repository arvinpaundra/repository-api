package request

import (
	"github.com/arvinpaundra/repository-api/models/domain"
)

type UpdatePemustakaRequest struct {
	StudyProgramId              string `json:"study_program_id" form:"study_program_id" validate:"required"`
	DepartementId               string `json:"departement_id" form:"departement_id" validate:"required"`
	RoleId                      string `json:"role_id" form:"role_id" validate:"required"`
	Fullname                    string `json:"fullname" form:"fullname" validate:"required"`
	Email                       string `json:"email" form:"email" validate:"required,email"`
	IdentityNumber              string `json:"identity_number" form:"identity_number" validate:"required"`
	YearGen                     string `json:"year_gen" form:"year_gen"`
	Gender                      string `json:"gender" form:"gender"`
	Telp                        string `json:"telp" form:"telp"`
	BirthDate                   string `json:"birth_date" form:"birth_date"`
	Address                     string `json:"address" form:"address"`
	IsActive                    string `json:"is_active" form:"is_active" validate:"required"`
	IsCollectedFinalProject     string `json:"is_collected_final_project" form:"is_collected_final_project" validate:"required"`
	IsCollectedInternshipReport string `json:"is_collected_internship_report" form:"is_collected_internship_report" validate:"required"`
}

func (req *UpdatePemustakaRequest) ToDomainPemustaka() domain.Pemustaka {
	return domain.Pemustaka{
		DepartementId:               req.DepartementId,
		StudyProgramId:              req.StudyProgramId,
		RoleId:                      req.RoleId,
		Fullname:                    req.Fullname,
		IdentityNumber:              req.IdentityNumber,
		YearGen:                     req.YearGen,
		Gender:                      req.Gender,
		Telp:                        req.Telp,
		BirthDate:                   req.BirthDate,
		Address:                     req.Address,
		IsActive:                    req.IsActive,
		IsCollectedFinalProject:     req.IsCollectedFinalProject,
		IsCollectedInternshipReport: req.IsCollectedInternshipReport,
	}
}

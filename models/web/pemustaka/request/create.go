package request

import "github.com/arvinpaundra/repository-api/models/domain"

type CreatePemustakaRequest struct {
	RoleId                      string `json:"role_id" form:"role_id" validate:"required"`
	DepartementId               string `json:"departement_id" form:"departement_id" validate:"required"`
	StudyProgramId              string `json:"study_program_id" form:"study_program_id" validate:"required"`
	Fullname                    string `json:"fullname" form:"fullname" validate:"required"`
	Email                       string `json:"email" form:"email" validate:"required,email"`
	Telp                        string `json:"telp" form:"telp" validate:"max=13"`
	BirthDate                   string `json:"birth_date" form:"birth_date"`
	Gender                      string `json:"gender" form:"gender"`
	Address                     string `json:"address" form:"address"`
	IdentityNumber              string `json:"identity_number" form:"identity_number" validate:"required,numeric,min=9"`
	YearGen                     string `json:"year_gen" form:"year_gen" validate:"max=4"`
	IsActive                    string `json:"is_active" form:"is_active" validate:"required"`
	IsCollectedFinalProject     string `json:"is_collected_final_project" form:"is_collected_final_project" validate:"required"`
	IsCollectedInternshipReport string `json:"is_collected_internship_report" form:"is_collected_internship_report" validate:"required"`
}

func (req *CreatePemustakaRequest) ToPemustakaDomain() domain.Pemustaka {
	return domain.Pemustaka{
		RoleId:                      req.RoleId,
		DepartementId:               req.DepartementId,
		StudyProgramId:              req.StudyProgramId,
		Fullname:                    req.Fullname,
		Telp:                        req.Telp,
		BirthDate:                   req.BirthDate,
		Gender:                      req.Gender,
		Address:                     req.Address,
		IdentityNumber:              req.IdentityNumber,
		YearGen:                     req.YearGen,
		IsActive:                    req.IsActive,
		IsCollectedFinalProject:     req.IsCollectedFinalProject,
		IsCollectedInternshipReport: req.IsCollectedInternshipReport,
	}
}

func (req *CreatePemustakaRequest) ToUserDomain() domain.User {
	return domain.User{
		Email: req.Email,
	}
}

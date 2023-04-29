package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type PemustakaResponse struct {
	ID                          string    `json:"id"`
	UserId                      string    `json:"user_id"`
	StudyProgramId              string    `json:"study_program_id"`
	DepartementId               string    `json:"departement_id"`
	StudyProgram                string    `json:"study_program"`
	Departement                 string    `json:"departement"`
	Role                        string    `json:"role"`
	MemberCode                  string    `json:"member_code"`
	Email                       string    `json:"email"`
	Fullname                    string    `json:"fullname"`
	IdentityNumber              string    `json:"identity_number"`
	YearGen                     string    `json:"year_gen"`
	Gender                      string    `json:"gender"`
	Telp                        string    `json:"telp"`
	BirthDate                   string    `json:"birth_date"`
	Address                     string    `json:"address"`
	IsCollectedFinalProject     string    `json:"is_collected_final_project"`
	IsCollectedInternshipReport string    `json:"is_collected_internship_report"`
	IsActive                    string    `json:"is_active"`
	Avatar                      string    `json:"avatar"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
}

func ToPemustakaResponse(pemustakaDomain domain.Pemustaka) PemustakaResponse {
	return PemustakaResponse{
		ID:                          pemustakaDomain.ID,
		UserId:                      pemustakaDomain.UserId,
		StudyProgramId:              pemustakaDomain.StudyProgramId,
		DepartementId:               pemustakaDomain.DepartementId,
		StudyProgram:                pemustakaDomain.StudyProgram.Name,
		Departement:                 pemustakaDomain.Departement.Name,
		Role:                        pemustakaDomain.Role.Role,
		MemberCode:                  pemustakaDomain.MemberCode,
		Email:                       pemustakaDomain.User.Email,
		Fullname:                    pemustakaDomain.Fullname,
		IdentityNumber:              pemustakaDomain.IdentityNumber,
		YearGen:                     pemustakaDomain.YearGen,
		Gender:                      pemustakaDomain.Gender,
		Telp:                        pemustakaDomain.Telp,
		BirthDate:                   pemustakaDomain.BirthDate,
		Address:                     pemustakaDomain.Address,
		IsCollectedFinalProject:     pemustakaDomain.IsCollectedFinalProject,
		IsCollectedInternshipReport: pemustakaDomain.IsCollectedInternshipReport,
		IsActive:                    pemustakaDomain.IsActive,
		Avatar:                      pemustakaDomain.Avatar,
		CreatedAt:                   pemustakaDomain.CreatedAt,
		UpdatedAt:                   pemustakaDomain.UpdatedAt,
	}
}

func ToPemustakaArrayResponse(pemustakaDomain []domain.Pemustaka) []PemustakaResponse {
	var pemustakas []PemustakaResponse

	for _, pemustaka := range pemustakaDomain {
		pemustakas = append(pemustakas, ToPemustakaResponse(pemustaka))
	}

	return pemustakas
}

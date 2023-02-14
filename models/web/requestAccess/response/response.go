package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type RequestAccessResponse struct {
	ID              string    `json:"id"`
	PemustakaId     string    `json:"pemustaka_id"`
	Pemustaka       string    `json:"pemustaka"`
	Role            string    `json:"role"`
	StudyProgram    string    `json:"study_program"`
	Departement     string    `json:"departement"`
	IdentityNumber  string    `json:"identity_number"`
	SupportEvidence string    `json:"support_evidence"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func ToRequestAccessResponse(requestAccessDomain domain.RequestAccess) RequestAccessResponse {
	return RequestAccessResponse{
		ID:              requestAccessDomain.ID,
		PemustakaId:     requestAccessDomain.PemustakaId,
		Pemustaka:       requestAccessDomain.Pemustaka.Fullname,
		Role:            requestAccessDomain.Pemustaka.Role.Role,
		StudyProgram:    requestAccessDomain.Pemustaka.StudyProgram.Name,
		Departement:     requestAccessDomain.Pemustaka.Departement.Name,
		IdentityNumber:  requestAccessDomain.Pemustaka.IdentityNumber,
		SupportEvidence: requestAccessDomain.SupportEvidence,
		Status:          requestAccessDomain.Status,
		CreatedAt:       requestAccessDomain.CreatedAt,
		UpdatedAt:       requestAccessDomain.UpdatedAt,
	}
}

func ToRequestAccessesResponse(requestAccessDomain []domain.RequestAccess) []RequestAccessResponse {
	var requestAccesses []RequestAccessResponse

	for _, requestAccess := range requestAccessDomain {
		requestAccesses = append(requestAccesses, ToRequestAccessResponse(requestAccess))
	}

	return requestAccesses
}

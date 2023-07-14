package response

import (
	"time"

	"github.com/arvinpaundra/repository-api/models/domain"
)

type StaffResponse struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	RoleId    string    `json:"role_id"`
	Fullname  string    `json:"fullname"`
	Nip       string    `json:"nip"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Telp      string    `json:"telp"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	Avatar    string    `json:"avatar"`
	BirthDate string    `json:"birth_date"`
	IsActive  string    `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToStaffResponse(staffDomain domain.Staff) StaffResponse {
	return StaffResponse{
		ID:        staffDomain.ID,
		UserId:    staffDomain.UserId,
		RoleId:    staffDomain.Role.ID,
		Fullname:  staffDomain.Fullname,
		Nip:       staffDomain.Nip,
		Email:     staffDomain.User.Email,
		Role:      staffDomain.Role.Role,
		Telp:      staffDomain.Telp,
		Address:   staffDomain.Address,
		Gender:    staffDomain.Gender,
		Avatar:    staffDomain.Avatar,
		BirthDate: staffDomain.BirthDate,
		IsActive:  staffDomain.IsActive,
		CreatedAt: staffDomain.CreatedAt,
		UpdatedAt: staffDomain.UpdatedAt,
	}
}

func ToStaffArrayResponse(staffDomain []domain.Staff) []StaffResponse {
	var staffs []StaffResponse

	for _, staff := range staffDomain {
		staffs = append(staffs, ToStaffResponse(staff))
	}

	return staffs
}

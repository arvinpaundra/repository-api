package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateStaffRequest struct {
	RoleId    string `json:"role_id" form:"role_id" validate:"required"`
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Nip       string `json:"nip" form:"nip"`
	Telp      string `json:"telp" form:"telp"`
	Address   string `json:"address" form:"address"`
	Gender    string `json:"gender" form:"gender"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	IsActive  string `json:"is_active" form:"is_active" validate:"required"`
}

func (req *UpdateStaffRequest) ToDomainStaff() domain.Staff {
	return domain.Staff{
		RoleId:    req.RoleId,
		Fullname:  req.Fullname,
		Nip:       req.Nip,
		Telp:      req.Telp,
		Address:   req.Address,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
		IsActive:  req.IsActive,
	}
}

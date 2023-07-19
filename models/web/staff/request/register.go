package request

import "github.com/arvinpaundra/repository-api/models/domain"

type RegisterStaffRequest struct {
	RoleId    string `json:"role_id" form:"role_id" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
	Fullname  string `json:"fullname" form:"fullname" validate:"required"`
	Nip       string `json:"nip" form:"nip" validate:"max=18"`
	BirthDate string `json:"birth_date" form:"birth_date"`
	Gender    string `json:"gender" form:"gender"`
	Telp      string `json:"telp" form:"telp" validate:"max=13"`
	Address   string `json:"address" form:"address"`
	IsActive  string `json:"is_active" form:"is_active"`
}

func (req *RegisterStaffRequest) ToDomainStaff() domain.Staff {
	return domain.Staff{
		Fullname:  req.Fullname,
		RoleId:    req.RoleId,
		Nip:       req.Nip,
		Telp:      req.Telp,
		Address:   req.Address,
		Gender:    req.Gender,
		BirthDate: req.BirthDate,
		IsActive:  req.IsActive,
	}
}

package request

import "github.com/arvinpaundra/repository-api/models/domain"

type RegisterStaffRequest struct {
	RoleId   string `json:"role_id" form:"role_id" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Nip      string `json:"nip" form:"nip"`
}

func (req *RegisterStaffRequest) ToDomainStaff() domain.Staff {
	return domain.Staff{
		Fullname: req.Fullname,
		RoleId:   req.RoleId,
		Nip:      req.Nip,
	}
}

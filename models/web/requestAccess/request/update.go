package request

import "github.com/arvinpaundra/repository-api/models/domain"

type UpdateRequestAccessRequest struct {
	Status string `json:"status" form:"status" validate:"required"`
}

func (req *UpdateRequestAccessRequest) ToDomainRequestAccess() domain.RequestAccess {
	return domain.RequestAccess{
		Status: req.Status,
	}
}

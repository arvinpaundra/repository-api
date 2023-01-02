package service

import (
	"context"

	"github.com/arvinpaundra/repository-api/drivers/mysql/role"
	"github.com/arvinpaundra/repository-api/models/web/role/request"
	"github.com/arvinpaundra/repository-api/models/web/role/response"
	"github.com/google/uuid"
)

type RoleServiceImpl struct {
	roleRepository role.RoleRepository
}

func NewRoleService(roleRepository role.RoleRepository) RoleService {
	return RoleServiceImpl{
		roleRepository: roleRepository,
	}
}

func (service RoleServiceImpl) Create(ctx context.Context, role request.CreateRoleRequest) error {
	roleDomain := role.ToDomainRole()

	roleDomain.ID = uuid.NewString()

	err := service.roleRepository.Save(ctx, roleDomain)

	if err != nil {
		return err
	}

	return nil
}

func (service RoleServiceImpl) Update(ctx context.Context, role request.UpdateRoleRequest, roleId string) error {
	if _, err := service.roleRepository.FindById(ctx, roleId); err != nil {
		return err
	}

	err := service.roleRepository.Update(ctx, role.ToDomainRole(), roleId)

	if err != nil {
		return err
	}

	return nil
}

func (service RoleServiceImpl) FindAll(ctx context.Context, visibility string) ([]response.RoleResponse, error) {
	roles, err := service.roleRepository.FindAll(ctx, visibility)

	if err != nil {
		return []response.RoleResponse{}, err
	}

	return response.ToRolesResponse(roles), nil
}

func (service RoleServiceImpl) FindById(ctx context.Context, roleId string) (response.RoleResponse, error) {
	role, err := service.roleRepository.FindById(ctx, roleId)

	if err != nil {
		return response.RoleResponse{}, err
	}

	return response.ToRoleResponse(role), nil
}

func (service RoleServiceImpl) Delete(ctx context.Context, roleId string) error {
	panic("not implemented")
}

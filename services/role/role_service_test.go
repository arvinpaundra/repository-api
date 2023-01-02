package role_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arvinpaundra/repository-api/drivers/mysql/role/mocks"
	"github.com/arvinpaundra/repository-api/models/domain"
	"github.com/arvinpaundra/repository-api/models/web/role/request"
	"github.com/arvinpaundra/repository-api/services/role"
	"github.com/arvinpaundra/repository-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	roleRepository mocks.RoleRepository
	roleService    role.RoleService

	roleDomain    domain.Role
	createRoleDTO request.CreateRoleRequest
	updateRoleDTO request.UpdateRoleRequest

	ctx context.Context
)

func TestMain(m *testing.M) {
	roleService = role.NewRoleService(&roleRepository)

	roleDomain = domain.Role{
		ID:         uuid.NewString(),
		Role:       "test",
		Visibility: "test",
	}

	createRoleDTO = request.CreateRoleRequest{
		Role:       roleDomain.Role,
		Visibility: roleDomain.Visibility,
	}

	updateRoleDTO = request.UpdateRoleRequest{
		Role:       roleDomain.Role,
		Visibility: roleDomain.Visibility,
	}

	ctx = context.Background()

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		roleRepository.Mock.On("Save", ctx, mock.Anything).Return(nil).Once()

		err := roleService.Create(ctx, createRoleDTO)

		assert.NoError(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		roleRepository.Mock.On("Save", ctx, mock.Anything).Return(errors.New("error occurred")).Once()

		err := roleService.Create(ctx, createRoleDTO)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		roleRepository.Mock.On("FindById", ctx, roleDomain.ID).Return(roleDomain, nil).Once()

		roleRepository.Mock.On("Update", ctx, updateRoleDTO.ToDomainRole(), roleDomain.ID).Return(nil).Once()

		err := roleService.Update(ctx, updateRoleDTO, roleDomain.ID)

		assert.NoError(t, err)
	})

	t.Run("Failed | Role not found", func(t *testing.T) {
		roleRepository.Mock.On("FindById", ctx, roleDomain.ID).Return(domain.Role{}, utils.ErrRoleNotFound).Once()

		err := roleService.Update(ctx, updateRoleDTO, roleDomain.ID)

		assert.Error(t, err)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		roleRepository.Mock.On("FindById", ctx, roleDomain.ID).Return(roleDomain, nil).Once()

		roleRepository.Mock.On("Update", ctx, updateRoleDTO.ToDomainRole(), roleDomain.ID).Return(errors.New("error occurred")).Once()

		err := roleService.Update(ctx, updateRoleDTO, roleDomain.ID)

		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		roleRepository.Mock.On("FindAll", ctx, "").Return([]domain.Role{roleDomain}, nil).Once()

		results, err := roleService.FindAll(ctx, "")

		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})

	t.Run("Failed | Error occurred", func(t *testing.T) {
		roleRepository.Mock.On("FindAll", ctx, "").Return([]domain.Role{}, errors.New("error occurred")).Once()

		results, err := roleService.FindAll(ctx, "")

		assert.Error(t, err)
		assert.Empty(t, results)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		roleRepository.Mock.On("FindById", ctx, roleDomain.ID).Return(roleDomain, nil).Once()

		result, err := roleService.FindById(ctx, roleDomain.ID)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Failed | Role not found", func(t *testing.T) {
		roleRepository.Mock.On("FindById", ctx, roleDomain.ID).Return(domain.Role{}, utils.ErrRoleNotFound).Once()

		result, err := roleService.FindById(ctx, roleDomain.ID)

		assert.Error(t, err)
		assert.Empty(t, result)
	})
}

// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/arvinpaundra/repository-api/models/domain"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	request "github.com/arvinpaundra/repository-api/models/web/repository/request"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, repositoryId
func (_m *Repository) Delete(ctx context.Context, repositoryId string) error {
	ret := _m.Called(ctx, repositoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, repositoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, query, limit, offset
func (_m *Repository) FindAll(ctx context.Context, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByAuthorId provides a mock function with given fields: ctx, pemustakaId, query, limit, offset
func (_m *Repository) FindByAuthorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, pemustakaId, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByCollectionId provides a mock function with given fields: ctx, collectionId, query, limit, offset
func (_m *Repository) FindByCollectionId(ctx context.Context, collectionId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, collectionId, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, collectionId, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, collectionId, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, collectionId, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByDepartementId provides a mock function with given fields: ctx, departementId, query, limit, offset
func (_m *Repository) FindByDepartementId(ctx context.Context, departementId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, departementId, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, departementId, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, departementId, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, departementId, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByExaminerId provides a mock function with given fields: ctx, pemustakaId, query, limit, offset
func (_m *Repository) FindByExaminerId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, pemustakaId, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: ctx, repositoryId
func (_m *Repository) FindById(ctx context.Context, repositoryId string) (domain.Repository, error) {
	ret := _m.Called(ctx, repositoryId)

	var r0 domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Repository); ok {
		r0 = rf(ctx, repositoryId)
	} else {
		r0 = ret.Get(0).(domain.Repository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repositoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByMentorId provides a mock function with given fields: ctx, pemustakaId, query, limit, offset
func (_m *Repository) FindByMentorId(ctx context.Context, pemustakaId string, query request.RepositoryRequestQuery, limit int, offset int) ([]domain.Repository, int, error) {
	ret := _m.Called(ctx, pemustakaId, query, limit, offset)

	var r0 []domain.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, request.RepositoryRequestQuery, int, int) []domain.Repository); ok {
		r0 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Repository)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, request.RepositoryRequestQuery, int, int) int); ok {
		r1 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, request.RepositoryRequestQuery, int, int) error); ok {
		r2 = rf(ctx, pemustakaId, query, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Save provides a mock function with given fields: ctx, tx, _a2
func (_m *Repository) Save(ctx context.Context, tx *gorm.DB, _a2 domain.Repository) error {
	ret := _m.Called(ctx, tx, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.Repository) error); ok {
		r0 = rf(ctx, tx, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, tx, _a2, repositoryId
func (_m *Repository) Update(ctx context.Context, tx *gorm.DB, _a2 domain.Repository, repositoryId string) error {
	ret := _m.Called(ctx, tx, _a2, repositoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.Repository, string) error); ok {
		r0 = rf(ctx, tx, _a2, repositoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
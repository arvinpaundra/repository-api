// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/arvinpaundra/repository-api/models/domain"
	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, categoryId
func (_m *CategoryRepository) Delete(ctx context.Context, categoryId string) error {
	ret := _m.Called(ctx, categoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, categoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, keyword, limit, offset
func (_m *CategoryRepository) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]domain.Category, int64, error) {
	ret := _m.Called(ctx, keyword, limit, offset)

	var r0 []domain.Category
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []domain.Category); ok {
		r0 = rf(ctx, keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Category)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) int64); ok {
		r1 = rf(ctx, keyword, limit, offset)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int, int) error); ok {
		r2 = rf(ctx, keyword, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindById provides a mock function with given fields: ctx, categoryId
func (_m *CategoryRepository) FindById(ctx context.Context, categoryId string) (domain.Category, error) {
	ret := _m.Called(ctx, categoryId)

	var r0 domain.Category
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Category); ok {
		r0 = rf(ctx, categoryId)
	} else {
		r0 = ret.Get(0).(domain.Category)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *CategoryRepository) Save(ctx context.Context, _a1 domain.Category) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Category) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, _a1, categoryId
func (_m *CategoryRepository) Update(ctx context.Context, _a1 domain.Category, categoryId string) error {
	ret := _m.Called(ctx, _a1, categoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Category, string) error); ok {
		r0 = rf(ctx, _a1, categoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCategoryRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCategoryRepository creates a new instance of CategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCategoryRepository(t mockConstructorTestingTNewCategoryRepository) *CategoryRepository {
	mock := &CategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
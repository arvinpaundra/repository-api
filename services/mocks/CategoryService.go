// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	request "github.com/arvinpaundra/repository-api/models/web/category/request"
	mock "github.com/stretchr/testify/mock"

	response "github.com/arvinpaundra/repository-api/models/web/category/response"
)

// CategoryService is an autogenerated mock type for the CategoryService type
type CategoryService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, category
func (_m *CategoryService) Create(ctx context.Context, category request.CreateCategoryRequest) error {
	ret := _m.Called(ctx, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateCategoryRequest) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, categoryId
func (_m *CategoryService) Delete(ctx context.Context, categoryId string) error {
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
func (_m *CategoryService) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.CategoryResponse, int, int, error) {
	ret := _m.Called(ctx, keyword, limit, offset)

	var r0 []response.CategoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []response.CategoryResponse); ok {
		r0 = rf(ctx, keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.CategoryResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) int); ok {
		r1 = rf(ctx, keyword, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(context.Context, string, int, int) int); ok {
		r2 = rf(ctx, keyword, limit, offset)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, string, int, int) error); ok {
		r3 = rf(ctx, keyword, limit, offset)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// FindById provides a mock function with given fields: ctx, categoryId
func (_m *CategoryService) FindById(ctx context.Context, categoryId string) (response.CategoryResponse, error) {
	ret := _m.Called(ctx, categoryId)

	var r0 response.CategoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) response.CategoryResponse); ok {
		r0 = rf(ctx, categoryId)
	} else {
		r0 = ret.Get(0).(response.CategoryResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, category, categoryId
func (_m *CategoryService) Update(ctx context.Context, category request.UpdateCategoryRequest, categoryId string) error {
	ret := _m.Called(ctx, category, categoryId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UpdateCategoryRequest, string) error); ok {
		r0 = rf(ctx, category, categoryId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCategoryService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCategoryService creates a new instance of CategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCategoryService(t mockConstructorTestingTNewCategoryService) *CategoryService {
	mock := &CategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
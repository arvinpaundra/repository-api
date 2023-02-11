// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	request "github.com/arvinpaundra/repository-api/models/web/departement/request"

	response "github.com/arvinpaundra/repository-api/models/web/departement/response"
)

// DepartementService is an autogenerated mock type for the DepartementService type
type DepartementService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *DepartementService) Create(ctx context.Context, _a1 request.CreateDepartementRequest) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateDepartementRequest) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, departementId
func (_m *DepartementService) Delete(ctx context.Context, departementId string) error {
	ret := _m.Called(ctx, departementId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, departementId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx, keyword, limit, offset
func (_m *DepartementService) FindAll(ctx context.Context, keyword string, limit int, offset int) ([]response.DepartementResponse, int, int, error) {
	ret := _m.Called(ctx, keyword, limit, offset)

	var r0 []response.DepartementResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) []response.DepartementResponse); ok {
		r0 = rf(ctx, keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.DepartementResponse)
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

// FindById provides a mock function with given fields: ctx, departementId
func (_m *DepartementService) FindById(ctx context.Context, departementId string) (response.DepartementResponse, error) {
	ret := _m.Called(ctx, departementId)

	var r0 response.DepartementResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) response.DepartementResponse); ok {
		r0 = rf(ctx, departementId)
	} else {
		r0 = ret.Get(0).(response.DepartementResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, departementId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByProgramStudyId provides a mock function with given fields: ctx, studyProgramId
func (_m *DepartementService) FindByProgramStudyId(ctx context.Context, studyProgramId string) ([]response.DepartementResponse, error) {
	ret := _m.Called(ctx, studyProgramId)

	var r0 []response.DepartementResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) []response.DepartementResponse); ok {
		r0 = rf(ctx, studyProgramId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.DepartementResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, studyProgramId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1, departementId
func (_m *DepartementService) Update(ctx context.Context, _a1 request.UpdateDepartementRequest, departementId string) error {
	ret := _m.Called(ctx, _a1, departementId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UpdateDepartementRequest, string) error); ok {
		r0 = rf(ctx, _a1, departementId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDepartementService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDepartementService creates a new instance of DepartementService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDepartementService(t mockConstructorTestingTNewDepartementService) *DepartementService {
	mock := &DepartementService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
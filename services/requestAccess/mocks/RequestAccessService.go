// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	request "github.com/arvinpaundra/repository-api/models/web/requestAccess/request"
	mock "github.com/stretchr/testify/mock"

	response "github.com/arvinpaundra/repository-api/models/web/requestAccess/response"
)

// RequestAccessService is an autogenerated mock type for the RequestAccessService type
type RequestAccessService struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: ctx, keyword, status, limit, offset
func (_m *RequestAccessService) FindAll(ctx context.Context, keyword string, status string, limit int, offset int) ([]response.RequestAccessResponse, int, int, error) {
	ret := _m.Called(ctx, keyword, status, limit, offset)

	var r0 []response.RequestAccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []response.RequestAccessResponse); ok {
		r0 = rf(ctx, keyword, status, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.RequestAccessResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) int); ok {
		r1 = rf(ctx, keyword, status, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int, int) int); ok {
		r2 = rf(ctx, keyword, status, limit, offset)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, string, string, int, int) error); ok {
		r3 = rf(ctx, keyword, status, limit, offset)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// FindById provides a mock function with given fields: ctx, requestAccessId
func (_m *RequestAccessService) FindById(ctx context.Context, requestAccessId string) (response.RequestAccessResponse, error) {
	ret := _m.Called(ctx, requestAccessId)

	var r0 response.RequestAccessResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) response.RequestAccessResponse); ok {
		r0 = rf(ctx, requestAccessId)
	} else {
		r0 = ret.Get(0).(response.RequestAccessResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, requestAccessId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, requestAccessDTO, requestAccessId
func (_m *RequestAccessService) Update(ctx context.Context, requestAccessDTO request.UpdateRequestAccessRequest, requestAccessId string) error {
	ret := _m.Called(ctx, requestAccessDTO, requestAccessId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.UpdateRequestAccessRequest, string) error); ok {
		r0 = rf(ctx, requestAccessDTO, requestAccessId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRequestAccessService interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequestAccessService creates a new instance of RequestAccessService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequestAccessService(t mockConstructorTestingTNewRequestAccessService) *RequestAccessService {
	mock := &RequestAccessService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
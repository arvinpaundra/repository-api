// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"
)

// Cloudinary is an autogenerated mock type for the Cloudinary type
type Cloudinary struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, filename
func (_m *Cloudinary) Delete(ctx context.Context, filename string) error {
	ret := _m.Called(ctx, filename)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, folder, filename, file
func (_m *Cloudinary) Upload(ctx context.Context, folder string, filename string, file *multipart.FileHeader) (string, error) {
	ret := _m.Called(ctx, folder, filename, file)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *multipart.FileHeader) string); ok {
		r0 = rf(ctx, folder, filename, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *multipart.FileHeader) error); ok {
		r1 = rf(ctx, folder, filename, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCloudinary interface {
	mock.TestingT
	Cleanup(func())
}

// NewCloudinary creates a new instance of Cloudinary. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCloudinary(t mockConstructorTestingTNewCloudinary) *Cloudinary {
	mock := &Cloudinary{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

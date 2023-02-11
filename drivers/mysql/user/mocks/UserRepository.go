// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/arvinpaundra/repository-api/models/domain"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	ret := _m.Called(ctx, email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, tx, _a2
func (_m *UserRepository) Save(ctx context.Context, tx *gorm.DB, _a2 domain.User) error {
	ret := _m.Called(ctx, tx, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, domain.User) error); ok {
		r0 = rf(ctx, tx, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, _a1, email
func (_m *UserRepository) Update(ctx context.Context, _a1 domain.User, email string) error {
	ret := _m.Called(ctx, _a1, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User, string) error); ok {
		r0 = rf(ctx, _a1, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, username, pasword
func (_m *UserUsecase) Login(ctx context.Context, username string, pasword string) (*domain.User, string, error) {
	ret := _m.Called(ctx, username, pasword)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *domain.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*domain.User, string, error)); ok {
		return rf(ctx, username, pasword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.User); ok {
		r0 = rf(ctx, username, pasword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = rf(ctx, username, pasword)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, username, pasword)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Promote provides a mock function with given fields: ctx, id
func (_m *UserUsecase) Promote(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Promote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignUp provides a mock function with given fields: ctx, user
func (_m *UserUsecase) SignUp(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) (*domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

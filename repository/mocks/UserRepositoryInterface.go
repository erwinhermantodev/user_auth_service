// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/erwinhermantodev/user_auth_service/repository"
	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) CreateUser(user *repository.User) (int, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*repository.User) (int, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*repository.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*repository.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByPhone provides a mock function with given fields: phone
func (_m *UserRepositoryInterface) FindUserByPhone(phone string) (*repository.User, error) {
	ret := _m.Called(phone)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByPhone")
	}

	var r0 *repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.User, error)); ok {
		return rf(phone)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.User); ok {
		r0 = rf(phone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: userID
func (_m *UserRepositoryInterface) GetUserByID(userID int) (*repository.User, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *repository.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*repository.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) *repository.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementLoginCount provides a mock function with given fields: userID
func (_m *UserRepositoryInterface) IncrementLoginCount(userID int) error {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for IncrementLoginCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserProfile provides a mock function with given fields: user
func (_m *UserRepositoryInterface) UpdateUserProfile(user *repository.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
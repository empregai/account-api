// Code generated by mockery v2.20.0. DO NOT EDIT.

package sessionmock

import (
	context "context"
	session "go-api/internal/core/session"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CreateSession provides a mock function with given fields: ctx, userID
func (_m *UseCase) CreateSession(ctx context.Context, userID uuid.UUID) (string, error) {
	ret := _m.Called(ctx, userID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (string, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) string); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, sessionID
func (_m *UseCase) DeleteByID(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSessionByID provides a mock function with given fields: ctx, sessionID
func (_m *UseCase) GetSessionByID(ctx context.Context, sessionID string) (*session.Session, error) {
	ret := _m.Called(ctx, sessionID)

	var r0 *session.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*session.Session, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *session.Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*session.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

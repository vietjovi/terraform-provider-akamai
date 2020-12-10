// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package iam

import (
	context "context"

	iam "github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/iam"
	mock "github.com/stretchr/testify/mock"
)

// IAM is an autogenerated mock type for the IAM type
type IAM struct {
	mock.Mock
}

// ListGroups provides a mock function with given fields: _a0, _a1
func (_m *IAM) ListGroups(_a0 context.Context, _a1 iam.ListGroupsRequest) ([]iam.Group, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []iam.Group
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListGroupsRequest) []iam.Group); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, iam.ListGroupsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: _a0
func (_m *IAM) ListProducts(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRoles provides a mock function with given fields: _a0, _a1
func (_m *IAM) ListRoles(_a0 context.Context, _a1 iam.ListRolesRequest) ([]iam.Role, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []iam.Role
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListRolesRequest) []iam.Role); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, iam.ListRolesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStates provides a mock function with given fields: _a0, _a1
func (_m *IAM) ListStates(_a0 context.Context, _a1 iam.ListStatesRequest) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, iam.ListStatesRequest) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, iam.ListStatesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTimeoutPolicies provides a mock function with given fields: _a0
func (_m *IAM) ListTimeoutPolicies(_a0 context.Context) ([]iam.TimeoutPolicy, error) {
	ret := _m.Called(_a0)

	var r0 []iam.TimeoutPolicy
	if rf, ok := ret.Get(0).(func(context.Context) []iam.TimeoutPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]iam.TimeoutPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedContactTypes provides a mock function with given fields: _a0
func (_m *IAM) SupportedContactTypes(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedCountries provides a mock function with given fields: _a0
func (_m *IAM) SupportedCountries(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedLanguages provides a mock function with given fields: _a0
func (_m *IAM) SupportedLanguages(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

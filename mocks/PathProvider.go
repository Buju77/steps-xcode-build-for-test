// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PathProvider is an autogenerated mock type for the PathProvider type
type PathProvider struct {
	mock.Mock
}

// CreateTempDir provides a mock function with given fields: prefix
func (_m *PathProvider) CreateTempDir(prefix string) (string, error) {
	ret := _m.Called(prefix)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

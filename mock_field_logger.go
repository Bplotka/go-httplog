package httplog

import "github.com/stretchr/testify/mock"

// MockFieldLogger is an autogenerated mock type for the FieldLogger type
type MockFieldLogger struct {
	mock.Mock
}

// Log provides a mock function with given fields: args
func (_m *MockFieldLogger) Log(args ...interface{}) {
	_m.Called(args)
}

// WithError provides a mock function with given fields: err
func (_m *MockFieldLogger) WithError(err error) FieldLogger {
	ret := _m.Called(err)

	var r0 FieldLogger
	if rf, ok := ret.Get(0).(func(error) FieldLogger); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(FieldLogger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: fields
func (_m *MockFieldLogger) WithFields(fields Fields) FieldLogger {
	ret := _m.Called(fields)

	var r0 FieldLogger
	if rf, ok := ret.Get(0).(func(Fields) FieldLogger); ok {
		r0 = rf(fields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(FieldLogger)
		}
	}

	return r0
}

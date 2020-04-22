// Code generated by mockery v1.0.0. DO NOT EDIT.

package amqpwrapper

import (
	amqp "github.com/streadway/amqp"
	mock "github.com/stretchr/testify/mock"
)

// MockIConnectionManager is an autogenerated mock type for the IConnectionManager type
type MockIConnectionManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockIConnectionManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateChannel provides a mock function with given fields: typeChan
func (_m *MockIConnectionManager) CreateChannel(typeChan uint64) (*amqp.Channel, error) {
	ret := _m.Called(typeChan)

	var r0 *amqp.Channel
	if rf, ok := ret.Get(0).(func(uint64) *amqp.Channel); ok {
		r0 = rf(typeChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(typeChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannel provides a mock function with given fields: key, typeChan
func (_m *MockIConnectionManager) GetChannel(key string, typeChan uint64) (*amqp.Channel, error) {
	ret := _m.Called(key, typeChan)

	var r0 *amqp.Channel
	if rf, ok := ret.Get(0).(func(string, uint64) *amqp.Channel); ok {
		r0 = rf(key, typeChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(key, typeChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChannel provides a mock function with given fields: fn, args
func (_m *MockIConnectionManager) InitChannel(fn InitializeChannel, args InitArgs) error {
	ret := _m.Called(fn, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(InitializeChannel, InitArgs) error); ok {
		r0 = rf(fn, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitChannelAndGet provides a mock function with given fields: fn, args
func (_m *MockIConnectionManager) InitChannelAndGet(fn InitializeChannel, args InitArgs) (*amqp.Channel, error) {
	ret := _m.Called(fn, args)

	var r0 *amqp.Channel
	if rf, ok := ret.Get(0).(func(InitializeChannel, InitArgs) *amqp.Channel); ok {
		r0 = rf(fn, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(InitializeChannel, InitArgs) error); ok {
		r1 = rf(fn, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsClosed provides a mock function with given fields:
func (_m *MockIConnectionManager) IsClosed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
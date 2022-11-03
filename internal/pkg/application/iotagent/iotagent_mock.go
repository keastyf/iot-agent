// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package iotagent

import (
	"context"
	app "github.com/diwise/iot-agent/internal/pkg/application"
	"sync"
)

// Ensure, that IoTAgentMock does implement IoTAgent.
// If this is not the case, regenerate this file with moq.
var _ IoTAgent = &IoTAgentMock{}

// IoTAgentMock is a mock implementation of IoTAgent.
//
//	func TestSomethingThatUsesIoTAgent(t *testing.T) {
//
//		// make and configure a mocked IoTAgent
//		mockedIoTAgent := &IoTAgentMock{
//			MessageReceivedFunc: func(ctx context.Context, ue app.SensorEvent) error {
//				panic("mock out the MessageReceived method")
//			},
//			MessageReceivedFnFunc: func(ctx context.Context, msg []byte, ue app.UplinkASFunc) error {
//				panic("mock out the MessageReceivedFn method")
//			},
//		}
//
//		// use mockedIoTAgent in code that requires IoTAgent
//		// and then make assertions.
//
//	}
type IoTAgentMock struct {
	// MessageReceivedFunc mocks the MessageReceived method.
	MessageReceivedFunc func(ctx context.Context, ue app.SensorEvent) error

	// MessageReceivedFnFunc mocks the MessageReceivedFn method.
	MessageReceivedFnFunc func(ctx context.Context, msg []byte, ue app.UplinkASFunc) error

	// calls tracks calls to the methods.
	calls struct {
		// MessageReceived holds details about calls to the MessageReceived method.
		MessageReceived []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Ue is the ue argument value.
			Ue app.SensorEvent
		}
		// MessageReceivedFn holds details about calls to the MessageReceivedFn method.
		MessageReceivedFn []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg []byte
			// Ue is the ue argument value.
			Ue app.UplinkASFunc
		}
	}
	lockMessageReceived   sync.RWMutex
	lockMessageReceivedFn sync.RWMutex
}

// MessageReceived calls MessageReceivedFunc.
func (mock *IoTAgentMock) MessageReceived(ctx context.Context, ue app.SensorEvent) error {
	if mock.MessageReceivedFunc == nil {
		panic("IoTAgentMock.MessageReceivedFunc: method is nil but IoTAgent.MessageReceived was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Ue  app.SensorEvent
	}{
		Ctx: ctx,
		Ue:  ue,
	}
	mock.lockMessageReceived.Lock()
	mock.calls.MessageReceived = append(mock.calls.MessageReceived, callInfo)
	mock.lockMessageReceived.Unlock()
	return mock.MessageReceivedFunc(ctx, ue)
}

// MessageReceivedCalls gets all the calls that were made to MessageReceived.
// Check the length with:
//
//	len(mockedIoTAgent.MessageReceivedCalls())
func (mock *IoTAgentMock) MessageReceivedCalls() []struct {
	Ctx context.Context
	Ue  app.SensorEvent
} {
	var calls []struct {
		Ctx context.Context
		Ue  app.SensorEvent
	}
	mock.lockMessageReceived.RLock()
	calls = mock.calls.MessageReceived
	mock.lockMessageReceived.RUnlock()
	return calls
}

// MessageReceivedFn calls MessageReceivedFnFunc.
func (mock *IoTAgentMock) MessageReceivedFn(ctx context.Context, msg []byte, ue app.UplinkASFunc) error {
	if mock.MessageReceivedFnFunc == nil {
		panic("IoTAgentMock.MessageReceivedFnFunc: method is nil but IoTAgent.MessageReceivedFn was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Msg []byte
		Ue  app.UplinkASFunc
	}{
		Ctx: ctx,
		Msg: msg,
		Ue:  ue,
	}
	mock.lockMessageReceivedFn.Lock()
	mock.calls.MessageReceivedFn = append(mock.calls.MessageReceivedFn, callInfo)
	mock.lockMessageReceivedFn.Unlock()
	return mock.MessageReceivedFnFunc(ctx, msg, ue)
}

// MessageReceivedFnCalls gets all the calls that were made to MessageReceivedFn.
// Check the length with:
//
//	len(mockedIoTAgent.MessageReceivedFnCalls())
func (mock *IoTAgentMock) MessageReceivedFnCalls() []struct {
	Ctx context.Context
	Msg []byte
	Ue  app.UplinkASFunc
} {
	var calls []struct {
		Ctx context.Context
		Msg []byte
		Ue  app.UplinkASFunc
	}
	mock.lockMessageReceivedFn.RLock()
	calls = mock.calls.MessageReceivedFn
	mock.lockMessageReceivedFn.RUnlock()
	return calls
}

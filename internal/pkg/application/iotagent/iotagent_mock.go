// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package iotagent

import (
	"context"
	"github.com/diwise/iot-agent/internal/pkg/application"
	"github.com/farshidtz/senml/v2"
	"sync"
)

// Ensure, that AppMock does implement App.
// If this is not the case, regenerate this file with moq.
var _ App = &AppMock{}

// AppMock is a mock implementation of App.
//
//	func TestSomethingThatUsesApp(t *testing.T) {
//
//		// make and configure a mocked App
//		mockedApp := &AppMock{
//			HandleSensorEventFunc: func(ctx context.Context, se application.SensorEvent) error {
//				panic("mock out the HandleSensorEvent method")
//			},
//			HandleSensorMeasurementListFunc: func(ctx context.Context, deviceID string, pack senml.Pack) error {
//				panic("mock out the HandleSensorMeasurementList method")
//			},
//		}
//
//		// use mockedApp in code that requires App
//		// and then make assertions.
//
//	}
type AppMock struct {
	// HandleSensorEventFunc mocks the HandleSensorEvent method.
	HandleSensorEventFunc func(ctx context.Context, se application.SensorEvent) error

	// HandleSensorMeasurementListFunc mocks the HandleSensorMeasurementList method.
	HandleSensorMeasurementListFunc func(ctx context.Context, deviceID string, pack senml.Pack) error

	// calls tracks calls to the methods.
	calls struct {
		// HandleSensorEvent holds details about calls to the HandleSensorEvent method.
		HandleSensorEvent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Se is the se argument value.
			Se application.SensorEvent
		}
		// HandleSensorMeasurementList holds details about calls to the HandleSensorMeasurementList method.
		HandleSensorMeasurementList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Pack is the pack argument value.
			Pack senml.Pack
		}
	}
	lockHandleSensorEvent           sync.RWMutex
	lockHandleSensorMeasurementList sync.RWMutex
}

// HandleSensorEvent calls HandleSensorEventFunc.
func (mock *AppMock) HandleSensorEvent(ctx context.Context, se application.SensorEvent) error {
	if mock.HandleSensorEventFunc == nil {
		panic("AppMock.HandleSensorEventFunc: method is nil but App.HandleSensorEvent was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Se  application.SensorEvent
	}{
		Ctx: ctx,
		Se:  se,
	}
	mock.lockHandleSensorEvent.Lock()
	mock.calls.HandleSensorEvent = append(mock.calls.HandleSensorEvent, callInfo)
	mock.lockHandleSensorEvent.Unlock()
	return mock.HandleSensorEventFunc(ctx, se)
}

// HandleSensorEventCalls gets all the calls that were made to HandleSensorEvent.
// Check the length with:
//
//	len(mockedApp.HandleSensorEventCalls())
func (mock *AppMock) HandleSensorEventCalls() []struct {
	Ctx context.Context
	Se  application.SensorEvent
} {
	var calls []struct {
		Ctx context.Context
		Se  application.SensorEvent
	}
	mock.lockHandleSensorEvent.RLock()
	calls = mock.calls.HandleSensorEvent
	mock.lockHandleSensorEvent.RUnlock()
	return calls
}

// HandleSensorMeasurementList calls HandleSensorMeasurementListFunc.
func (mock *AppMock) HandleSensorMeasurementList(ctx context.Context, deviceID string, pack senml.Pack) error {
	if mock.HandleSensorMeasurementListFunc == nil {
		panic("AppMock.HandleSensorMeasurementListFunc: method is nil but App.HandleSensorMeasurementList was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeviceID string
		Pack     senml.Pack
	}{
		Ctx:      ctx,
		DeviceID: deviceID,
		Pack:     pack,
	}
	mock.lockHandleSensorMeasurementList.Lock()
	mock.calls.HandleSensorMeasurementList = append(mock.calls.HandleSensorMeasurementList, callInfo)
	mock.lockHandleSensorMeasurementList.Unlock()
	return mock.HandleSensorMeasurementListFunc(ctx, deviceID, pack)
}

// HandleSensorMeasurementListCalls gets all the calls that were made to HandleSensorMeasurementList.
// Check the length with:
//
//	len(mockedApp.HandleSensorMeasurementListCalls())
func (mock *AppMock) HandleSensorMeasurementListCalls() []struct {
	Ctx      context.Context
	DeviceID string
	Pack     senml.Pack
} {
	var calls []struct {
		Ctx      context.Context
		DeviceID string
		Pack     senml.Pack
	}
	mock.lockHandleSensorMeasurementList.RLock()
	calls = mock.calls.HandleSensorMeasurementList
	mock.lockHandleSensorMeasurementList.RUnlock()
	return calls
}

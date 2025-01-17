package iotagent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/diwise/iot-agent/internal/pkg/application"
	"github.com/diwise/iot-agent/internal/pkg/application/conversion"
	"github.com/diwise/iot-agent/internal/pkg/application/decoder"
	"github.com/diwise/iot-agent/internal/pkg/application/decoder/payload"
	"github.com/diwise/iot-agent/internal/pkg/application/events"
	"github.com/diwise/iot-agent/internal/pkg/application/messageprocessor"
	core "github.com/diwise/iot-core/pkg/messaging/events"
	dmc "github.com/diwise/iot-device-mgmt/pkg/client"
	"github.com/diwise/service-chassis/pkg/infrastructure/o11y/logging"
	"github.com/farshidtz/senml/v2"
)

//go:generate moq -rm -out iotagent_mock.go . App

type App interface {
	HandleSensorEvent(ctx context.Context, se application.SensorEvent) error
	HandleSensorMeasurementList(ctx context.Context, deviceID string, pack senml.Pack) error
}

type app struct {
	messageProcessor       messageprocessor.MessageProcessor
	decoderRegistry        decoder.DecoderRegistry
	deviceManagementClient dmc.DeviceManagementClient
	eventSender            events.EventSender

	notFoundDevices   map[string]time.Time
	notFoundDevicesMu sync.Mutex
}

func New(dmc dmc.DeviceManagementClient, eventPub events.EventSender) App {
	c := conversion.NewConverterRegistry()
	d := decoder.NewDecoderRegistry()
	m := messageprocessor.NewMessageReceivedProcessor(c)

	return &app{
		messageProcessor:       m,
		decoderRegistry:        d,
		deviceManagementClient: dmc,
		eventSender:            eventPub,
		notFoundDevices:        make(map[string]time.Time),
	}
}

func (a *app) HandleSensorEvent(ctx context.Context, se application.SensorEvent) error {
	log := logging.GetFromContext(ctx).With().Str("devEui", se.DevEui).Logger()
	ctx = logging.NewContextWithLogger(ctx, log)

	device, err := a.findDevice(ctx, se.DevEui)
	if err != nil {
		if errors.Is(err, errDeviceOnBlackList) {
			log.Warn().Str("deviceName", se.DeviceName).Msg("blacklisted")
			return nil
		}

		return err
	}

	log = log.With().Str("device", device.ID()).Logger()
	ctx = logging.NewContextWithLogger(ctx, log)

	log.Debug().Str("type", device.SensorType()).Msg("message received")

	decodePayload := decoder.PayloadErrorDecoder
	if !se.HasError() {
		decodePayload = a.decoderRegistry.GetDecoderForSensorType(ctx, device.SensorType())
	}

	err = decodePayload(ctx, se, func(ctx context.Context, p payload.Payload) error {
		a.sendStatusMessage(ctx, device.ID(), device.Tenant(), p)

		packs, err := a.messageProcessor.ProcessMessage(ctx, p, device)
		if err != nil {
			return fmt.Errorf("failed to process message (%w)", err)
		}

		if device.IsActive() {
			for _, pack := range packs {
				a.HandleSensorMeasurementList(ctx, device.ID(), pack)
			}
		} else {
			log.Warn().Msg("ignored message from inactive device")
		}

		return nil
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to handle received message")
	}

	return err
}

func (a *app) HandleSensorMeasurementList(ctx context.Context, deviceID string, pack senml.Pack) error {
	m := core.MessageReceived{
		Device:    deviceID,
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		Pack:      pack,
	}

	a.eventSender.Send(ctx, &m)

	return nil
}

var errDeviceOnBlackList = errors.New("blacklisted")

func (a *app) deviceIsCurrentlyIgnored(ctx context.Context, devEui string) bool {
	a.notFoundDevicesMu.Lock()
	defer a.notFoundDevicesMu.Unlock()

	if timeOfNextAllowedRetry, ok := a.notFoundDevices[devEui]; ok {
		if !time.Now().UTC().After(timeOfNextAllowedRetry) {
			return true
		}

		delete(a.notFoundDevices, devEui)
	}

	return false
}

func (a *app) findDevice(ctx context.Context, devEui string) (dmc.Device, error) {
	if a.deviceIsCurrentlyIgnored(ctx, devEui) {
		return nil, errDeviceOnBlackList
	}

	device, err := a.deviceManagementClient.FindDeviceFromDevEUI(ctx, devEui)
	if err != nil {
		a.ignoreDeviceFor(ctx, devEui, 1*time.Hour)
		return nil, fmt.Errorf("device lookup failure (%w)", err)
	}

	return device, nil
}

func (a *app) ignoreDeviceFor(ctx context.Context, devEui string, period time.Duration) {
	a.notFoundDevicesMu.Lock()
	defer a.notFoundDevicesMu.Unlock()

	a.notFoundDevices[devEui] = time.Now().UTC().Add(period)
}

func (a *app) sendStatusMessage(ctx context.Context, deviceID, tenant string, p payload.Payload) {
	logger := logging.GetFromContext(ctx).
		With().Str("device_id", deviceID).Logger().
		With().Str("func", "sendStatusMessage").
		Logger()
		
	var decorators []func(*events.StatusMessage)

	decorators = append(decorators,
		events.WithStatus(p.Status().Code, p.Status().Messages),
		events.WithTenant(tenant),
	)

	if bat, ok := payload.Get[int](p, payload.BatteryLevelProperty); ok {
		decorators = append(decorators, events.WithBatteryLevel(bat))
	}

	statusMsg := events.NewStatusMessage(deviceID, decorators...)

	if statusMsg.Tenant == "" {
		logger.Warn().Msg("tenant information is missing")
	}

	err := a.eventSender.Publish(ctx, statusMsg)
	if err != nil {
		logger.Error().Err(err).Msg("failed to publish status message")
	}
}

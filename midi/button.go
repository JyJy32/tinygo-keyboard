package midi

import (
	"errors"
	"machine"
	"machine/usb/adc/midi"
	"time"
)

var debounce time.Duration = time.Microsecond * 8

type MidiButton interface {
	OnTick() error
	interruptR()
	interruptF()
}

type Callback func(pin machine.Pin, channel uint8, controller uint8)

type MidiControlButton struct {
	pin        machine.Pin
	channel    uint8
	controller uint8 // make a controller parent struct?
	value      uint8
	r_value    uint8 // default to 0

	pressed   bool
	released  bool
	lastPress time.Time
	onDown    Callback
}

// NewMidiControlButton create new midi control button
// params:
//
//	pin machine.Pin   the gpio pin on the microcontroller
//	channel uint8     [[1;16]] value defining the channel
//	controller        [[0;127]] value defining the controller
//	value             [[0;127]] value to send
func NewMidiControlButton(pin machine.Pin, channel uint8, controller uint8, value uint8) *MidiControlButton {
	b := new(MidiControlButton)
	b.pin = pin
	b.channel = channel
	b.controller = controller
	b.value = value
	b.r_value = 0
	return b
}

func (b *MidiControlButton) Init() *MidiControlButton {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	b.pin.SetInterrupt(machine.PinRising, b.interruptR)
	return b
}

func (b *MidiControlButton) OnTick() error {
	if b.pressed {
		b.onDownCallback()
		b.pressed = false
	} else if b.released {

	}
	return nil
}

func (b *MidiControlButton) onDownCallback() {
	if b.onDown != nil {
		b.onDown(b.pin, b.channel, b.controller)
	} else {
		midi.Port().ControlChange(0, b.channel, b.controller, b.value)
	}
}

func (b *MidiControlButton) SetOnDown(fn Callback) *MidiControlButton {
	b.onDown = fn
	return b
}

func (b *MidiControlButton) interruptR(pin machine.Pin) {
	now := time.Now()
	if now.Sub(b.lastPress) > debounce {
		b.pressed = true
		b.lastPress = now
	}
}

func (b *MidiControlButton) interruptF(pin machine.Pin) {
	now := time.Now()
	if now.Sub(b.lastPress) > debounce {
		b.released = true
		b.lastPress = now
	}
}

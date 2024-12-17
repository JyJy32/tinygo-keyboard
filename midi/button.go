package midi

import (
	"machine"
	"machine/usb/adc/midi"
	"time"
)

var debounce time.Duration = time.Millisecond * 80

type Callback func(pin machine.Pin, channel uint8, controller uint8)

// MidiControlButton
// fields:
//
//	pin         machine.Pin the pin the button is connected to
//	channel     uint8       midi channel to send message, only value in [[1;16]] possible
//	controller  uint8       control to send
//	value       uint8       value to send on button down
//	r_value     uint8       value to send on button up
type MidiControlButton struct {
	pin        machine.Pin
	channel    uint8
	controller uint8
	value      uint8
	r_value    uint8 // default to 0

	pressed      bool
	released     bool
	lastPress    time.Time
	lastReleased time.Time
	onDown       Callback
	onUp         Callback
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
	b.pin.SetInterrupt(machine.PinToggle, b.interrupt)

	return b
}

func (b *MidiControlButton) OnTick() error {
	if b.pressed {
		b.onDownCallback()
		b.pressed = false
	}
	if b.released {
		b.onUpCallback()
		b.released = false
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

func (b *MidiControlButton) onUpCallback() {
	if b.onUp != nil {
		b.onUp(b.pin, b.channel, b.controller)
	} else {
		midi.Port().ControlChange(0, b.channel, b.controller, b.r_value)
	}
}

func (b *MidiControlButton) SetOnUp(fn Callback) *MidiControlButton {
	b.onUp = fn
	return b
}

func (b *MidiControlButton) interrupt(pin machine.Pin) {
	now := time.Now()
	print("pressed")
	if now.Sub(b.lastPress) > debounce {
		state := pin.Get()
		b.pressed = state
		b.released = !state
		b.lastPress = now
	}
}

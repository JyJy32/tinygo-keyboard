package midi

import (
	"errors"
	"machine"
	"machine/usb/adc/midi"
)

var (
	debounce uint8 = 8
)

type MidiControlButton struct {
	pin        machine.Pin
	channel    uint8
	controller uint8 // make a controller parent struct?
	value      uint8
	r_value    uint8 // default to 0

	// TODO: set callback as value

	pressed bool
	bounce  uint8
}

// NewMidiControlButton create new midi control button
// params:
//
//	pin machine.Pin   the gpio pin on the microcontroller
//	channel uint8     [[1;16]] value defining the channel
//	controller        [[0;127]] value defining the controller
//	value             [[0;127]] value to send
func NewMidiControlButton(pin machine.Pin, channel uint8, controller uint8, value uint8) (*MidiControlButton, error) {
	b := new(MidiControlButton)
	b.pin = pin
	if channel > 16 || channel < 1 {
		return nil, errors.New("invalid channel")
	}
	b.channel = channel
	b.controller = controller
	b.value = value
	b.r_value = 0
	b.bounce = 0
	return b, nil
}

func (b *MidiControlButton) Init() *MidiControlButton {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	return b
}

func (b *MidiControlButton) OnTick() {
	if b.bounce != 0 {
		b.bounce -= 1
		return
	}
	if b.pin.Get() && !b.pressed {
		midi.Port().ControlChange(0, b.channel, b.controller, b.value)
		b.bounce = debounce
		b.pressed = true
	} else if !b.pin.Get() && b.pressed {
		midi.Port().ControlChange(0, b.channel, b.controller, b.r_value)
		b.pressed = false
	}
}

type MidiNoteButton struct {
	pin        machine.Pin
	channel    uint8
	note       uint8
	velocity   uint8
	r_velocity uint8

	pressed bool
	bounce  uint8
}

// NewMidiControlButton create new midi control button
// params:
//
//	pin machine.Pin   the gpio pin on the microcontroller
//	channel uint8     [[1;16]] value defining the channel
//	note              [[0;127]] value to send
//	vel               velocity
func newMidiNoteButton(pin machine.Pin, channel uint8, note uint8, vel uint8) (*MidiNoteButton, error) {
	b := new(MidiNoteButton)
	b.pin = pin
	if channel > 16 || channel < 1 {
		return nil, errors.New("invalid channel")
	}
	b.channel = channel
	b.note = note
	b.velocity = vel

	b.r_velocity = 0
	b.pressed = false
	b.bounce = 0
	return b, nil
}

func (b *MidiNoteButton) init() *MidiNoteButton {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	return b
}

func (b *MidiNoteButton) OnTick() {
	if b.bounce != 0 {
		b.bounce -= 1
		return
	}
	if b.pin.Get() && !b.pressed {
		midi.Port().NoteOn(0, 0, midi.Note(b.note), b.velocity)
		b.bounce = debounce
		b.pressed = true
	} else if !b.pin.Get() && b.pressed {
		midi.Port().NoteOff(0, 0, midi.Note(b.note), b.r_velocity)
		b.pressed = false
	}
}

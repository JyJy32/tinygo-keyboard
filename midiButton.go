package main

import (
	"machine"
	"machine/usb/adc/midi"
)

var (
	debounce uint8 = 8
)

type midiButton struct {
	pin  machine.Pin
	note uint8

	pressed bool
	bounce  uint8
}

func NewMidiButton(pin machine.Pin, note uint8) *midiButton {
	b := new(midiButton)
	b.pin = pin
	b.note = note
	b.bounce = 0
	return b
}

func (b *midiButton) init() *midiButton {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	return b
}

func (b *midiButton) OnTick() {
	if b.bounce != 0 {
		b.bounce -= 1
		return
	}
	if b.pin.Get() && !b.pressed {
		midi.Port().ControlChange(0, 1, 0, 5)
		b.bounce = debounce
		b.pressed = true
	} else if !b.pin.Get() {
		b.pressed = false
	}
}

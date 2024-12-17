package midi

import (
	"errors"
	"machine"
)

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

func (b *MidiNoteButton) OnTick() error {
	return nil
}

package keyboard

import (
	"machine"
	"machine/usb/hid/keyboard"
)

var debounce uint8 = 8

type Button struct {
	pin machine.Pin
	key keyboard.Keycode

	pressed bool
	bounce  uint8
}

func NewButton(pin machine.Pin, key keyboard.Keycode) *Button {
	b := new(Button)
	b.pin = pin
	b.key = key
	b.bounce = 0
	return b
}

func (b *Button) init() *Button {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	return b
}

func (b *Button) OnTick() {
	if b.bounce != 0 {
		b.bounce -= 1
		return
	}
	if b.pin.Get() && !b.pressed {
		keyboard.Port().Press(b.key)
		b.bounce = debounce
		b.pressed = true
	} else if !b.pin.Get() {
		b.pressed = false
	}
}

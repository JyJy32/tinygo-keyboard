package keyboard

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

var debounce time.Duration = time.Millisecond * 80

type Callback func(pin machine.Pin, key keyboard.Keycode)

type Button struct {
	pin machine.Pin
	key keyboard.Keycode

	pressed   bool
	released  bool
	lastPress time.Time
	onDown    Callback
	onUp      Callback
}

func NewButton(pin machine.Pin, key keyboard.Keycode) *Button {
	b := new(Button)
	b.pin = pin
	b.key = key
	return b
}

func (b *Button) Init() *Button {
	b.pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	b.pin.SetInterrupt(machine.PinToggle, b.interrupt)

	return b
}

func (b *Button) OnTick() error {
	if b.pressed {
		b.onDownCallback()
		b.pressed = false
	} else if b.released {
		b.OnUpCallback()
		b.released = false
	}
	return nil
}

func (b *Button) onDownCallback() {
	// if different behavior is defined do that else just press key
	if b.onDown != nil {
		b.onDown(b.pin, b.key)
	} else {
		keyboard.Port().Press(b.key)
	}
}

func (b *Button) SetOnDown(fn Callback) *Button {
	b.onDown = fn
	return b
}

func (b *Button) OnUpCallback() {
	if b.onUp != nil {
		b.onUp(b.pin, b.key)
	}
}

func (b *Button) SetOnUp(fn Callback) *Button {
	b.onUp = fn
	return b
}

func (b *Button) interrupt(pin machine.Pin) {
	now := time.Now()
	if now.Sub(b.lastPress) > debounce {
		state := pin.Get()
		b.pressed = state
		b.released = !state
		b.lastPress = now
	}
}

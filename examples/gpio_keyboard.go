package main

import (
	"machine"
	kb "machine/usb/hid/keyboard"
	"macropad/keyboard"
	"time"
)

func main() {
	button := keyboard.NewButton(machine.GP7, kb.KeyF13).Init().SetOnDown(func(pin machine.Pin, key kb.Keycode) {
		macro := []byte("ur mom")
		kb.Port().Write(macro)
	})
	ticker := time.NewTicker(time.Microsecond * 10)
	for range ticker.C {
		button.OnTick()
	}
}

package main

import (
	"machine"
	kb "machine/usb/hid/keyboard"
	"macropad/keyboard"
	"time"
)

func main() {
	boardLed := machine.LED
	boardLed.Configure(machine.PinConfig{Mode: machine.PinOutput})

	buttons := [...]*keyboard.Button{
		// set sequence of bytes, the key will be ignored
		keyboard.NewButton(machine.GP16, kb.KeyF13).
			Init().
			SetOnDown(func(pin machine.Pin, key kb.Keycode) {
				macro := []byte("what the skibidi")
				kb.Port().Write(macro)
			}),
		// set direct byte value (ascii) to be sent
		keyboard.NewButton(machine.GP17, 0x52).Init(),
		// set actual keycode (beware that these are scan codes and not the value)
		keyboard.NewButton(machine.GP18, kb.KeyEsc).Init(),
		// set a function to run when the button is released
		keyboard.NewButton(machine.GP19, 0x54).
			Init().
			SetOnUp(func(pin machine.Pin, key kb.Keycode) {
				kb.Port().Write([]byte("I am released"))
			}),
		keyboard.NewButton(machine.GP20, 0x55).Init(),
		keyboard.NewButton(machine.GP21, 0x56).Init(),
		keyboard.NewButton(machine.GP22, 0x57).Init(),
		keyboard.NewButton(machine.GP7, 0x58).Init(),
		keyboard.NewButton(machine.GP0, 0x59).Init(),
	}

	board := keyboard.Board{}

	for _, btn := range buttons {
		board.AddButton(btn)
	}

	boardLed.High()
	ticker := time.NewTicker(time.Microsecond * 10)
	for range ticker.C {
		board.OnTick()
	}
}

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
		keyboard.NewButton(machine.GP16, kb.KeyF13).
			Init().
			SetOnDown(func(pin machine.Pin, key kb.Keycode) {
				macro := []byte("what the skibidi")
				kb.Port().Write(macro)
			}),
		keyboard.NewButton(machine.GP17, 0x52).Init(),
		keyboard.NewButton(machine.GP18, 0x53).Init(),
		keyboard.NewButton(machine.GP19, 0x54).Init(),
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

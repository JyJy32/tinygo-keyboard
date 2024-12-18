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

	kb.Port().Press(kb.KeyA)
	buttons := [...]*keyboard.Button{
		// set sequence of bytes, the key will be ignored
		keyboard.NewButton(machine.GP7, kb.KeyF13).
			Init().
			SetOnDown(func(pin machine.Pin, key kb.Keycode) {
				macro := []byte("what the skibidi")
				kb.Port().Write(macro)
				println("HELP2")
			}),
		// set direct byte value (ascii) to be sent
		keyboard.NewButton(machine.GP8, 0x52).Init(),
		// set actual keycode (beware that these are scan codes and not the value)
		keyboard.NewButton(machine.GP9, kb.KeyEsc).Init(),
		// set a function to run when the button is released
		keyboard.NewButton(machine.GP10, 0x54).
			Init().
			SetOnUp(func(pin machine.Pin, key kb.Keycode) {
				kb.Port().Write([]byte("I am released"))
			}),
		keyboard.NewButton(machine.GP11, 0x55).Init(),
		keyboard.NewButton(machine.GP12, 0x56).Init(),
		keyboard.NewButton(machine.GP13, 0x57).Init(),
		keyboard.NewButton(machine.GP14, 0x58).Init(),
		keyboard.NewButton(machine.GP2, 0x59).Init(),
	}

	// TODO: this should maybe be part of the board struct
	encoder := keyboard.NewEncoder(machine.GP0, machine.GP1).
		Init().
		SetonCW(func(pin machine.Pin, key kb.Keycode) {
			kb.Port().Press(kb.KeyMediaVolumeInc)
		}).
		SetonCCW(func(pin machine.Pin, key kb.Keycode) {
			kb.Port().Press(kb.KeyMediaVolumeDec)
		})
	board := keyboard.Board{}

	for _, btn := range buttons {
		board.AddButton(btn)
	}

	boardLed.High()
	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		encoder.OnTick()
		board.OnTick()
	}
}

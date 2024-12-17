package main

import (
	"machine"
	"machine/usb"
	"macropad/midi"
	"time"
)

func main() {
	// led to see that the config worked
	boardLED := machine.LED
	boardLED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	usb.Product = "jyx-controller"

	board := new(midi.Board)

	button_pins := [...]machine.Pin{
		machine.GP7,
		machine.GP8,
		machine.GP9,
		machine.GP10,
		machine.GP11,
		machine.GP12,
		machine.GP13,
		machine.GP14,
		machine.GP2,
	}
	for i := range 9 {
		button := midi.NewMidiControlButton(button_pins[i], 1, uint8(i), 1).
			Init()
		board.AddButton(button)
	}

	boardLED.High()
	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		err := board.OnTick()
		if err != nil {
			break
		}
	}
}

package main

import (
	"machine"
	"machine/usb"
	"macropad/midi"
	"time"
)

func main() {
	boardLED := machine.LED
	boardLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	usb.Product = "jyx-controller"
	// port := midi.Port()
	//kb := keyboard.Port()
	board := new(midi.Board)

	// I have 9 buttons
	button_pins := [9]machine.Pin{
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
		button, _ := midi.NewMidiControlButton(button_pins[i], 1, uint8(i), 1)
		button.Init()
		board.AddButton(button)
	}
	// removing yhis coz I think I am smart enough to not do this wrong
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}

	boardLED.High()
	ticker := time.NewTicker(time.Millisecond * 10)
	for range ticker.C {
		board.OnTick()
	}
}

// TODO: this will eventually be where the actual loop is
func run() error {

	return nil
}

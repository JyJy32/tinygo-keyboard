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
	var button *midi.MidiControlButton = midi.NewMidiControlButton(machine.GP7, 1, 0, 1).Init()

	boardLED.High()
	ticker := time.NewTicker(time.Millisecond * 20)
	for range ticker.C {
		button.OnTick()
	}
}

// TODO: this will eventually be where the actual loop is
func run() error {

	return nil
}

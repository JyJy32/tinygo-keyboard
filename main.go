package main

import (
	"machine"
	"machine/usb"
	"time"
)

func main() {
	boardLED := machine.LED
	boardLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	usb.Product = "jyx-controller"
	// port := midi.Port()
	//kb := keyboard.Port()
	var button *midiButton = NewMidiButton(machine.GP7, 1)

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

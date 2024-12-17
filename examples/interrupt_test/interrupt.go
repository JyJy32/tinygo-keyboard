package main

import (
	"machine"
	"time"
)

func main() {
	pin := machine.GP7
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})

	pin.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
		currentState := p.Get()

		if currentState {
			print("rising")
		} else {
			print("falling")
		}
	})
	for {
		time.Sleep(time.Millisecond * 100)
	}
}

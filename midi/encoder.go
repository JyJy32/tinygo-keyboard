// used github.com/bgould/tinygo-rotary-encoder as inspiration
package midi

import "machine"

func New(pinA, pinB machine.Pin) *Encoder {
	return &Encoder{pinA: pinA, pinB: pinB}
}

type Encoder struct {
	pinA machine.Pin
	pinB machine.Pin

	hAB   uint8
	value uint8
}

func (e *Encoder) Init() *Encoder {
	e.pinA.Configure(machine.PinConfig{Mode: machine.PinInput})
	e.pinB.Configure(machine.PinConfig{Mode: machine.PinInput})
	e.pinA.SetInterrupt(machine.PinToggle, e.interrupt)
	e.pinB.SetInterrupt(machine.PinToggle, e.interrupt)
	return e
}

func (e *Encoder) interrupt(pin machine.Pin) {
	// CW   if 00 to 10 | 10 to 11 | 11 to 01 | 01 to 00
	// CCW  if 00 to 01 | 01 to 11 | 11 to 10 | 10 to 00
	aHigh, bHigh := e.pinA.Get(), e.pinB.Get()
	//move the AB value over by 2 to insert to current value and have to old one to compare to it
	e.hAB <<= 2
	if aHigh {
		e.hAB |= 1 << 1
	}
	if bHigh {
		e.hAB |= 1
	}
	// only need the first 4
	switch e.hAB & 0x0F {
	case 2, 11, 13, 4:
		e.cw()
	case 1, 7, 14, 8:
		e.ccw()
	}
}

func (e *Encoder) cw() {
	println("CLOCK WISE")
}
func (e *Encoder) ccw() {
	println("COUNTER CLOCK WISE")
}

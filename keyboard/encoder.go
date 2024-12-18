// used github.com/bgould/tinygo-rotary-encoder as inspiration
package keyboard

import "machine"

func NewEncoder(pinA, pinB machine.Pin) *Encoder {
	return &Encoder{pinA: pinA, pinB: pinB}
}

type Encoder struct {
	pinA machine.Pin
	pinB machine.Pin

	hAB   uint8
	value uint8
	cw    bool
	ccw   bool
	onCW  Callback
	onCCW Callback
}

func (e *Encoder) Init() *Encoder {
	e.pinA.Configure(machine.PinConfig{Mode: machine.PinInput})
	e.pinB.Configure(machine.PinConfig{Mode: machine.PinInput})
	e.pinA.SetInterrupt(machine.PinToggle, e.interrupt)
	e.pinB.SetInterrupt(machine.PinToggle, e.interrupt)
	return e
}

func (e *Encoder) OnTick() error {
	if e.cw {
		e.onClockWise()
		e.cw = false
	}
	if e.ccw {
		e.onCounterClockWise()
		e.ccw = false
	}
	return nil
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
	// only check the first else it will fire 4 times per click of the encoder
	case 2: //, 11, 13, 4:
		e.cw = true
	case 1: //, 7, 14, 8:
		e.ccw = true
	}
}

func (e *Encoder) SetonCW(fn Callback) *Encoder {
	e.onCW = fn
	return e
}

func (e *Encoder) SetonCCW(fn Callback) *Encoder {
	e.onCCW = fn
	return e
}

func (e *Encoder) onClockWise() {
	if e.onCW != nil {
		e.onCW(e.pinA, 0)
	}
}

func (e *Encoder) onCounterClockWise() {
	if e.onCCW != nil {
		e.onCCW(e.pinB, 0)
	}
}

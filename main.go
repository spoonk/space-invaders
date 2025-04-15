package main

import (
	"time"
)

const NANOSECOND = 1000000
const FRAME_DURATION = 1000 / 60

func main() {
	r := Renderer{}
	r.init()
	state := NewGameState()
	for {
		state = state.advance()
		var ui = state.getUI()
		time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
		r.draw(ui)
	}

}

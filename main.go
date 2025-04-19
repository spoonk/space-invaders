package main

import (
	// "fmt"
	"time"
)

const NANOSECOND = 1000000

// const FRAME_DURATION = 1000 / 60
const FRAME_DURATION = 1000 / 10

var kp = rune(NO_INPUT)

func main() {
	testKeyboard()
	// gameLoop()
}

func gameLoop() {
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

func testKeyboard() {
	handler := NewKeyboardInput()
	handler.init()

	controller := NewKeyBoardInputController()
	controller.init(handler)

	go handler.loop()

	for {
		// just wait
	}
}

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
	// testKeyboard()
	gameLoop()
}

func gameLoop() {
	r := Renderer{}
	r.init()
	state := NewGameState()
	keyboardInput := newKeyBoardInputController()
	keyboardInput.init()

	go keyboardInput.refreshEternally()
	// go asyncReadKeyboard()

	for {
		state = state.advance()
		var ui = state.getUI()
		time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
		r.draw(ui)
	}
}

func testKeyboard() {
	//
	// keyboardInput := newKeyBoardInputController()
	// keyboardInput.init()
	// go keyboardInput.refreshEternally()
	// // go asyncReadKeyboard()
	// for {
	// 	clearScreen()
	// 	fmt.Println(kp)
	// 	time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
	// }
}

// dummy fn to consume last key pressed
// func asyncReadKeyboard() {
// 	for {
// 		// kp = key
// 	}
// }

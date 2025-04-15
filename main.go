package main

import (
	"fmt"
	"time"
)

const NANOSECOND = 1000000

const FRAME_DURATION = 1000 / 60

var keyPress = make(chan rune) // ermmm

func main() {
	r := Renderer{}
	r.init()
	state := NewGameState()
	keyboardInput := newKeyBoardInputController()
	keyboardInput.init()

	go keyboardInput.refreshEternally()
	go asyncReadKeyboard()

	for {
		state = state.advance()
		var ui = state.getUI()
		time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
		r.draw(ui)
	}

}

// dummy fn to consume last key pressed
func asyncReadKeyboard() {
	for {
		key := <-keyPress
		fmt.Printf("%c\n", key)
	}
}

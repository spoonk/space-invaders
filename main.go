package main

import (
	// "fmt"
	"fmt"
	"time"
)

const NANOSECOND = 1000000

var controller *KeyboardInputController

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

	handler := NewKeyboardInput()
	handler.init()

	controller = NewKeyBoardInputController()
	controller.init(handler)

	state := NewGameState()

	go handler.loop()

	for {
		state = state.advance()
		var ui = state.getUI()
		r.draw(ui)
		time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
	}
}

func testKeyboard() {
	handler := NewKeyboardInput()
	handler.init()

	controller := NewKeyBoardInputController()
	controller.init(handler)

	r := Renderer{}
	r.init()

	go handler.loop()

	for {

		r.draw([]AbstractUiComponent{})
		fmt.Printf("Last pressed: %c\n", controller.getLastKeypress())
		fmt.Printf("Currently pressed pressed: %c\n", controller.getCurrentKeypress())
		time.Sleep(FRAME_DURATION * NANOSECOND)
	}
}

func GetController() *KeyboardInputController {
	if controller == nil {
		// exit()
		fmt.Println("controller not instantiated")
	}
	return controller
}

package main

import (
	// "fmt"
	"fmt"
	"time"
)

const NANOSECOND = 1000000

var controller *KeyboardInputController

// const FRAME_DURATION = 1000 / 60
const FRAME_DURATION = 1000 / 60

var kp = rune(NO_INPUT)

var Run = true

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
	debugPane := NewDebugPane(state)

	go handler.loop()

	defer cleanup(&r, handler)

	for Run {
		state.advance()
		ui := state.getUI()
		ui = append(ui, debugPane.getUI()...)
		r.draw(ui)
		time.Sleep(time.Duration(FRAME_DURATION * NANOSECOND))
	}
}

func cleanup(r *Renderer, k *KeyboardInputHandler) {
	r.cleanup()
	k.cleanup()
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
		fmt.Println("controller not instantiated")
	}
	return controller
}

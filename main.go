package main

import (
	"fmt"
	"space-invaders/constants"
	"time"
)

var controller *KeyboardInputController

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

	program := NewProgramStateMaanger()
	program.init()

	go handler.loop()

	defer cleanup(&r, handler)

	for Run {
		program.update()
		ui := program.getUI()

		r.draw(ui)

		time.Sleep(time.Duration(constants.FRAME_DURATION * constants.NANOSECOND))
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
		time.Sleep(constants.FRAME_DURATION * constants.NANOSECOND)
	}
}

func GetController() *KeyboardInputController {
	if controller == nil {
		fmt.Println("controller not instantiated")
	}
	return controller
}

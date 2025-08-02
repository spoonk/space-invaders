package main

import (
	"fmt"
	"space-invaders/constants"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"time"
)

var controller *keyboard.KeyboardInputController

var Run = true

func main() {
	// testKeyboard()
	gameLoop()
}

func gameLoop() {
	r := Renderer{}
	r.init()

	handler := keyboard.NewKeyboardInput()
	handler.Init()
	handler.RegisterCallback('q', func(_ rune) { handler.Cleanup(); Run = false })

	controller = keyboard.NewKeyBoardInputController()
	controller.Init(handler)

	program := NewProgramStateMaanger()
	program.init()

	go handler.Loop()

	defer cleanup(&r, handler)

	for Run {
		program.update()
		ui := program.getUI()

		r.draw(ui)

		time.Sleep(time.Duration(constants.FRAME_DURATION * constants.NANOSECOND))
	}
}

func cleanup(r *Renderer, k *keyboard.KeyboardInputHandler) {
	r.cleanup()
	k.Cleanup()
}

func testKeyboard() {
	handler := keyboard.NewKeyboardInput()
	handler.Init()

	handler.RegisterCallback('q', func(_ rune) { handler.Cleanup(); Run = false })

	controller := keyboard.NewKeyBoardInputController()
	controller.Init(handler)

	r := Renderer{}
	r.init()

	go handler.Loop()

	for {

		r.draw([]ui.AbstractUiComponent{})
		fmt.Printf("Last pressed: %c\n", controller.GetLastKeypress())
		fmt.Printf("Currently pressed pressed: %c\n", controller.GetCurrentKeypress())
		time.Sleep(constants.FRAME_DURATION * constants.NANOSECOND)
	}
}

func GetController() *keyboard.KeyboardInputController {
	if controller == nil {
		fmt.Println("controller not instantiated")
	}
	return controller
}

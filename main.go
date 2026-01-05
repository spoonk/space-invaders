package main

import (
	"fmt"
	"space-invaders/constants"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"space-invaders/utils"
	"time"
)

var Run = true

func main() {
	// runtime.Breakpoint()
	// testKeyboard()
	gameLoop()
}

func gameLoop() {
	r := Renderer{center: utils.Point{X: 65, Y: 30}}
	r.init()

	handler := keyboard.NewKeyboardInput()
	handler.Init()
	handler.RegisterCallback('q', func(_ rune) { handler.Cleanup(); Run = false })

	controller := keyboard.GetController()
	controller.Init(handler)

	program := NewProgramStateMaanger()
	program.init()

	go handler.Loop()

	defer cleanup(&r, handler)

	for Run {
		program.update()
		staticUI := program.GetStaticUI()
		// dynamicUI := program.GetDynamicUI()

		// map dynamic -> static UI

		r.draw(staticUI)

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
		r.draw([]ui.StaticUI{})
		fmt.Printf("Last pressed: %c\n", controller.GetLastKeypress())
		fmt.Printf("Currently pressed pressed: %c\n", controller.GetCurrentKeypress())
		time.Sleep(constants.FRAME_DURATION * constants.NANOSECOND)
	}
}

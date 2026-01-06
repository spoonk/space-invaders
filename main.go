package main

import (
	"fmt"
	"space-invaders/constants"
	"space-invaders/keyboard"
	"space-invaders/ui"
	"time"
)

var Run = true

func main() {
	// testKeyboard()
	gameLoop()
}

func gameLoop() {
	r := Renderer{rasterizedCache: make(map[string][]string)}
	resolver := NewImageResolver()
	r.init()

	handler := keyboard.NewKeyboardInput()
	handler.Init()

	// global game interrupt
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
		dynamicUI := program.GetDynamicUI()

		hydratedUI := resolver.GetHydratedUI(dynamicUI)
		scaledUI := r.ScaleHydratedImages(hydratedUI)

		r.draw(append(staticUI, scaledUI...))

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

	r := Renderer{rasterizedCache: make(map[string][]string)}
	r.init()

	go handler.Loop()

	for {
		r.draw([]ui.StaticUI{})
		fmt.Printf("Last pressed: %c\n", controller.GetLastKeypress())
		fmt.Printf("Currently pressed pressed: %c\n", controller.GetCurrentKeypress())
		time.Sleep(constants.FRAME_DURATION * constants.NANOSECOND)
	}
}

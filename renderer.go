package main

import (
	"fmt"
	"space-invaders/ui"
	"space-invaders/utils"
)

const (
	CLEAR_ANSI  = "\033[2J"
	HOME_ANSI   = "\033[H"
	SHOW_CURSOR = "\x1b[?25h"
	HIDE_CURSOR = "\x1b[?25l"
)

type Renderer struct{}

func (r *Renderer) draw(components []ui.AbstractUiComponent) {
	clearScreen()
	for i := 0; i < len(components); i++ {
		elm := components[i]
		text := elm.GetRasterized()
		drawAtPosition(text, elm.GetTopLeft())
	}
}

func (r Renderer) init() {
	clearScreen()
	hideCursor()
}

// stuff to clear screen

func clearScreen() {
	fmt.Print(CLEAR_ANSI)
	fmt.Print(HOME_ANSI)
}

func hideCursor() {
	fmt.Print(HIDE_CURSOR)
}

func (r *Renderer) cleanup() {
	fmt.Print(SHOW_CURSOR)
}

func drawAtPosition(sprite string, p utils.Point) {
	moveCursorTo(p)
	fmt.Print(sprite)
}

func moveCursorTo(p utils.Point) {
	fmt.Printf("\033[%d;%dH", p.Y, p.X)
}

package main

import "fmt"

const CLEAR_ANSI = "\033[2J"
const HOME_ANSI = "\033[H"
const SHOW_CURSOR = "\x1b[?25h"
const HIDE_CURSOR = "\x1b[?25l"

type Renderer struct{}

func (r Renderer) draw(ui []AbstractUiComponent) {
	clearScreen()
	for i := 0; i < len(ui); i++ {
		elm := ui[i]
		text := elm.getRasterized()
		drawAtPosition(text, elm.getTopLeft())
	}
}

func (r Renderer) init() {
	clearScreen()
	hideCursor()
}

// stuff to clear screen

func clearScreen() {
	fmt.Printf(CLEAR_ANSI)
	fmt.Printf(HOME_ANSI)
}

func hideCursor() {
	fmt.Printf(HIDE_CURSOR)
}

func drawAtPosition(sprite string, p Point) {
	moveCursorTo(p)
	fmt.Print(sprite)

}

func moveCursorTo(p Point) {
	fmt.Printf("\033[%d;%dH", p.y, p.x)
}

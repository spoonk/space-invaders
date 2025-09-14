package main

import (
	"fmt"
	"golang.org/x/term"
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"
)

const (
	CLEAR_ANSI  = "\033[2J"
	HOME_ANSI   = "\033[H"
	SHOW_CURSOR = "\x1b[?25h"
	HIDE_CURSOR = "\x1b[?25l"
)

type Renderer struct {
	center utils.Point
}

func (r *Renderer) draw(components []ui.AbstractUiComponent) {
	clearScreen()
	for i := 0; i < len(components); i++ {
		elm := components[i]
		text := elm.GetRasterized()
		r.drawAtPosition(text, elm.GetTopLeft())
	}
}

func (r *Renderer) init() {
	clearScreen()
	hideCursor()

}

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

func (r *Renderer) drawAtPosition(sprite string, p utils.Point) {
	normalizedPoint := r.normalizePosition(p)
	moveCursorTo(normalizedPoint)
	// todo: check if a sprite would go off screeen & don't render otherwise - actually, should never be off screen
	fmt.Print(sprite)
}

func (r *Renderer) normalizePosition(p utils.Point) utils.Point {
	width, height, err := term.GetSize(0)
	if err != nil {
		width = 100  // default values
		height = 100 // default values
	}
	// TODO: introduce scaling from 'game space' to 'screen resolution'
	normalizedX := p.X - constants.GAME_BOUNDARY.W/2
	normalizedY := p.Y - constants.GAME_BOUNDARY.H/2
	return utils.Point{X: normalizedX + width/2, Y: normalizedY + height/2}
	// return utils.Point{X: normalizedX + r.center.X, Y: normalizedY + r.center.Y}
}

func moveCursorTo(p utils.Point) {
	fmt.Printf("\033[%d;%dH", p.Y, p.X)
}

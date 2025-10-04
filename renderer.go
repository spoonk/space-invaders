package main

import (
	"fmt"
	"math"
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"

	"golang.org/x/term"
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

// centers a point around the middle of the screen
func (r *Renderer) normalizePosition(p utils.Point) utils.Point {
	screenWidth, screenHeight, err := term.GetSize(0)
	if err != nil {
		screenWidth = constants.GAME_BOUNDARY.W
		screenHeight = constants.GAME_BOUNDARY.H
	}
	normalizedX := p.X - constants.GAME_BOUNDARY.W/2
	normalizedY := p.Y - constants.GAME_BOUNDARY.H/2

	screenSpacePoint := r.gameSpaceToScreenSpace(utils.Point{X: normalizedX, Y: normalizedY})

	return utils.Point{X: screenSpacePoint.X + screenWidth/2, Y: screenSpacePoint.Y + screenHeight/2}
}

func (r *Renderer) gameSpaceToScreenSpace(gamePoint utils.Point) utils.Point {
	screenWidth, screenHeight, err := term.GetSize(0)
	if err != nil {
		screenWidth = constants.GAME_BOUNDARY.W
		screenHeight = constants.GAME_BOUNDARY.H
	}

	xScalar := float64(screenWidth) / float64(constants.GAME_BOUNDARY.W)
	yScalar := float64(screenHeight) / float64(constants.GAME_BOUNDARY.H)

	return utils.Point{X: int(math.Round(xScalar * float64(gamePoint.X))), Y: int(math.Round(yScalar * float64(gamePoint.Y)))}
}

func moveCursorTo(p utils.Point) {
	fmt.Printf("\033[%d;%dH", p.Y, p.X)
}

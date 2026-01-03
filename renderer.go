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
		// TODO: test out rasterizing here
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

func (r *Renderer) getScreenSize() (int, int) {

	screenWidth, screenHeight, err := term.GetSize(0)
	if err != nil {
		screenWidth = constants.GAME_BOUNDARY.W
		screenHeight = constants.GAME_BOUNDARY.H
	}

	return screenHeight, screenWidth
}

func (r *Renderer) gameSpaceToScreenSpace(gamePoint utils.Point) utils.Point {
	// TODO: this will eventually need to be a bit more complicated. Need to leave some buffer room outside of the game. Can't have
	screenWidth, screenHeight := r.getScreenSize()
	xScalar := float64(screenWidth) / float64(constants.GAME_BOUNDARY.W)
	yScalar := float64(screenHeight) / float64(constants.GAME_BOUNDARY.H)

	return utils.Point{X: int(math.Round(xScalar * float64(gamePoint.X))), Y: int(math.Round(yScalar * float64(gamePoint.Y)))}
}

func moveCursorTo(p utils.Point) {
	fmt.Printf("\033[%d;%dH", p.Y, p.X)
}

// TODO: just make rasterization udpates here

func (r *Renderer) scaleSprite(image [][]float64, gameSpaceContainer *utils.Box) []string {
	screenWidthPx, screenHeightPx := r.getScreenSize()

	// proportion of game space taken up by bounding box
	relativeHeight := float64(gameSpaceContainer.H) / float64(constants.GAME_BOUNDARY.H)
	relativeWidth := float64(gameSpaceContainer.W) / float64(constants.GAME_BOUNDARY.W)

	// scale the image to these values
	finalHeightPx := int32(relativeHeight * float64(screenHeightPx))
	finalWidthPx := int32(relativeWidth * float64(screenWidthPx))

	finalImage := scaleImageToResolution(&image, int(finalHeightPx), int(finalWidthPx))

	return mapImageToAscii(&finalImage)

}

func scaleImageToResolution(image *[][]float64, height int, width int) [][]int {
	src := *image
	srcH := len(src)
	if srcH == 0 {
		return nil
	}
	srcW := len(src[0])

	finalImage := make([][]int, height)
	for i := range finalImage {
		finalImage[i] = make([]int, width)
	}

	// Calculate scale factors
	rowScale := float64(srcH-1) / float64(height-1)
	colScale := float64(srcW-1) / float64(width-1)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			// Map destination coordinates to source coordinates
			srcR := float64(r) * rowScale
			srcC := float64(c) * colScale

			// Find the indices of the 4 surrounding pixels
			rLow := int(srcR)
			cLow := int(srcC)
			rHigh := rLow + 1
			cHigh := cLow + 1

			// Clamp high indices to stay within bounds
			if rHigh >= srcH {
				rHigh = srcH - 1
			}
			if cHigh >= srcW {
				cHigh = srcW - 1
			}

			// Calculate the weights (fractions)
			rWeight := srcR - float64(rLow)
			cWeight := srcC - float64(cLow)

			// Interpolate between the four pixels
			// Top-left (rLow, cLow), Top-right (rLow, cHigh)
			// Bottom-left (rHigh, cLow), Bottom-right (rHigh, cHigh)
			val := src[rLow][cLow]*(1-rWeight)*(1-cWeight) +
				src[rLow][cHigh]*(1-rWeight)*cWeight +
				src[rHigh][cLow]*rWeight*(1-cWeight) +
				src[rHigh][cHigh]*rWeight*cWeight

			// Cast the result to int for the final image
			finalImage[r][c] = int(val)
		}
	}

	return finalImage
}

func mapImageToAscii(image *[][]int) []string {
	chars := []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,^`'. ")

	height := len(*image)
	width := len((*image)[0])

	finalImage := make([]string, height)

	max := getMaxValue(image)
	valToIndex := max / len(chars)

	for r := 0; r < height; r++ {
		row := ""
		for c := 0; c < width; c++ {
			ind := int(valToIndex * ((*image)[r][c]))
			row += string(chars[ind])
		}

		finalImage[r] = row
	}
	return finalImage
}

func getMaxValue(image *[][]int) int {
	max := 0
	for y := range *image {
		for x := range (*image)[0] {
			if (*image)[y][x] > max {
				max = (*image)[y][x]
			}
		}
	}
	return max
}

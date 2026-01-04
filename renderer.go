package main

import (
	"fmt"
	"math"
	"space-invaders/constants"
	"space-invaders/ui"
	"space-invaders/utils"

	"golang.org/x/term"
	"image"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"os"
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
	img, _ := readImageToFloat64("invader.png")
	scaled := r.scaleSprite(img, &utils.Box{X: 0, Y: 0, H: 1, W: 3})
	for i := 0; i < len(components); i++ {
		elm := components[i]
		// TODO: test out rasterizing here
		// text := elm.GetRasterized()
		// r.drawAtPosition([]string{text}, elm.GetTopLeft())
		r.drawAtPosition(scaled, elm.GetTopLeft())
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

func (r *Renderer) drawAtPosition(sprite []string, p utils.Point) {
	normalizedPoint := r.normalizePosition(p)
	for r := range sprite {
		newPoint := utils.Point{X: normalizedPoint.X, Y: normalizedPoint.Y + r}
		if newPoint.X < 0 || newPoint.Y < 0 {
			continue
		}
		moveCursorTo(newPoint)
		fmt.Print(sprite[r])
	}
}

// centers a point around the middle of the screen
func (r *Renderer) normalizePosition(p utils.Point) utils.Point {
	screenWidth, screenHeight := r.getScreenSize()

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

	return screenWidth, screenHeight
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
	asciified := mapImageToAscii(&finalImage)

	return asciified
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
	rowScale := float64(srcH) / float64(height)
	colScale := float64(srcW) / float64(width)

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
	valToIndex := float64(len(chars)) / float64(max)

	for r := 0; r < height; r++ {
		row := ""
		for c := 0; c < width; c++ {
			ind := int(valToIndex * float64(((*image)[r][c])))
			if ind >= len(chars) {
				ind = len(chars) - 1
			}

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

func readImageToFloat64(filePath string) ([][]float64, error) {
	// 1. Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 2. Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// 3. Initialize the 2D slice
	pixels := make([][]float64, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]float64, width)
		for x := 0; x < width; x++ {
			// 4. Get pixel color and convert to RGBA
			r, g, b, _ := img.At(x, y).RGBA()

			// Go's RGBA() returns values in range [0, 65535].
			// We convert them to [0, 255] for standard float representation.
			lum := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)

			pixels[y][x] = lum
		}
	}

	return pixels, nil
}

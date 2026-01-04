package utils

import (
	"image"
	"os"
)

func ReadImageToFloat64(filePath string) ([][]float64, error) {
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

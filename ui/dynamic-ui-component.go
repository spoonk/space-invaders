package ui

import "space-invaders/utils"

type DynamicUI struct {
	Path        string
	BoundingBox utils.Box
}

type HydratedDynamicUI struct {
	BoundingBox utils.Box
	Image       *[][]float64
}

func NewDynamicUI(path string, box utils.Box) DynamicUI {
	return DynamicUI{Path: path, BoundingBox: box}
}

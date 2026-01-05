package main

import (
	"space-invaders/ui"
	"space-invaders/utils"
)

type ImageResolver struct {
	imageCache map[string]*[][]float64
}

func NewImageResolver() *ImageResolver {
	return &ImageResolver{imageCache: make(map[string]*[][]float64)}
}

func (resolver *ImageResolver) GetHydratedUI(components []ui.DynamicUI) []ui.HydratedDynamicUI {
	hydratedUI := []ui.HydratedDynamicUI{}
	for _, component := range components {
		hydratedUI = append(hydratedUI, ui.HydratedDynamicUI{
			BoundingBox: component.BoundingBox,
			Image:       resolver.resolveImage(component.Path),
			Path:        component.Path})
	}

	return hydratedUI
}

func (resolver *ImageResolver) resolveImage(imagePath string) *[][]float64 {
	value, ok := resolver.imageCache[imagePath]
	if ok {
		return value
	}

	image, err := utils.ReadImageToFloat64(imagePath)
	if err != nil {
		panic("failed to read image")
	}

	resolver.imageCache[imagePath] = &image
	return &image
}

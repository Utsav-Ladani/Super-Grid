package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Direction int

const (
	DirectionHorizontal Direction = 1
	DirectionVertical   Direction = 2
)

type SuperGridOptions struct {
	direction Direction
	spacing   float32
}

func NewSuperGrid(superGridOptions SuperGridOptions, superGridElements []*SuperGridElement) fyne.CanvasObject {
	superGrid := &SuperGrid{
		Elements:         superGridElements,
		SuperGridOptions: superGridOptions,
	}

	superGrid.ExtendBaseWidget(superGrid)

	return superGrid
}

type SuperGrid struct {
	widget.BaseWidget
	Elements         []*SuperGridElement
	SuperGridOptions SuperGridOptions
}

func (s *SuperGrid) Size() fyne.Size {
	var width float32 = 0.0
	var height float32 = 0.0

	totalSpace := s.SuperGridOptions.spacing * float32(len(s.Elements)-1)

	if s.SuperGridOptions.direction == DirectionHorizontal {
		width += totalSpace

		for _, element := range s.Elements {
			width += element.Size().Width
			height = float32(math.Max(float64(height), float64(element.Size().Height)))
		}
	} else {
		height += totalSpace

		for _, element := range s.Elements {
			height += element.Size().Height
			width = float32(math.Max(float64(width), float64(element.Size().Width)))
		}
	}

	return fyne.NewSize(width, height)
}

func (s *SuperGrid) CreateRenderer() fyne.WidgetRenderer {
	canvasObjs := []fyne.CanvasObject{}

	for _, element := range s.Elements {
		canvasObjs = append(canvasObjs, element)
	}

	return &superGridRenderer{
		superGrid:  s,
		canvasObjs: canvasObjs,
		elements:   s.Elements,
	}
}

type superGridRenderer struct {
	superGrid  *SuperGrid
	canvasObjs []fyne.CanvasObject
	elements   []*SuperGridElement
}

func (s *superGridRenderer) Destroy() {
}

func (s *superGridRenderer) Layout(size fyne.Size) {
	if s.superGrid.SuperGridOptions.direction == DirectionHorizontal {
		s.LayoutHorizontal(size)
	} else {
		s.LayoutVertical(size)
	}
}

func (s *superGridRenderer) LayoutHorizontal(size fyne.Size) {
	width := size.Width
	height := float32(0.0)

	blockElements := 0

	for _, element := range s.elements {
		if !element.IsBlock {
			width -= element.Size().Width
		} else {
			blockElements++
		}
		height = float32(math.Max(float64(height), float64(element.Size().Height)))
	}

	len := len(s.elements)
	spacersWidth := s.superGrid.SuperGridOptions.spacing * float32(len-1)
	width -= spacersWidth

	perElementWidth := width / float32(blockElements)
	perElementHeight := float32(math.Min(float64(height), float64(size.Height)))

	for _, element := range s.elements {
		if element.IsBlock {
			element.Resize(fyne.NewSize(perElementWidth, element.Size().Height))
		}

		if element.Fill {
			element.Resize(fyne.NewSize(element.Size().Width, size.Height))
		}
	}

	posX := float32(0.0)
	posY := float32(0.0)

	spacerWidth := s.superGrid.SuperGridOptions.spacing

	for _, element := range s.elements {
		elePosX := posX
		elePosY := posY

		if element.Alignment == AlignmentCenter {
			elePosY += (perElementHeight - element.Size().Height) / 2
		} else if element.Alignment == AlignmentBottom {
			elePosY += (perElementHeight - element.Size().Height)
		}

		element.Move(fyne.NewPos(elePosX, elePosY))

		posX += element.Size().Width + spacerWidth
	}
}

func (s *superGridRenderer) LayoutVertical(size fyne.Size) {
	width := float32(0.0)
	height := size.Height

	blockElements := 0

	for _, element := range s.elements {
		if !element.IsBlock {
			height -= element.Size().Height
		} else {
			blockElements++
		}
		width = float32(math.Max(float64(width), float64(element.Size().Width)))
	}

	len := len(s.elements)
	spacersHeight := s.superGrid.SuperGridOptions.spacing * float32(len-1)
	height -= spacersHeight

	perElementWidth := float32(math.Min(float64(width), float64(size.Width)))
	perElementHeight := height / float32(blockElements)

	for _, element := range s.elements {
		if element.IsBlock {
			element.Resize(fyne.NewSize(element.Size().Width, perElementHeight))
		}

		if element.Fill {
			element.Resize(fyne.NewSize(size.Width, element.Size().Height))
		}
	}

	posX := float32(0.0)
	posY := float32(0.0)

	spacerHeight := s.superGrid.SuperGridOptions.spacing

	for _, element := range s.elements {
		elePosX := posX
		elePosY := posY

		if element.Alignment == AlignmentCenter {
			elePosX += (perElementWidth - element.Size().Width) / 2
		} else if element.Alignment == AlignmentBottom {
			elePosX += (perElementWidth - element.Size().Width)
		}

		element.Move(fyne.NewPos(elePosX, elePosY))

		posY += element.Size().Height + spacerHeight
	}
}

func (s *superGridRenderer) MinSize() fyne.Size {
	var width float32 = 0.0
	var height float32 = 0.0

	totalSpace := s.superGrid.SuperGridOptions.spacing * float32(len(s.elements)-1)

	if s.superGrid.SuperGridOptions.direction == DirectionHorizontal {
		width += totalSpace

		for _, element := range s.elements {
			if !element.IsBlock {
				width += element.Size().Width
			}

			if !element.Fill {
				height = float32(math.Max(float64(height), float64(element.Size().Height)))
			}

		}
	} else {
		height += totalSpace

		for _, element := range s.elements {
			if !element.IsBlock {
				height += element.Size().Height
			}

			if !element.Fill {
				width = float32(math.Max(float64(width), float64(element.Size().Width)))
			}
		}
	}

	return fyne.NewSize(width, height)
}

func (s *superGridRenderer) Objects() []fyne.CanvasObject {
	return s.canvasObjs
}

func (s *superGridRenderer) Refresh() {
}

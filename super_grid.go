package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Alignment int

const (
	AlignmentTop    Alignment = 1
	AlignmentCenter Alignment = 2
	AlignmentBottom Alignment = 3
)

type SuperGridElement struct {
	Obj       fyne.CanvasObject
	IsBlock   bool
	Alignment Alignment
	Fill      bool
}

type Direction int

const (
	DirectionHorizontal Direction = 1
	DirectionVertical   Direction = 2
)

type SuperGridOptions struct {
	direction Direction
	spacing   float32
}

func NewSuperGrid(superGridElements []*SuperGridElement, superGridOptions SuperGridOptions) fyne.CanvasObject {
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

// func (s *SuperGrid) MinSize() fyne.Size {
// 	return fyne.NewSize(0, 0)
// }

func (s *SuperGrid) CreateRenderer() fyne.WidgetRenderer {
	canvasObjs := []fyne.CanvasObject{}

	for _, element := range s.Elements {
		canvasObjs = append(canvasObjs, element.Obj)
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
			width -= element.Obj.Size().Width
		} else {
			blockElements++
		}
		height = float32(math.Max(float64(height), float64(element.Obj.Size().Height)))
	}

	len := len(s.elements)
	spacersWidth := s.superGrid.SuperGridOptions.spacing * float32(len-1)
	width -= spacersWidth

	perElementWidth := width / float32(blockElements)
	perElementHeight := float32(math.Min(float64(height), float64(size.Height)))

	for _, element := range s.elements {
		if element.IsBlock {
			element.Obj.Resize(fyne.NewSize(perElementWidth, element.Obj.Size().Height))
		}

		if element.Fill {
			element.Obj.Resize(fyne.NewSize(element.Obj.Size().Width, size.Height))
		}
	}

	posX := float32(0.0)
	posY := float32(0.0)

	spacerWidth := s.superGrid.SuperGridOptions.spacing

	for _, element := range s.elements {
		elePosX := posX
		elePosY := posY

		if element.Alignment == AlignmentCenter {
			elePosY += (perElementHeight - element.Obj.Size().Height) / 2
		} else if element.Alignment == AlignmentBottom {
			elePosY += (perElementHeight - element.Obj.Size().Height)
		}

		element.Obj.Move(fyne.NewPos(elePosX, elePosY))

		posX += element.Obj.Size().Width + spacerWidth
	}
}

func (s *superGridRenderer) LayoutVertical(size fyne.Size) {
	width := float32(0.0)
	height := size.Height

	blockElements := 0

	for _, element := range s.elements {
		if !element.IsBlock {
			height -= element.Obj.Size().Height
		} else {
			blockElements++
		}
		width = float32(math.Max(float64(width), float64(element.Obj.Size().Width)))
	}

	len := len(s.elements)
	spacersHeight := s.superGrid.SuperGridOptions.spacing * float32(len-1)
	height -= spacersHeight

	perElementWidth := float32(math.Min(float64(width), float64(size.Width)))
	perElementHeight := height / float32(blockElements)

	for _, element := range s.elements {
		if element.IsBlock {
			element.Obj.Resize(fyne.NewSize(element.Obj.Size().Width, perElementHeight))
		}

		if element.Fill {
			element.Obj.Resize(fyne.NewSize(size.Width, element.Obj.Size().Height))
		}
	}

	posX := float32(0.0)
	posY := float32(0.0)

	spacerHeight := s.superGrid.SuperGridOptions.spacing

	for _, element := range s.elements {
		elePosX := posX
		elePosY := posY

		if element.Alignment == AlignmentCenter {
			elePosX += (perElementWidth - element.Obj.Size().Width) / 2
		} else if element.Alignment == AlignmentBottom {
			elePosX += (perElementWidth - element.Obj.Size().Width)
		}

		element.Obj.Move(fyne.NewPos(elePosX, elePosY))

		posY += element.Obj.Size().Height + spacerHeight
	}
}

func (s *superGridRenderer) MinSize() fyne.Size {
	var width float32 = 0.0
	var height float32 = 0.0

	spacersWidth := s.superGrid.SuperGridOptions.spacing * float32(len(s.elements)-1)
	width += spacersWidth

	for _, element := range s.elements {
		if !element.IsBlock {
			width += element.Obj.Size().Width
		}

		height = float32(math.Max(float64(height), float64(element.Obj.Size().Height)))
	}

	return fyne.NewSize(width, height)
}

func (s *superGridRenderer) Objects() []fyne.CanvasObject {
	return s.canvasObjs
}

func (s *superGridRenderer) Refresh() {
}

func (s *superGridRenderer) Resize(size fyne.Size) {
	s.Layout(size)
}

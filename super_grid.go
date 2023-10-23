package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
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
}

func NewSuperGrid(superGridElements []*SuperGridElement) fyne.CanvasObject {
	superGrid := &SuperGrid{
		Elements: superGridElements,
	}

	superGrid.ExtendBaseWidget(superGrid)

	return superGrid
}

type SuperGrid struct {
	widget.BaseWidget
	Elements []*SuperGridElement
}

// func (s *SuperGrid) MinSize() fyne.Size {
// 	return fyne.NewSize(0, 0)
// }

func (s *SuperGrid) CreateRenderer() fyne.WidgetRenderer {
	canvasObjs := []fyne.CanvasObject{}
	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(20, 10))

	for _, element := range s.Elements {
		canvasObjs = append(canvasObjs, element.Obj)
	}

	return &superGridRenderer{
		canvasObjs: canvasObjs,
		elements:   s.Elements,
		spacer:     spacer,
	}
}

type superGridRenderer struct {
	canvasObjs []fyne.CanvasObject
	elements   []*SuperGridElement
	spacer     fyne.CanvasObject
}

func (s *superGridRenderer) Destroy() {
}

func (s *superGridRenderer) Layout(size fyne.Size) {
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
	spacersWidth := s.spacer.Size().Width * float32(len-1)
	width -= spacersWidth

	perElementWidth := width / float32(blockElements)
	perElementHeight := float32(math.Min(float64(height), float64(size.Height)))

	for _, element := range s.elements {
		if element.IsBlock {
			element.Obj.Resize(fyne.NewSize(perElementWidth, element.Obj.Size().Height))
		}
	}

	posX := float32(0.0)
	posY := float32(0.0)

	spacerWidth := s.spacer.Size().Width

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

func (s *superGridRenderer) MinSize() fyne.Size {
	var width float32 = 0.0
	var height float32 = 0.0

	spacersWidth := s.spacer.Size().Width * float32(len(s.elements)-1)
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

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Alignment int

const (
	AlignmentStart  Alignment = 1
	AlignmentCenter Alignment = 2
	AlignmentEnd    Alignment = 3
)

type SuperGridElement struct {
	widget.BaseWidget
	Obj       fyne.CanvasObject
	IsBlock   bool
	Alignment Alignment
	Fill      bool
	Margin    [4]float32
}

func NewSuperGridElement(obj fyne.CanvasObject) *SuperGridElement {
	superGridElement := &SuperGridElement{
		Obj: obj,
	}

	superGridElement.ExtendBaseWidget(superGridElement)

	return superGridElement
}

func (s *SuperGridElement) MinSize() fyne.Size {
	size := s.Obj.MinSize()

	size = size.Add(fyne.NewSize(s.Margin[1]+s.Margin[3], s.Margin[0]+s.Margin[2]))

	return size
}

func (s *SuperGridElement) Resize(size fyne.Size) {
	size = size.Subtract(fyne.NewSize(s.Margin[1]+s.Margin[3], s.Margin[0]+s.Margin[2]))

	s.Obj.Resize(size)
}

func (s *SuperGridElement) Size() fyne.Size {
	size := s.Obj.Size()

	size = size.Add(fyne.NewSize(s.Margin[1]+s.Margin[3], s.Margin[0]+s.Margin[2]))

	return size
}

func (s *SuperGridElement) CreateRenderer() fyne.WidgetRenderer {
	return &superGridElementRenderer{
		superGridElement: s,
	}
}

type superGridElementRenderer struct {
	superGridElement *SuperGridElement
}

func (s *superGridElementRenderer) Destroy() {
}

func (s *superGridElementRenderer) Layout(size fyne.Size) {
	size = size.Subtract(fyne.NewSize(s.superGridElement.Margin[1]+s.superGridElement.Margin[3], s.superGridElement.Margin[0]+s.superGridElement.Margin[2]))

	s.superGridElement.Obj.Resize(size)

	pos := fyne.NewPos(s.superGridElement.Margin[3], s.superGridElement.Margin[0])

	s.superGridElement.Obj.Move(pos)
}

func (s *superGridElementRenderer) MinSize() fyne.Size {
	size := s.superGridElement.Obj.MinSize()
	size = size.Add(fyne.NewSize(s.superGridElement.Margin[1]+s.superGridElement.Margin[3], s.superGridElement.Margin[0]+s.superGridElement.Margin[2]))

	return size
}

func (s *superGridElementRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		s.superGridElement.Obj,
	}
}

func (s *superGridElementRenderer) Refresh() {
}

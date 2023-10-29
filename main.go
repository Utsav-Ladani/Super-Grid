package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	sApp := app.New()

	sWindow := sApp.NewWindow("Super Grid")

	var superGrid fyne.CanvasObject

	image := canvas.NewImageFromFile("github.png")
	image.Resize(fyne.NewSize(48, 48))

	label := canvas.NewRectangle(color.RGBA{255, 255, 0, 255})
	label.Resize(fyne.NewSize(50, 32))

	button := widget.NewButton("Click Me", func() {
		if label.FillColor == (color.RGBA{255, 255, 0, 255}) {
			label.FillColor = color.RGBA{0, 255, 0, 255}
		} else {
			label.FillColor = color.RGBA{255, 255, 0, 255}
		}
	})
	button.Resize(fyne.NewSize(100, 40))

	message := canvas.NewRectangle(color.RGBA{0, 0, 255, 255})
	message.Resize(fyne.NewSize(50, 32))

	superGridElements := []*SuperGridElement{
		{
			IsBlock:   false,
			Obj:       image,
			Alignment: AlignmentCenter,
			Fill:      false,
			Margin:    [4]float32{8, 8, 8, 8},
		},
		{
			IsBlock:   true,
			Obj:       label,
			Alignment: AlignmentEnd,
			Fill:      false,
			Margin:    [4]float32{0, 20, 0, 0},
		},
		{
			IsBlock:   false,
			Obj:       button,
			Alignment: AlignmentCenter,
			Fill:      false,
		},
		{
			IsBlock:   true,
			Obj:       message,
			Alignment: AlignmentStart,
			Fill:      true,
			Margin:    [4]float32{10, 10, 10, 10},
		},
	}

	superGridOptions := SuperGridOptions{
		// Direction: DirectionHorizontal,
		Direction: DirectionVertical,
		Spacing:   10.0,
	}

	superGrid = NewSuperGrid(superGridOptions, superGridElements)

	sWindow.SetContent(superGrid)

	sWindow.Resize(fyne.NewSize(480, 400))

	sWindow.ShowAndRun()
}

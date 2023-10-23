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

	image := canvas.NewImageFromFile("github.png")
	image.Resize(fyne.NewSize(48, 48))

	label := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	// label := canvas.NewText("Hello World", color.RGBA{255, 0, 0, 255})
	label.Resize(fyne.NewSize(50, 32))

	button := widget.NewButton("Click Me", func() {})
	button.Resize(fyne.NewSize(100, 32))

	message := canvas.NewRectangle(color.RGBA{0, 0, 255, 255})
	// message := canvas.NewText("This is cool message", color.RGBA{0, 0, 255, 255})
	message.Resize(fyne.NewSize(50, 32))

	superGridElements := []*SuperGridElement{
		{
			IsBlock: false,
			Obj:     image,
		},
		{
			IsBlock: true,
			Obj:     label,
		},
		{
			IsBlock: false,
			Obj:     button,
		},
		{
			IsBlock: true,
			Obj:     message,
		},
	}

	superGrid := NewSuperGrid(superGridElements)

	sWindow.SetContent(superGrid)

	sWindow.Resize(fyne.NewSize(480, 400))

	sWindow.ShowAndRun()
}

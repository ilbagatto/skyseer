package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Skyseer")

	label := widget.NewLabel("Welcome to Skyseer – your astrological workspace.")
	w.SetContent(container.NewCenter(label))

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

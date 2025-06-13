package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/ilbagatto/skyseer/internal/ui"
)

func main() {
	a := app.New()
	w := ui.BuildWindow(a)
	w.ShowAndRun()
}

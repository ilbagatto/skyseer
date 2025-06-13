package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// BuildWindow constructs the main application window.
// Layout: global header, sidebar navigation, and right panel with
// metadata, tabs, and vertical actions.
func BuildWindow(app fyne.App) fyne.Window {
	w := app.NewWindow("Skyseer")

	// === Global Header Toolbar ===
	hdr := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			// TODO: New Chart
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			// TODO: Settings
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			// TODO: Help
		}),
	)

	// === Sidebar: saved charts ===
	chartItems := []string{"Ivan · Radix", "Ivan · Transit"}
	chartList := widget.NewList(
		func() int { return len(chartItems) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(chartItems[i])
		},
	)
	sidebar := container.NewVScroll(chartList)

	// === Chart Metadata (placed above tabs) ===
	metadataLabels := container.NewHBox(
		widget.NewLabel("Name: Ivan"),
		widget.NewLabel("Born: 01 Jan 1980 12:34"),
		widget.NewLabel("Place: Moscow"),
	)

	// === Main Tabs: views and techniques ===
	tabs := container.NewAppTabs(
		container.NewTabItem("Graphics", container.NewCenter(widget.NewLabel("Graphics view"))),
		container.NewTabItem("Ephemeris", container.NewCenter(widget.NewLabel("Ephemeris view"))),
		container.NewTabItem("Houses", container.NewCenter(widget.NewLabel("Houses view"))),
		container.NewTabItem("Statistics", container.NewCenter(widget.NewLabel("Statistics view"))),
		container.NewTabItem("Techniques", container.NewCenter(widget.NewLabel("Techniques view"))),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Combine metadata and tabs into main content
	mainContent := container.NewVBox(metadataLabels, tabs)

	// === Vertical actions ===
	actions := container.NewVBox(
		widget.NewButtonWithIcon("", theme.ContentCutIcon(), func() {
			// TODO: Edit Chart Data
		}),
		widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
			// TODO: Save Chart
		}),
	)

	// Split main content and actions
	rightPanel := container.NewHSplit(mainContent, actions)
	rightPanel.SetOffset(0.90)

	// Split sidebar and rightPanel
	mainSplit := container.NewHSplit(sidebar, rightPanel)
	mainSplit.SetOffset(0.25)

	// Assemble final layout: header above mainSplit
	layout := container.NewBorder(hdr, nil, nil, nil, mainSplit)

	w.SetContent(layout)
	w.Resize(fyne.NewSize(1024, 768))
	return w
}

package frontend

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (f *frontend) renderFooter() *fyne.Container {
	return fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(2),
		&widget.Button{
			Text:          "Router Console",
			Icon:          theme.ViewRefreshIcon(),
			Alignment:     widget.ButtonAlignCenter,
			IconPlacement: 0,
			OnTapped:      func() { log.Println("Not yet implemented") },
		},
		&widget.Button{
			Text:          "Settings",
			Icon:          theme.SettingsIcon(),
			Alignment:     widget.ButtonAlignCenter,
			IconPlacement: 0,
			OnTapped:      func() { log.Println("Not yet implemented") },
		})
}

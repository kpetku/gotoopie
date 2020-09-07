package frontend

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (f *frontend) renderHeader() *fyne.Container {
	top := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		renderToolbar(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), renderBackgroundImage(), &widget.Button{
			Text:          "Disconnect",
			Icon:          theme.VisibilityOffIcon(),
			Alignment:     widget.ButtonAlignCenter,
			IconPlacement: 0,
			OnTapped:      func() { log.Println("Not yet implemented") },
		}))
	return top
}

func renderToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() { log.Println("Not yet implemented") }),
	)
}

func renderBackgroundImage() *canvas.Image {
	b := canvas.NewImageFromResource(resourceCurvePng)
	b.FillMode = canvas.ImageFillStretch
	b.SetMinSize(fyne.NewSize(400, 50))
	return b
}

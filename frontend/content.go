package frontend

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (f *frontend) renderContent() *fyne.Container {
	return fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		renderLogoImage(),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), f.vbox),
		f.renderAccordionMenus(),
		layout.NewSpacer())
}

func renderLogoImage() *canvas.Image {
	l := canvas.NewImageFromResource(resourceLogoPng)
	l.FillMode = canvas.ImageFillContain
	l.SetMinSize(fyne.NewSize(172, 172))
	return l
}

func (f *frontend) renderAccordionMenus() *widget.AccordionContainer {
	ac := widget.NewAccordionContainer()
	ac.Append(renderI2PApplications())
	return ac
}

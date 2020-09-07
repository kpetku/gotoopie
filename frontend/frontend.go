package frontend

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/kpetku/gotoopie/backend"
)

type frontend struct {
	client *backend.Client
	app    fyne.App
	window fyne.Window
	vbox   *widget.Box
}

// Start launches the GUI frontend.
func Start() {
	f := startFrontendAndBackend()
	err := f.client.Authenticate()
	if err != nil {
		log.Printf("Error authenticating: %s", err)
	}
	f.startGUI()
}

func startFrontendAndBackend() *frontend {
	f := new(frontend)
	f.client = backend.NewClient(&backend.Options{
		Address:  "localhost",
		Port:     "7650",
		Password: "itoopie",
	})
	return f
}

func (f *frontend) startGUI() {
	f.vbox = widget.NewVBox()
	f.app = app.New()
	f.window = f.app.NewWindow("gotoopie")

	f.renderNetworkStatus()
	f.renderPeers()
	f.renderOther()
	f.renderAccordionMenus()

	header := f.renderHeader()
	content := f.renderContent()
	footer := f.renderFooter()

	layout := layout.NewBorderLayout(header, footer, nil, nil)
	f.window.SetContent(fyne.NewContainerWithLayout(layout, header, footer, content))

	fyne.CurrentApp().Settings().SetTheme(newToopieTheme())
	f.window.SetIcon(resourceLogoPng)
	f.window.Resize(fyne.NewSize(400, 800))
	f.window.SetFixedSize(true)
	f.window.ShowAndRun()
}

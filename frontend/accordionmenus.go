package frontend

import (
	"log"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func renderI2PApplications() *widget.AccordionItem {
	return widget.NewAccordionItem("I2P Applications",
		widget.NewVBox(
			&widget.Button{
				Text:          "I2P Email",
				Icon:          theme.MailComposeIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
			&widget.Button{
				Text:          "Torrent",
				Icon:          theme.DownloadIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
			&widget.Button{
				Text:          "Web Server",
				Icon:          theme.ComputerIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
			&widget.Button{
				Text:          "Address Book",
				Icon:          theme.FileApplicationIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
			&widget.Button{
				Text:          "I2P Network Manager",
				Icon:          theme.MoveDownIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
			&widget.Button{
				Text:          "Plugins",
				Icon:          theme.FileVideoIcon(),
				Alignment:     widget.ButtonAlignLeading,
				IconPlacement: 0,
				OnTapped:      func() { log.Println("Not yet implemented") },
			},
		))
}

package frontend

import (
	"strconv"
	"time"

	"fyne.io/fyne/widget"
)

func (f *frontend) renderOther() error {
	var u string
	version, err := f.client.RouterVersion()
	if err != nil {
		return err
	}
	uptime, err := f.client.RouterUptime()
	if err != nil {
		return err
	}
	hours := uptime / 3600000
	days := (uptime / 86400000) % 7
	weeks := uptime / (86400000 * 7)
	if hours < 48 {
		u = time.Duration(int64(uptime / 3600000)).String()
	} else if hours > 48 {
		u = strconv.Itoa(days) + " days"
	} else if days > 13 {
		u = strconv.Itoa(weeks) + " weeks"
	}
	status, err := f.client.RouterStatus()
	if err != nil {
		return err
	}
	f.vbox.Append(widget.NewGroup("Other",
		widget.NewLabel("Version: "+version),
		widget.NewLabel("Uptime: "+u),
		widget.NewLabel("Status: "+status),
	))
	return nil
}

package frontend

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func (f *frontend) renderNetworkStatus() error {
	ns, err := f.client.NetStatus()
	if err != nil {
		return err
	}
	var out string
	switch ns {
	case 0:
		out = "Network OK"
	case 1:
		out = "Testing Network"
	case 2:
		out = "Firewalled"
	case 3:
		out = "Hidden"
	case 4:
		out = "Warn: Firewalled and fast"
	case 5:
		out = "Warn: Firewalled and floodfill"
	case 6:
		out = "Warn: Firewalled with inbound TCP"
	case 7:
		out = "Warn: Firewalled with UDP disabled"
	case 8:
		out = "Error: I2CP"
	case 9:
		out = " Error: Clock Skew"
	case 10:
		out = "Error: Private TCP address"
	case 11:
		out = "Error: Symmetric NAT"
	case 12:
		out = "Error: UDP port in use"
	case 13:
		out = "Error: No active peers check connection and firewall"
	case 14:
		out = "Error: UDP disabled and TCP unset"
	}
	f.vbox.Append(widget.NewLabelWithStyle(out, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
	return nil
}

package frontend

import (
	"strconv"

	"fyne.io/fyne/widget"
)

func (f *frontend) renderPeers() error {
	knownPeers, err := f.client.NetDBKnownPeers()
	if err != nil {
		return err
	}
	activePeers, err := f.client.NetDBActivePeers()
	if err != nil {
		return err
	}
	fastPeers, err := f.client.NetDBFastPeers()
	if err != nil {
		return err
	}
	participating, err := f.client.NetTunnelsParticipating()
	if err != nil {
		return err
	}
	f.vbox.Append(
		widget.NewGroup("Peers",
			widget.NewLabel("Active: "+strconv.Itoa(activePeers)+"/"+strconv.Itoa(knownPeers)),
			widget.NewLabel("Fast: "+strconv.Itoa(fastPeers)+"/"+strconv.Itoa(knownPeers)),
			widget.NewLabel("Active: "+strconv.Itoa(knownPeers)),
		))
	f.vbox.Append(widget.NewGroup("Tunnels",
		widget.NewLabel("Participating: "+strconv.Itoa(participating)),
	))
	return nil
}

package backend

import (
	"fmt"
)

// RouterInfo fetches basic information about the I2P router. Uptime, version etc.
type RouterInfo struct {
	Status  string
	Uptime  int
	Version string

	NetBwInbound1s          int
	NetBwInbound15s         int
	NetBwOutbound1s         int
	NetBwOutbound15s        int
	NetStatus               int
	NetTunnelsParticipating int

	NetDBActivePeers       int
	NetDBFastPeers         int
	NetDBHighCapacityPeers int
	NetDBIsReseeding       bool
	NetDBKnownPeers        int
}

// RouterInfo fetches basic information about the I2P router. Uptime, version etc.
func (c *Client) RouterInfo(ri RouterInfo) (RouterInfo, error) {
	params := make(map[string]interface{})
	if ri.Status != "" {
		params["i2p.router.status"] = nil
	}
	if ri.Uptime != 0 {
		params["i2p.router.uptime"] = nil
	}
	if ri.Version != "" {
		params["i2p.router.version"] = nil
	}

	if ri.NetBwInbound1s != 0 {
		params["i2p.router.net.bw.inbound.1s"] = nil
	}
	if ri.NetBwInbound15s != 0 {
		params["i2p.router.net.bw.inbound.15s"] = nil
	}
	if ri.NetBwOutbound1s != 0 {
		params["i2p.router.net.bw.outbound.1s"] = nil
	}
	if ri.NetBwOutbound15s != 0 {
		params["i2p.router.net.bw.outbound.15s"] = nil
	}
	if ri.NetStatus != 0 {
		params["i2p.router.net.status"] = nil
	}
	if ri.NetTunnelsParticipating != 0 {
		params["i2p.router.net.tunnels.participating"] = nil
	}

	if ri.NetDBActivePeers != 0 {
		params["i2p.router.netdb.activepeers"] = nil
	}
	if ri.NetDBFastPeers != 0 {
		params["i2p.router.netdb.fastpeers"] = nil
	}
	if ri.NetDBHighCapacityPeers != 0 {
		params["i2p.router.netdb.highcapacitypeers"] = nil
	}
	if ri.NetDBIsReseeding != false {
		params["i2p.router.netdb.isreseeding"] = nil
	}
	if ri.NetDBKnownPeers != 0 {
		params["i2p.router.netdb.knownpeers"] = nil
	}

	reply, err := c.Send("RouterInfo", params)
	result := RouterInfo{}

	if err != nil || reply.Result == nil {
		return result, fmt.Errorf("RouterInfo failed: %w", err)
	}

	if reply.Result["i2p.router.status"] != nil {
		result.Status = reply.Result["i2p.router.status"].(string)
	}
	if reply.Result["i2p.router.uptime"] != nil {
		result.Uptime = int(reply.Result["i2p.router.uptime"].(float64))
	}
	if reply.Result["i2p.router.version"] != nil {
		result.Version = reply.Result["i2p.router.version"].(string)
	}
	if reply.Result["i2p.router.net.bw.inbound.1s"] != nil {
		result.NetBwInbound1s = int(reply.Result["i2p.router.net.bw.inbound.1s"].(float64))
	}
	if reply.Result["i2p.router.net.bw.inbound.15s"] != nil {
		result.NetBwInbound15s = int(reply.Result["i2p.router.net.bw.inbound.15s"].(float64))
	}
	if reply.Result["i2p.router.net.bw.outbound.1s"] != nil {
		result.NetBwOutbound1s = int(reply.Result["i2p.router.net.bw.outbound.1s"].(float64))
	}
	if reply.Result["i2p.router.net.bw.outbound.15s"] != nil {
		result.NetBwOutbound15s = int(reply.Result["i2p.router.net.bw.outbound.15s"].(float64))
	}
	if reply.Result["i2p.router.net.status"] != nil {
		result.NetStatus = int(reply.Result["i2p.router.net.status"].(float64))
	}
	if reply.Result["i2p.router.net.tunnels.participating"] != nil {
		result.NetTunnelsParticipating = int(reply.Result["i2p.router.net.tunnels.participating"].(float64))
	}
	if reply.Result["i2p.router.netdb.activepeers"] != nil {
		result.NetDBActivePeers = int(reply.Result["i2p.router.netdb.activepeers"].(float64))
	}
	if reply.Result["i2p.router.netdb.fastpeers"] != nil {
		result.NetDBFastPeers = int(reply.Result["i2p.router.netdb.fastpeers"].(float64))
	}
	if reply.Result["i2p.router.netdb.highcapacitypeers"] != nil {
		result.NetDBHighCapacityPeers = int(reply.Result["i2p.router.netdb.highcapacitypeers"].(float64))
	}
	if reply.Result["i2p.router.netdb.isreseeding"] != "" {
		if reply.Result["i2p.router.netdb.isreseeding"] == "true" {
			result.NetDBIsReseeding = true
		}
	}
	if reply.Result["i2p.router.netdb.knownpeers"] != nil {
		result.NetDBKnownPeers = int(reply.Result["i2p.router.netdb.knownpeers"].(float64))
	}
	return result, err
}

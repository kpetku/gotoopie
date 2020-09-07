package backend

import "fmt"

// NetworkSetting fetches or sets various network related settings. Ports, addresses etc.
type NetworkSetting struct {
	NTCPPort      string // What port is used for the TCP transport.
	NTCPHostname  string // What hostname is used for the TCP transport.
	NTCPAutoIP    string // Use automatically detected ip for TCP transport.
	SSUPort       string // What port is used for the UDP transport.
	SSUHostname   string // What hostname is used for the UDP transport.
	SSUAutoIP     string // Which methods should be used for detecting the ip address of the UDP transport.
	SSUDetectedIP string // What ip has been detected by the UDP transport.
	NetUPNP       string // Is UPnP enabled.
	NetBWShare    string // How many percent of bandwidth is usable for participating tunnels.
	NetBWIn       string // How many KB/s of inbound bandwidth is allowed.
	NetBWOut      string // How many KB/s of outbound bandwidth is allowed.
	NetLaptopMode string // Is laptop mode enabled (change router identity and UDP port when IP changes).
}

// NetworkSettingReply contains a NetworkSetting reply.
type NetworkSettingReply struct {
	*NetworkSetting
	SettingsSaved bool // Have the provided settings been saved.
	RestartNeeded bool // Is a restart needed for the new settings to be used.
}

// NetworkSetting fetches or sets various network related settings. Ports, addresses etc.
func (c *Client) NetworkSetting(ns NetworkSetting) (NetworkSettingReply, error) {
	params := make(map[string]interface{})
	if ns.NTCPPort != "" {
		params["i2p.router.net.ntcp.port"] = ns.NTCPPort
	}
	if ns.NTCPHostname != "" {
		params["i2p.router.net.ntcp.hostname"] = ns.NTCPHostname
	}
	if ns.NTCPAutoIP != "" {
		params["i2p.router.net.ntcp.autoip"] = ns.NTCPAutoIP
	}
	if ns.SSUPort != "" {
		params["i2p.router.net.ssu.port"] = ns.SSUPort
	}
	if ns.SSUHostname != "" {
		params["i2p.router.net.ssu.hostname"] = ns.SSUHostname
	}
	if ns.SSUAutoIP != "" {
		params["i2p.router.net.ssu.autoip"] = ns.SSUAutoIP
	}
	if ns.SSUDetectedIP != "" {
		params["i2p.router.net.ssu.detectedip"] = ns.SSUDetectedIP
	}
	if ns.NetUPNP != "" {
		params["i2p.router.net.upnp"] = ns.NetUPNP
	}
	if ns.NetBWShare != "" {
		params["i2p.router.net.bw.share"] = ns.NetBWShare
	}
	if ns.NetBWIn != "" {
		params["i2p.router.net.bw.in"] = ns.NetBWIn
	}
	if ns.NetBWOut != "" {
		params["i2p.router.net.bw.out"] = ns.NetBWOut
	}
	if ns.NetLaptopMode != "" {
		params["i2p.router.net.laptopmode"] = ns.NetLaptopMode
	}

	reply, err := c.Send("NetworkSetting", params)
	result := NetworkSettingReply{}

	if err != nil || reply.Result == nil {
		return result, fmt.Errorf("NetworkSetting failed: %w", err)
	}
	if reply.Result["i2p.router.net.ntcp.port"] != "" {
		result.NTCPPort = reply.Result["i2p.router.net.ntcp.port"].(string)
	}
	if reply.Result["i2p.router.net.ntcp.hostname"] != "" {
		result.NTCPHostname = reply.Result["i2p.router.net.ntcp.hostname"].(string)
	}
	if reply.Result["i2p.router.net.ntcp.autoip"] != "" {
		result.NTCPAutoIP = reply.Result["i2p.router.net.ntcp.autoip"].(string)
	}
	if reply.Result["i2p.router.net.ssu.port"] != "" {
		result.SSUPort = reply.Result["i2p.router.net.ssu.port"].(string)
	}
	if reply.Result["i2p.router.net.ssu.hostname"] != "" {
		result.SSUHostname = reply.Result["i2p.router.net.ssu.hostname"].(string)
	}
	if reply.Result["i2p.router.net.ssu.autoip"] != "" {
		result.SSUAutoIP = reply.Result["i2p.router.net.ssu.autoip"].(string)
	}
	if reply.Result["i2p.router.net.ssu.detectedip"] != "" {
		result.SSUDetectedIP = reply.Result["i2p.router.net.ssu.detectedip"].(string)
	}
	if reply.Result["i2p.router.net.upnp"] != "" {
		result.NetUPNP = reply.Result["i2p.router.net.upnp"].(string)
	}
	if reply.Result["i2p.router.net.bw.share"] != "" {
		result.NetBWShare = reply.Result["i2p.router.net.bw.share"].(string)
	}
	if reply.Result["i2p.router.net.bw.in"] != "" {
		result.NetBWIn = reply.Result["i2p.router.net.bw.in"].(string)
	}
	if reply.Result["i2p.router.net.bw.out"] != "" {
		result.NetBWOut = reply.Result["i2p.router.net.bw.out"].(string)
	}
	if reply.Result["i2p.router.net.laptopmode"] != "" {
		result.NetLaptopMode = reply.Result["i2p.router.net.laptopmode"].(string)
	}
	return result, err
}

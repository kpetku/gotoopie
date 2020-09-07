package backend

import (
	"fmt"
)

const apiVersion float64 = 1 // The version of the I2PControl API used by the client
const nonEmptyString = "x"
const nonEmptyint = 1

// Authenticate creates and returns an authentication token used for further communication.
func (c *Client) Authenticate() error {
	reply, err := c.Send("Authenticate", map[string]interface{}{
		"API":      apiVersion,
		"Password": c.Config.Password,
	})
	if err != nil || reply.Result == nil {
		return fmt.Errorf("Authenticate failed: %w", err)
	}
	c.Config.Token = reply.Result["Token"].(string)
	return err
}

// Echo echoes the value of the echo key, used for debugging and testing.
func (c *Client) Echo(echo string) (string, error) {
	reply, err := c.Send("Echo", map[string]interface{}{
		"Echo": echo,
	})
	if err != nil || reply.Result == nil {
		return "", fmt.Errorf("Echo failed: %w", err)
	}
	return reply.Result["Result"].(string), err
}

// GetRate fetches rateStat from router statManager. Creates stat if not already created.
func (c *Client) GetRate(stat string, period int) (int, error) {
	reply, err := c.Send("GetRate", map[string]interface{}{
		"Stat":   stat,
		"Period": float64(period),
	})
	if err != nil || reply.Result == nil {
		return 0, fmt.Errorf("GetRate failed: %w", err)
	}
	return int(reply.Result["Result"].(float64)), err
}

// I2PControl wrappers

// SetAddress sets a new listen address for I2PControl (only 127.0.0.1 and 0.0.0.0 are implemented in I2PControl currently).
func (c *Client) SetAddress(address string) (bool, bool, error) {
	reply, err := c.I2PControl(I2PControl{
		Address: address,
	})
	if err != nil {
		return false, false, fmt.Errorf("SetAddress failed: %w", err)
	}
	return reply.RestartNeeded, reply.SettingsSaved, err
}

// SetPassword ets a new password for I2PControl, all Authentication tokens will be revoked.
func (c *Client) SetPassword(password string) (bool, bool, error) {
	reply, err := c.I2PControl(I2PControl{
		Password: password,
	})
	if err != nil {
		return false, false, fmt.Errorf("SetPassword failed: %w", err)
	}
	return reply.RestartNeeded, reply.SettingsSaved, err
}

// SetPort switches which port I2PControl will listen for connections on.
func (c *Client) SetPort(port string) (bool, bool, error) {
	reply, err := c.I2PControl(I2PControl{
		Port: port,
	})
	if err != nil {
		return false, false, fmt.Errorf("SetPort failed: %w", err)
	}
	return reply.RestartNeeded, reply.SettingsSaved, err
}

// RouterInfo wrappers

// RouterStatus returns what the status of the router is.
func (c *Client) RouterStatus() (string, error) {
	reply, err := c.RouterInfo(RouterInfo{
		Status: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("RouterStatus failed: %w", err)
	}
	return reply.Status, err
}

// RouterUptime returns what the uptime of the router is in ms.
func (c *Client) RouterUptime() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		Uptime: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("RouterUptime failed: %w", err)
	}
	return reply.Uptime, err
}

// RouterVersion returns what version of I2P the router is running.
func (c *Client) RouterVersion() (string, error) {
	reply, err := c.RouterInfo(RouterInfo{
		Version: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("RouterVersion failed: %w", err)
	}
	return reply.Version, err
}

// NetBwInbound1s returns the 1 second average inbound bandwidth in Bps.
func (c *Client) NetBwInbound1s() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetBwInbound1s: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetBwInbound1s failed: %w", err)
	}
	return reply.NetBwInbound1s, err
}

// NetBwInbound15s returns the 15 second average inbound bandwidth in Bps.
func (c *Client) NetBwInbound15s() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetBwInbound15s: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetBwInbound15s failed: %w", err)
	}
	return reply.NetBwInbound15s, err
}

// NetBwOutbound1s returns the 1 second average outbound bandwidth in Bps.
func (c *Client) NetBwOutbound1s() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetBwOutbound1s: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetBwOutbound1s failed: %w", err)
	}
	return reply.NetBwOutbound1s, err
}

// NetBwOutbound15s returns the 15 second average outbound bandwidth in Bps
func (c *Client) NetBwOutbound15s() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetBwOutbound15s: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetBwOutbound15s failed: %w", err)
	}
	return reply.NetBwOutbound15s, err
}

// NetStatus returns what the current network status is.
func (c *Client) NetStatus() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetStatus: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetStatus failed: %w", err)
	}
	return reply.NetStatus, err
}

// NetTunnelsParticipating returns how many tunnels on the I2P net are we participating in.
func (c *Client) NetTunnelsParticipating() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetTunnelsParticipating: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetTunnelsParticipating failed: %w", err)
	}
	return reply.NetTunnelsParticipating, err
}

// NetDBActivePeers returns how many tunnels on the I2P net are we participating in.
func (c *Client) NetDBActivePeers() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetDBActivePeers: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetDBActivePeers failed: %w", err)
	}
	return reply.NetDBActivePeers, err
}

// NetDBFastPeers returns how many peers are considered 'fast'.
func (c *Client) NetDBFastPeers() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetDBFastPeers: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetDBFastPeers failed: %w", err)
	}
	return reply.NetDBFastPeers, err
}

// NetDBHighCapacityPeers returns how many peers are considered 'high capacity'.
func (c *Client) NetDBHighCapacityPeers() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetDBHighCapacityPeers: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetDBHighCapacityPeers failed: %w", err)
	}
	return reply.NetDBHighCapacityPeers, err
}

// NetDBIsReseeding returns is the router reseeding hosts to its NetDB?
func (c *Client) NetDBIsReseeding() (bool, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetDBIsReseeding: true,
	})
	if err != nil {
		return false, fmt.Errorf("NetDBIsReseeding failed: %w", err)
	}
	return reply.NetDBIsReseeding, err
}

// NetDBKnownPeers returns how many peers are known to us (listed in our NetDB).
func (c *Client) NetDBKnownPeers() (int, error) {
	reply, err := c.RouterInfo(RouterInfo{
		NetDBKnownPeers: nonEmptyint,
	})
	if err != nil {
		return 0, fmt.Errorf("NetDBKnownPeers failed: %w", err)
	}
	return reply.NetDBKnownPeers, err
}

// RouterManager wrappers

// BlockingFindUpdates initiates a search for signed updates.
func (c *Client) BlockingFindUpdates() (bool, error) {
	reply, err := c.RouterManager(RouterManager{
		FindUpdates: true,
	})
	if err != nil {
		return false, fmt.Errorf("BlockingFindUpdates failed: %w", err)
	}
	return reply.FindUpdates, err
}

// Reseed initiates a router reseed, fetching peers into our NetDB from a remote host.
func (c *Client) Reseed() error {
	_, err := c.RouterManager(RouterManager{
		Reseed: "",
	})
	if err != nil {
		return fmt.Errorf("Reseed failed: %w", err)
	}
	return err
}

// Restart restarts the router.
func (c *Client) Restart() error {
	_, err := c.RouterManager(RouterManager{
		Restart: "",
	})
	if err != nil {
		return fmt.Errorf("Restart failed: %w", err)
	}
	return err
}

// RestartGraceful restarts the router gracefully (waits for participating tunnels to expire).
func (c *Client) RestartGraceful() error {
	_, err := c.RouterManager(RouterManager{
		RestartGraceful: "",
	})
	if err != nil {
		return fmt.Errorf("RestartGraceful failed: %w", err)
	}
	return err
}

// Shutdown shuts down the router.
func (c *Client) Shutdown() error {
	_, err := c.RouterManager(RouterManager{
		Shutdown: "",
	})
	if err != nil {
		return fmt.Errorf("Shutdown failed: %w", err)
	}
	return err
}

// ShutdownGraceful shuts down the router.
func (c *Client) ShutdownGraceful() error {
	_, err := c.RouterManager(RouterManager{
		ShutdownGraceful: "",
	})
	if err != nil {
		return fmt.Errorf("ShutdownGraceful failed: %w", err)
	}
	return err
}

// Update initiates a router update from signed sources.
func (c *Client) Update() error {
	_, err := c.RouterManager(RouterManager{
		Update: "",
	})
	if err != nil {
		return fmt.Errorf("Update failed: %w", err)
	}
	return err
}

// NetworkSetting wrappers

// NTCPPort gets/sets the port used for the TCP transport.
func (c *Client) NTCPPort(port string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NTCPPort: port,
	})
	if err != nil {
		return "", fmt.Errorf("NTCPPort failed: %w", err)
	}
	return reply.NTCPPort, err
}

// NTCPHostname gets/sets hostname is used for the TCP transport.
func (c *Client) NTCPHostname(hostname string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NTCPHostname: hostname,
	})
	if err != nil {
		return "", fmt.Errorf("NTCPHostname failed: %w", err)
	}
	return reply.NTCPHostname, err
}

// NTCPAutoIP use automatically detected ip for TCP transport.
func (c *Client) NTCPAutoIP(setting string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NTCPAutoIP: setting,
	})
	if err != nil {
		return "", fmt.Errorf("NTCPAutoIP failed: %w", err)
	}
	return reply.NTCPAutoIP, err
}

// SSUPort what port is used for the UDP transport.
func (c *Client) SSUPort(port string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		SSUPort: port,
	})
	if err != nil {
		return "", fmt.Errorf("SSUPort failed: %w", err)
	}
	return reply.SSUPort, err
}

// SSUHostname what hostname is used for the UDP transport.
func (c *Client) SSUHostname(hostname string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		SSUHostname: hostname,
	})
	if err != nil {
		return "", fmt.Errorf("SSUHostname failed: %w", err)
	}
	return reply.SSUHostname, err
}

// SSUAutoIP which methods should be used for detecting the ip address of the UDP transport.
func (c *Client) SSUAutoIP(setting string) (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		SSUAutoIP: setting,
	})
	if err != nil {
		return "", fmt.Errorf("SSUAutoIP failed: %w", err)
	}
	return reply.SSUAutoIP, err
}

// SSUDetectedIP what ip has been detected by the UDP transport.
func (c *Client) SSUDetectedIP() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		SSUDetectedIP: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("SSUDetectedIP failed: %w", err)
	}
	return reply.SSUDetectedIP, nil
}

// NetUPNP is UPnP enabled?
func (c *Client) NetUPNP() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NetUPNP: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("NetUPNP failed: %w", err)
	}
	return reply.SSUDetectedIP, nil
}

// NetBWShare how many percent of bandwidth is usable for participating tunnels.
func (c *Client) NetBWShare() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NetBWShare: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("NetBWShare failed: %w", err)
	}
	return reply.NetBWShare, nil
}

// NetBWIn how many KB/s of inbound bandwidth is allowed.
func (c *Client) NetBWIn() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NetBWIn: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("NetBWIn failed: %w", err)
	}
	return reply.NetBWIn, nil
}

// NetBWOut How many KB/s of outbound bandwidth is allowed.
func (c *Client) NetBWOut() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NetBWOut: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("NetBWOut failed: %w", err)
	}
	return reply.NetBWOut, nil
}

// NetLaptopMode is laptop mode enabled (change router identity and UDP port when IP changes).
func (c *Client) NetLaptopMode() (string, error) {
	reply, err := c.NetworkSetting(NetworkSetting{
		NetLaptopMode: nonEmptyString,
	})
	if err != nil {
		return "", fmt.Errorf("NetLaptopMode failed: %w", err)
	}
	return reply.NetLaptopMode, nil
}

// AdvancedSettings wrappers

// GetAll lists key value options for AdvancedSettings
func (c *Client) GetAll() (map[string]interface{}, error) {
	reply, err := c.Send("AdvancedSettings", map[string]interface{}{
		"getAll": nonEmptyString,
	})
	if err != nil || reply.Result == nil {
		return nil, fmt.Errorf("GetAll failed: %w", err)
	}
	return reply.Result, err
}

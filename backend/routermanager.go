package backend

import "fmt"

// RouterManager manages I2P router restart/shutdown.
type RouterManager struct {
	FindUpdates      bool   // Blocking. Initiates a search for signed updates.
	Reseed           string // Initiates a router reseed, fetching peers into our NetDB from a remote host.
	Restart          string // Restarts the router.
	RestartGraceful  string // Restarts the router gracefully (waits for participating tunnels to expire).
	Shutdown         string // Shuts down the router.
	ShutdownGraceful string // Shuts down the router gracefully (waits for participating tunnels to expire).
	Update           string // Initiates a router update from signed sources.
}

// RouterManager manages I2P router restart/shutdown.
func (c *Client) RouterManager(rm RouterManager) (RouterManager, error) {
	params := make(map[string]interface{})
	if rm.FindUpdates != false {
		params["FindUpdates"] = nil
	}
	if rm.Reseed != "" {
		params["Reseed"] = nil
	}
	if rm.Restart != "" {
		params["Restart"] = nil
	}
	if rm.RestartGraceful != "" {
		params["RestartGraceful"] = nil
	}
	if rm.Shutdown != "" {
		params["Shutdown"] = nil
	}
	if rm.ShutdownGraceful != "" {
		params["ShutdownGraceful"] = nil
	}
	if rm.Update != "" {
		params["Update"] = nil
	}

	reply, err := c.Send("RouterManager", params)
	result := RouterManager{}

	if err != nil || reply.Result == nil {
		return result, fmt.Errorf("RouterManager failed: %w", err)
	}
	if reply.Result["FindUpdates"] != "" {
		if reply.Result["FindUpdates"] == "true" {
			result.FindUpdates = true
		}
	}
	if reply.Result["Reseed"] != nil {
		result.Reseed = reply.Result["Reseed"].(string)
	}
	if reply.Result["Restart"] != nil {
		result.Restart = reply.Result["Restart"].(string)
	}
	if reply.Result["RestartGraceful"] != nil {
		result.RestartGraceful = reply.Result["RestartGraceful"].(string)
	}
	if reply.Result["Shutdown"] != nil {
		result.Shutdown = reply.Result["Shutdown"].(string)
	}
	if reply.Result["ShutdownGraceful"] != nil {
		result.ShutdownGraceful = reply.Result["ShutdownGraceful"].(string)
	}
	if reply.Result["Update"] != nil {
		result.Update = reply.Result["Update"].(string)
	}
	return result, err
}

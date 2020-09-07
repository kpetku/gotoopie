package backend

import "fmt"

// I2PControl manages I2PControl. Ports, passwords and the like.
type I2PControl struct {
	Address  string `json:",omitempty"` // Sets a new listen address for I2PControl (only 127.0.0.1 and 0.0.0.0 are implemented in I2PControl currently).
	Password string `json:",omitempty"` // Sets a new password for I2PControl, all Authentication tokens will be revoked.
	Port     string `json:",omitempty"` // Switches which port I2PControl will listen for connections on.
	Token    string `json:",omitempty"`
}

// I2PControlReply contains an I2PControl reply.
type I2PControlReply struct {
	*I2PControl
	SettingsSaved bool // Returns true if any changes were made.
	RestartNeeded bool // Returns true if any changes requiring a restart to take effect were made.
}

// I2PControl manages I2PControl. Ports, passwords and the like.
func (c *Client) I2PControl(i2pcontrol I2PControl) (I2PControlReply, error) {
	params := make(map[string]interface{})
	if i2pcontrol.Address != "" {
		params["i2pcontrol.address"] = i2pcontrol.Address
	}
	if i2pcontrol.Password != "" {
		params["i2pcontrol.password"] = i2pcontrol.Password
	}
	if i2pcontrol.Port != "" {
		params["i2pcontrol.port"] = i2pcontrol.Port
	}

	reply, err := c.Send("I2PControl", params)
	result := I2PControlReply{}

	if err != nil || reply.Result == nil {
		return result, fmt.Errorf("I2PControl failed: %w", err)
	}
	if reply.Result["i2pcontrol.address"] != nil {
		result.Address = reply.Result["Address"].(string)
	}
	if reply.Result["i2pcontrol.password"] != nil {
		result.Address = reply.Result["i2pcontrol.password"].(string)
	}
	if reply.Result["i2pcontrol.port"] != nil {
		result.Address = reply.Result["i2pcontrol.port"].(string)
	}
	if reply.Result["SettingsSaved"] == "true" {
		result.SettingsSaved = true
	}
	if reply.Result["RestartNeeded"] == "true" {
		result.RestartNeeded = true
	}
	return result, err
}

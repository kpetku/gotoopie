package backend

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

const jsonrpcVersion string = "2.0"
const scheme = "https://"

// Client is an I2PControl client
type Client struct {
	Config   *Options
	MsgCount int
	mutex    sync.RWMutex
}

// Options is atype alias of I2PControl
type Options = I2PControl

// NewClient creates a new I2PControl session with the configuration options Options
func NewClient(opts *Options) *Client {
	c := new(Client)
	c.Config = opts
	return c
}

// Request contains a JSONRPC2.0 I2PControl request
type Request struct {
	ID      string                 `json:"id"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	JSONRPC string                 `json:"jsonrpc"`
}

// Reply contains a Request reply
type Reply struct {
	ID      string                 `json:"id"`
	Result  map[string]interface{} `json:"result"`
	JSONRPC string                 `json:"jsonrpc"`
}

// Send sends a JSONRPC2.0 I2PControl Request down the wire and returns a Result
func (c *Client) Send(method string, params map[string]interface{}) (*Reply, error) {
	c.mutex.Lock()
	c.MsgCount++
	params["Token"] = c.Config.Token
	c.mutex.Unlock()
	return c.buildRequest(Request{
		ID:      strconv.Itoa(c.MsgCount),
		Method:  method,
		Params:  params,
		JSONRPC: jsonrpcVersion,
	})
}

func (c *Client) buildRequest(request Request) (*Reply, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return &Reply{}, err
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", scheme+c.Config.Address+":"+c.Config.Port, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	r := new(Reply)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&r)
	return r, nil
}

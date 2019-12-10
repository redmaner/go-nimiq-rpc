package nimiqrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/redmaner/go-nimiq-rpc/jsonrpc"
)

var (

	// ErrRespBodyEmpty is returned when the underlying HTTP response body is empty
	ErrRespBodyEmpty = errors.New("The HTTP response body was empty")

	// ErrResultUnexpected is returned when the expected result could not be read
	ErrResultUnexpected = errors.New("Unexpected result")

	// ErrUnauthorized is returned when the user is not authorized to call the function
	ErrUnauthorized = errors.New("Unauthorized to do the request")
)

// Client contains a Nimiq RPC client
type Client struct {

	// Address is the address of the RPC server / Nimiq node
	Address string

	// Transport is used to handle HTTP requests to the RPC server
	Transport http.RoundTripper

	// Headers contains the headers that will be copied over to each
	// JSON-RPC request. This allows to set custom headers for all requests
	// handled by the client.
	Headers http.Header
}

// NewClient returns a new Nimiq RPC client
func NewClient(address string) *Client {
	return &Client{
		Address: address,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   8 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Headers: make(map[string][]string),
	}
}

// RawCall can be used to send a raw JSON-RPC request. This function is used internally
// to handle all the RPC functions provided by the client. Therefore this function
// should generally not be used. It does however provide the functionality to do
// RPC requests that are (not yet) implemented by the client.
func (nc *Client) RawCall(req *jsonrpc.Request) (resp *jsonrpc.Response, err error) {

	// Marshall request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request
	httpReq, err := http.NewRequest("POST", nc.Address, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set Content-Type
	httpReq.Header.Set("Content-Type", "application/json")

	// Copy custom headers to request
	for k, vv := range nc.Headers {
		for _, v := range vv {
			httpReq.Header.Add(k, v)
		}
	}

	// Do the HTTP request
	httpResp, err := nc.Transport.RoundTrip(httpReq)
	if err != nil {
		return nil, err
	}

	if httpResp.Body == nil {
		return nil, ErrRespBodyEmpty
	}

	// Check status codes
	switch httpResp.StatusCode {
	case 403:
		return nil, ErrUnauthorized
	}

	// Parse body
	bodyData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshall response
	if resp == nil {
		resp = &jsonrpc.Response{}
	}
	err = json.Unmarshal(bodyData, resp)
	if err != nil {
		return nil, err
	}

	// Check for JSONRPC errors
	if err := resp.Error.Parse(); err != nil {
		return nil, err
	}

	return resp, nil
}

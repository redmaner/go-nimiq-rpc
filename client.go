// Copyright 2019 Jake "redmaner" van der Putten.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nimiqrpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	// ErrRespBodyEmpty is returned when the underlying HTTP response body is empty
	ErrRespBodyEmpty = errors.New("the HTTP response body was empty")

	// ErrResultUnexpected is returned when the expected result could not be read
	ErrResultUnexpected = errors.New("unexpected result")

	// ErrUnauthorized is returned when the user is not authorized to call the function
	ErrUnauthorized = errors.New("unauthorized")

	// ErrNotAuthenticated is returned when the user is required to be authenticated
	ErrNotAuthenticated = errors.New("not authenticated")
)

// Client contains a Nimiq RPC client
type Client struct {
	Address   string // URL of the RPC server
	Transport http.RoundTripper
	Headers   http.Header // additional headers on all requests

	id int64
}

// NewClient returns a new Nimiq RPC client
func NewClient(address string) *Client {
	return &Client{
		Address: address,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   8 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Headers: make(map[string][]string),
	}
}

// NewClientWithAuth returns a RPC client with the given username and password set for
// authentication
func NewClientWithAuth(address, username, password string) *Client {
	cl := NewClient(address)
	cl.Authenticate(username, password)
	return cl
}

// Authenticate can be used to set a username and password that is used to authenticate
// to the RPC server
func (nc *Client) Authenticate(username, password string) {
	nc.Headers.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))))
}

// RawCall can be used to send a raw JSON-RPC request. This function is used internally
// to handle all the RPC functions provided by the client. Therefore this function
// should generally not be used. It does however provide the functionality to do
// RPC requests that are (not yet) implemented by the client.
func (nc *Client) RawCall(req *RPCRequest) (resp *RPCResponse, err error) {
	requestID := atomic.AddInt64(&nc.id, 1)
	req.ID = requestID

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

	// Check status codes
	switch httpResp.StatusCode {
	case 401:
		return nil, ErrNotAuthenticated
	case 403:
		return nil, ErrUnauthorized
	}

	// Check if HTTP body exists
	if httpResp.Body == nil {
		return nil, ErrRespBodyEmpty
	}

	// Parse body
	bodyData, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response
	resp = new(RPCResponse)
	err = json.Unmarshal(bodyData, resp)
	if err != nil {
		return nil, err
	}

	// Validate ID match
	if requestID != resp.ID {
		return nil, ErrRPCIDMismatch
	}

	// Check for JSONRPC errors
	if err := resp.Error.Parse(); err != nil {
		return nil, err
	}

	return resp, nil
}

package nimiqrpc

import (
	"encoding/json"

	"github.com/redmaner/go-nimiq-rpc/jsonrpc"
)

// Accounts returns a list of addresses owned by client.
func (nc *Client) Accounts() (accounts []Account, err error) {

	// Make a new jsonrpc request
	rpcReq := jsonrpc.NewRequest("accounts", nil, jsonrpc.NewID())

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return []Account{}, err
	}

	// Unmarshal result
	var result []Account
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return []Account{}, ErrResultUnexpected
	}

	return result, nil
}

// BlockNumber returns the height of most recent block.
func (nc *Client) BlockNumber() (blockHeight int, err error) {

	// Make a new jsonrpc request
	rpcReq := jsonrpc.NewRequest("blockNumber", nil, jsonrpc.NewID())

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result int
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, ErrResultUnexpected
	}

	return result, nil
}

// Consensus returns information on the current consensus state.
func (nc *Client) Consensus() (consensus string, err error) {

	// Make a new jsonrpc request
	rpcReq := jsonrpc.NewRequest("consensus", nil, jsonrpc.NewID())

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

	// Unmarshal result
	var result string
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return "", ErrResultUnexpected
	}

	return result, nil
}

// CreateAccount creates a new account and stores its private key in the client store.
func (nc *Client) CreateAccount() (wallet *Wallet, err error) {

	// Make a new jsonrpc request
	rpcReq := jsonrpc.NewRequest("createAccount", nil, jsonrpc.NewID())

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Wallet
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	return &result, nil
}

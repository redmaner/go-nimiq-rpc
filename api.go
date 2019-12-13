package nimiqrpc

import (
	"encoding/json"
	"fmt"
)

// Accounts returns a list of addresses owned by client.
func (nc *Client) Accounts() (accounts []Account, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("accounts", nil)

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
	rpcReq := NewRPCRequest("blockNumber", nil)

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
	rpcReq := NewRPCRequest("consensus", nil)

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
	rpcReq := NewRPCRequest("createAccount", nil)

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

// CreateRawTransaction creates and signs a transaction without sending it.
// The transaction can then be send via sendRawTransaction without accidentally replaying it.
func (nc *Client) CreateRawTransaction(trn OutgoingTransaction) (transactionHex string, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("createRawTransaction", trn)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

	fmt.Println(string(rpcResp.Result))

	// Unmarshal result
	var result string
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

// GetAccount returns details for the account of given address.
func (nc *Client) GetAccount(address string) (account *Account, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getAccount", address)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Account
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	return &result, nil
}

// GetBalance returns the balance of the account of given address.
func (nc *Client) GetBalance(address string) (balance Luna, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getBalance", address)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result Luna
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, ErrResultUnexpected
	}

	return result, nil
}

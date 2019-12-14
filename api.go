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
	"encoding/json"
	"fmt"
)

// Accounts returns a list of addresses owned by client.
func (nc *Client) Accounts() (accounts []Account, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("accounts", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return []Account{}, err
	}

	// Unmarshal result
	var result []Account
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return []Account{}, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// BlockNumber returns the height of most recent block.
func (nc *Client) BlockNumber() (blockHeight int, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("blockNumber", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result int
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// Consensus returns information on the current consensus state.
func (nc *Client) Consensus() (consensus string, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("consensus", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

	// Unmarshal result
	var result string
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// CreateAccount creates a new account and stores its private key in the client store.
func (nc *Client) CreateAccount() (wallet *Wallet, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("createAccount", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Wallet
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// CreateRawTransaction creates and signs a transaction without sending it.
// The transaction can then be send via sendRawTransaction without accidentally replaying it.
func (nc *Client) CreateRawTransaction(trn OutgoingTransaction) (transactionHex string, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("createRawTransaction", trn)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

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

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getAccount", address)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Account
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBalance returns the balance of the account of given address.
func (nc *Client) GetBalance(address string) (balance Luna, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBalance", address)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result Luna
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// GetBlockByHash returns information about a block by block hash.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) GetBlockByHash(blockHash string, fullTransactions bool) (block *Block, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockHash)
	params = append(params, fullTransactions)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBlockByNumber", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	// Transaction result
	switch {
	case fullTransactions:
		err = json.Unmarshal(result.Transactions, &result.TransactionsObjects)
	default:
		err = json.Unmarshal(result.Transactions, &result.TransactionsHashes)
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBlockByNumber returns information about a block by block number.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) GetBlockByNumber(blockNumber int, fullTransactions bool) (block *Block, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockNumber)
	params = append(params, fullTransactions)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBlockByNumber", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	// Transaction result
	switch {
	case fullTransactions:
		err = json.Unmarshal(result.Transactions, &result.TransactionsObjects)
	default:
		err = json.Unmarshal(result.Transactions, &result.TransactionsHashes)
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil
}

// GetBlockTemplate returns a template to build the next block for mining.
// This will consider pool instructions when connected to a pool.
// Optional parameters: (1) The address to use as a miner for this block.
// This overrides the address provided during startup or from the pool.
// and (2)  Hex-encoded value for the extra data field. This overrides the address
// provided during startup or from the pool.
func (nc *Client) GetBlockTemplate(params ...interface{}) (template *BlockTemplate, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBlockTemplate", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result BlockTemplate
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return &result, nil

}

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (nc *Client) GetBlockTransactionCountByHash(blockHash string) (transactionCount int, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBlockTransactionCountByHash", blockHash)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result int
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (nc *Client) GetBlockTransactionCountByNumber(blockNumber int) (transactionCount int, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getBlockTransactionCountByNumber", blockNumber)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result int
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// GetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (nc *Client) GetTransactionByBlockHashAndIndex(blockHash string, index int) (transaction *Transaction, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockHash)
	params = append(params, index)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getTransactionByBlockHashAndIndex", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionByBlockNumberAndIndex returns information about a transaction by block hash and transaction index position.
func (nc *Client) GetTransactionByBlockNumberAndIndex(blockNumber int, index int) (transaction *Transaction, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockNumber)
	params = append(params, index)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getTransactionByBlockNumberAndIndex", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionByHash Returns the information about a transaction requested by transaction hash.
func (nc *Client) GetTransactionByHash(transactionHash string) (transaction *Transaction, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getTransactionByHash", transactionHash)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionReceipt returns the receipt of a transaction by transaction hash.
func (nc *Client) GetTransactionReceipt(transactionHash string) (transactionReceipt *TransactionReceipt, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getTransactionReceipt", transactionHash)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result TransactionReceipt
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.TransactionHash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionsByAddress returns the latest transactions successfully performed by or for an address.
// The array will not contain more than maxEntries, but might contain less, even when more transactions happened.
// Any interpretation of the length of this array might result in worng assumptions.
func (nc *Client) GetTransactionsByAddress(address string, maxEntries int) (transactions []Transaction, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, address)
	params = append(params, maxEntries)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getTransactionsByAddress", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result []Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// GetWork returns instructions to mine the next block. This will consider pool instructions when connected to a pool.
// Optional arameters: (1)  The address to use as a miner for this block. This overrides the address provided during startup or from the pool.
// and (2) Hex-encoded value for the extra data field. This overrides the address provided during startup or from the pool
func (nc *Client) GetWork(params ...interface{}) (work *Work, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("getWork", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Work
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Data == "" {
		return nil, nil
	}

	return &result, nil
}

// Hashrate returns the number of hashes per second that the node is mining with.
func (nc *Client) Hashrate() (hashrate float64, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("hashrate", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result float64
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// Log sets the log level of the node.
func (nc *Client) Log(tag string, level LogLevel) (succes bool, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, tag)
	params = append(params, level)

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("log", params)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return false, err
	}

	// Unmarshal result
	var result bool
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return false, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// Mempool Returns information on the current mempool situation.
// This will provide an overview of the number of transactions sorted into buckets
// based on their fee per byte (in smallest unit).
func (nc *Client) Mempool() (mempool *Mempool, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("mempool", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Mempool
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	if result.Total == 0 && len(result.Buckets) == 0 {
		return nil, nil
	}

	return &result, nil
}

// Mining returns if client is actively mining new blocks.
func (nc *Client) Mining() (status bool, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("mining", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return false, err
	}

	// Unmarshal result
	var result bool
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return false, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// PeerCount returns number of peers currently connected to the client.
func (nc *Client) PeerCount() (peers int, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("peerCount", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return 0, err
	}

	// Unmarshal result
	var result int
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return 0, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// SendRawTransaction sends a signed message call transaction or a contract creation, if the data field contains code.
func (nc *Client) SendRawTransaction(signedTransaction string) (transactionHash string, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("sendRawTransaction", signedTransaction)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

	// Unmarshal result
	var result string
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return "", fmt.Errorf("%v: %v", ErrResultUnexpected, err)
	}

	return result, nil
}

// SendTransaction creates new message call transaction or a contract creation, if the data field contains code.
func (nc *Client) SendTransaction(trn OutgoingTransaction) (transactionHash string, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("sendTransaction", trn)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return "", err
	}

	// Unmarshal result
	var result string
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

// SubmitBlock submits a block to the node. When the block is valid, the node will forward it to other nodes in the network.
// The argument is a hex-encoded full block (including header, interlink and body).
// When submitting work from getWork, remember to include the suffix.
func (nc *Client) SubmitBlock(fullBlock string) (err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("submitBlock", fullBlock)

	// Make JSON-RPC call
	_, err = nc.RawCall(rpcReq)
	if err != nil {
		return err
	}

	return nil
}

// Syncing returns whether the node is syncing and when it is syncing, data about the sync status.
func (nc *Client) Syncing() (syncing bool, syncStatus *SyncStatus, err error) {

	// Make a new JSON-RPC request
	rpcReq := NewRPCRequest("syncing", nil)

	// Make JSON-RPC call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return false, nil, err
	}

	// Unmarshal result
	var result SyncStatus
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {

		var boolResult bool
		err = json.Unmarshal(rpcResp.Result, &boolResult)
		if err != nil {
			return false, nil, fmt.Errorf("%v: %v", ErrResultUnexpected, err)
		}

		return false, nil, nil
	}

	return true, &result, nil
}

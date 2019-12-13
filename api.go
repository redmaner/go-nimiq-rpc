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

// GetBlockByHash returns information about a block by block hash.
// If fullTransactions is true it returns a block with the full transaction objects,
// if false only the hashes of the transactions will be returned.
func (nc *Client) GetBlockByHash(blockHash string, fullTransactions bool) (block *Block, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockHash)
	params = append(params, fullTransactions)

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getBlockByNumber", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
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
		return nil, ErrResultUnexpected
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

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getBlockByNumber", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Block
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
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
		return nil, ErrResultUnexpected
	}

	return &result, nil
}

// TODO: IMPLEMENT function GetBlockTemplate

// GetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (nc *Client) GetBlockTransactionCountByHash(blockHash string) (transactionCount int, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getBlockTransactionCountByHash", blockHash)

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

// GetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block number.
func (nc *Client) GetBlockTransactionCountByNumber(blockNumber int) (transactionCount int, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getBlockTransactionCountByNumber", blockNumber)

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

// GetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (nc *Client) GetTransactionByBlockHashAndIndex(blockHash string, index int) (transaction *Transaction, err error) {

	// Encapsulate parameters in a interface slice
	var params []interface{}
	params = append(params, blockHash)
	params = append(params, index)

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getTransactionByBlockHashAndIndex", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
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

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getTransactionByBlockNumberAndIndex", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionByHash Returns the information about a transaction requested by transaction hash.
func (nc *Client) GetTransactionByHash(transactionHash string) (transaction *Transaction, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getTransactionByHash", transactionHash)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	if result.Hash == "" {
		return nil, nil
	}

	return &result, nil
}

// GetTransactionReceipt returns the receipt of a transaction by transaction hash.
func (nc *Client) GetTransactionReceipt(transactionHash string) (transactionReceipt *TransactionReceipt, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getTransactionReceipt", transactionHash)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result TransactionReceipt
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
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

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getTransactionsByAddress", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result []Transaction
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	return result, nil
}

// GetWork returns instructions to mine the next block. This will consider pool instructions when connected to a pool.
// Optional arameters: (1)  The address to use as a miner for this block. This overrides the address provided during startup or from the pool.
// and (2) Hex-encoded value for the extra data field. This overrides the address provided during startup or from the pool
func (nc *Client) GetWork(params ...interface{}) (work *Work, err error) {

	// Make a new jsonrpc request
	rpcReq := NewRPCRequest("getWork", params)

	// Make jsonrpc call
	rpcResp, err := nc.RawCall(rpcReq)
	if err != nil {
		return nil, err
	}

	// Unmarshal result
	var result Work
	err = json.Unmarshal(rpcResp.Result, &result)
	if err != nil {
		return nil, ErrResultUnexpected
	}

	if result.Data == "" {
		return nil, nil
	}

	return &result, nil
}

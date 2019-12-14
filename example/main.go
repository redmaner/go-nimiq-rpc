package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	nimiqrpc "github.com/redmaner/go-nimiq-rpc"
)

var (

	// Commands
	cmdAccounts                            = flag.Bool("accounts", false, "Returns a list of addresses owned by client.")
	cmdBlockNumber                         = flag.Bool("blockNumber", false, "Returns the height of most recent block.")
	cmdConsensus                           = flag.Bool("consensus", false, "Returns information on the current consensus state")
	cmdCreateAccount                       = flag.Bool("createAccount", false, "Creates a new account and stores its private key in the client store.")
	cmdCreateRawTransaction                = flag.Bool("createRawTransaction", false, "Creates and signs a transaction without sending it.")
	cmdGetAccount                          = flag.Bool("getAccount", false, "Returns details for the account of given address.")
	cmdGetBalance                          = flag.Bool("getBalance", false, "Returns the balance of the account of given address.")
	cmdGetBlockByHash                      = flag.Bool("getBlockByHash", false, "Returns information about a block by hash.")
	cmdGetBlockByNumber                    = flag.Bool("getBlockByNumber", false, "Returns information about a block by block number.")
	cmdGetBlockTemplate                    = flag.Bool("getBlockTemplate", false, "Returns a template to build the next block for mining.")
	cmdGetBlockTransactionCountByHash      = flag.Bool("getBlockTransactionCountByHash", false, "Returns the number of transactions in a block from a block matching the given block hash.")
	cmdGetBlockTransactionCountByNumber    = flag.Bool("getBlockTransactionCountByNumber", false, "Returns the number of transactions in a block matching the given block number.")
	cmdGetTransactionByBlockHashAndIndex   = flag.Bool("getTransactionByBlockHashAndIndex", false, "Returns information about a transaction by block hash and transaction index position.")
	cmdGetTransactionByBlockNumberAndIndex = flag.Bool("getTransactionByBlockNumberAndIndex", false, "Returns information about a transaction by block number and transaction index position.")
	cmdGetTransactionByHash                = flag.Bool("getTransactionByHash", false, "Returns the information about a transaction requested by transaction hash.")
	cmdGetTransactionReceipt               = flag.Bool("getTransactionReceipt", false, "Returns the receipt of a transaction by transaction hash.")
	cmdGetTransactionsByAddress            = flag.Bool("getTransactionsByAddress", false, "Returns the latest transactions successfully performed by or for an address.")
	cmdGetWork                             = flag.Bool("getWork", false, "Returns instructions to mine the next block. This will consider pool instructions when connected to a pool.")
	cmdHashrate                            = flag.Bool("hashrate", false, "Returns the number of hashes per second that the node is mining with.")
	cmdLog                                 = flag.Bool("log", false, "Sets the log level of the node.")
	cmdMempool                             = flag.Bool("mempool", false, "Returns information on the current mempool situation.")
	cmdMining                              = flag.Bool("mining", false, "Returns true if client is actively mining new blocks.")
	cmdPeerCount                           = flag.Bool("peerCount", false, "Returns number of peers currently connected to the client.")
	cmdSyncing                             = flag.Bool("syncing", false, "Returns data about the sync status.")

	argAddress          = flag.String("addr", "", "The Nimiq account address")
	argAddressFrom      = flag.String("addrFrom", "", "The Nimiq account address from")
	argAddressTo        = flag.String("addrTo", "", "The Nimiq account address to")
	argValue            = flag.Int("value", 0, "Value used by various functions as argument")
	argHash             = flag.String("hash", "", "The hash of a block or transaction")
	argNumber           = flag.Int("number", 0, "The number of a block")
	argFullTransactions = flag.Bool("fullTransactions", false, "Retrieve full transactions of a block")
	argLogLevel         = flag.String("level", "", "Log level")
	argLogTag           = flag.String("tag", "", "Log tag")
)

func main() {

	flag.Parse()
	nc := nimiqrpc.NewClient("http://nimiq.example.node.com:1234")
	nc.Headers.Set("Origin", "")

	switch {
	case *cmdAccounts:
		result, err := nc.Accounts()
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Accounts: %v\n", result)

	case *cmdBlockNumber:
		result, err := nc.BlockNumber()
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Block number: %v\n", result)

	case *cmdConsensus:
		result, err := nc.Consensus()
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Consensus: %v\n", result)

	case *cmdCreateAccount:
		result, err := nc.CreateAccount()
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("ID: %s\nAddress: %s\nPublicKey: %s\nPrivateKey: %s\n", result.ID, result.Address, result.PublicKey, result.PrivateKey)

	case *cmdCreateRawTransaction:

		if *argAddressTo == "" || *argAddressFrom == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		trn := nimiqrpc.OutgoingTransaction{
			From:  *argAddressFrom,
			To:    *argAddressTo,
			Fee:   0,
			Value: *argValue,
		}

		rawResult, err := nc.CreateRawTransaction(trn)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Transation in HEX: %v\n", rawResult)

	case *cmdGetAccount:
		if *argAddress == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetAccount(*argAddress)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Account: %v\n", *result)

	case *cmdGetBalance:
		if *argAddress == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetBalance(*argAddress)
		if err != nil {
			log.Panic(err)
		}
		resultNim := result.ToNIM()
		resultLuna := resultNim.ToLuna()
		fmt.Printf("Balance raw: %d\nBalance in NIM: %v\nBalance in Luna: %d\n", result, resultNim, resultLuna)

	case *cmdGetBlockByHash:
		if *argHash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetBlockByHash(*argHash, *argFullTransactions)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("block: %v\n", *result)

	case *cmdGetBlockByNumber:
		if *argNumber == 0 {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetBlockByNumber(*argNumber, *argFullTransactions)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("block: %v\n", *result)

	case *cmdGetBlockTemplate:

		result, err := nc.GetBlockTemplate()
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("block template: %v\n", *result)

	case *cmdGetBlockTransactionCountByHash:
		if *argHash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetBlockTransactionCountByHash(*argHash)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Transactions in block: %v\n", result)

	case *cmdGetBlockTransactionCountByNumber:
		if *argNumber == 0 {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetBlockTransactionCountByNumber(*argNumber)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("Transactions in block: %v\n", result)

	case *cmdGetTransactionByBlockHashAndIndex:
		if *argHash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetTransactionByBlockHashAndIndex(*argHash, *argValue)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Transaction: %v\n", *result)

	case *cmdGetTransactionByBlockNumberAndIndex:
		if *argNumber == 0 {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetTransactionByBlockNumberAndIndex(*argNumber, *argValue)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Transaction: %v\n", *result)

	case *cmdGetTransactionByHash:
		if *argHash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetTransactionByHash(*argHash)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Transaction: %v\n", *result)

	case *cmdGetTransactionReceipt:
		if *argHash == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		result, err := nc.GetTransactionReceipt(*argHash)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Transaction receipt: %v\n", *result)

	case *cmdGetTransactionsByAddress:
		if *argAddress == "" {
			flag.PrintDefaults()
			os.Exit(1)
		}

		if *argValue == 0 {
			*argValue = 1000
		}

		result, err := nc.GetTransactionsByAddress(*argAddress, *argValue)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Transactions: %v\n", result)

	case *cmdGetWork:

		result, err := nc.GetWork()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Work: %v\n", *result)

	case *cmdHashrate:

		result, err := nc.Hashrate()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Hashrate: %v\n", result)

	case *cmdLog:

		if *argLogTag == "" || *argLogLevel == "" {
			flag.PrintDefaults()
			os.Exit(2)
		}

		result, err := nc.Log(*argLogTag, nimiqrpc.LogLevel(*argLogLevel))
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Log set succes: %v\n", result)

	case *cmdMempool:

		result, err := nc.Mempool()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Mempool information: %v\n", *result)

	case *cmdMining:

		result, err := nc.Mining()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Mining: %v\n", result)

	case *cmdPeerCount:

		result, err := nc.PeerCount()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Peers: %v\n", result)

	case *cmdSyncing:
		status, result, err := nc.Syncing()
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("Syncing: %v\nSync status: %v\n", status, result)

	default:
		flag.PrintDefaults()
	}
}

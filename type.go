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

import "encoding/json"

const (

	// LogLevelTrace represents the trace log level
	LogLevelTrace LogLevel = "trace"

	// LogLevelVerbose represents the verbose log level
	LogLevelVerbose LogLevel = "verbose"

	// LogLevelDebug represents the debug log level
	LogLevelDebug LogLevel = "debug"

	// LogLevelInfo represents the info log level
	LogLevelInfo LogLevel = "info"

	// LogLevelWarn represents the warn log level
	LogLevelWarn LogLevel = "warn"

	// LogLevelError represents the error log level
	LogLevelError LogLevel = "error"

	// LogLevelAssert represents the assert log level
	LogLevelAssert LogLevel = "assert"
)

// LogLevel is de level of logging that is enabled on a node
type LogLevel string

// 100000 Luna is 1 NIM
// See https://www.nimiq.com/whitepaper/#nimiq-supply-distribution
const lunaNimRate float64 = 100000.00000

// NIM is the token transacted within Nimiq as a store and transfer of value: it acts as digital cash
type NIM float64

// ToLuna converts NIM to Luna
func (n *NIM) ToLuna() Luna {
	return Luna(float64(*n) * lunaNimRate)
}

// Luna is the smallest unit of NIM and 100’000 (1e5) Luna equals 1 NIM
type Luna int64

// ToNIM converts Luna to NIM
func (l *Luna) ToNIM() NIM {
	return NIM(float64(*l) / lunaNimRate)
}

// Account holds the details on an account
type Account struct {

	// ID - Hex-encoded 20 byte address
	ID string `json:"id,omitempty"`

	// address: String - User friendly address (NQ-address).
	Address string `json:"address,omitempty"`

	// balance: Integer - Balance of the account (in smallest unit).
	Balance Luna `json:"balance,omitempty"`

	// type: Integer - The account type associated with the account (BASIC: 0, VESTING: 1, HTLC: 2).
	Type int `json:"type,omitempty"`

	/*
	 * Additional fields for vesting contracts (type = 1)
	 */

	// owner: String - Hex-encoded 20 byte address of the owner of the vesting contract.
	Owner string `json:"owner,omitempty"`

	// ownerAddress: String - User friendly address (NQ-address) of the owner of the vesting contract.
	OwnerAddress string `json:"ownerAddress,omitempty"`

	// vestingStart: Integer - The block that the vesting contracted commenced.
	VestingStart int `json:"vestingStart,omitempty"`

	// vestingStepBlocks: Integer - The number of blocks after which some part of the vested funds is released.
	VestingStepBlocks int `json:"vestingStepBlocks,omitempty"`

	// vestingStepAmount: Integer - The amount (in smallest unit) released every vestingStepBlocks blocks.
	VestingStepAmount int `json:"vestingStepAmount,omitempty"`

	// vestingTotalAmount: Integer - The total amount (in smallest unit) that was provided at the contract creation.
	VestingTotalAmount int `json:"vestingTotalAmount,omitempty"`

	/*
	 * Additional fields for Hashed Time-Locked Contracts (type = 2)
	 */

	// sender: String - Hex-encoded 20 byte address of the sender of the HTLC.
	Sender string `json:"sender,omitempty"`

	// senderAddress: String - User friendly address (NQ-address) of the sender of the HTLC.
	SenderAddress string `json:"senderAddress,omitempty"`

	// recipient: String - Hex-encoded 20 byte address of the recipient of the HTLC.
	Recipient string `json:"recipient,omitempty"`

	// recipientAddress: String - User friendly address (NQ-address) of the recipient of the HTLC.
	RecipientAddress string `json:"recipientAddress,omitempty"`

	// hashRoot: String - Hex-encoded 32 byte hash root.
	HashRoot string `json:"hashRoot,omitempty"`

	// hashCount: Integer - Number of hashes this HTLC is split into
	HashCount int `json:"hashCount,omitempty"`

	// timeout: Integer - Block after which the contract can only be used by the original sender to recover funds.
	Timeout int `json:"timeout,omitempty"`

	// totalAmount: Integer - The total amount (in smallest unit) that was provided at the contract creation.
	TotalAmount int `json:"totalAmount,omitempty"`
}

// AddressObject holds the representation of a Nimiq address in two formats.
type AddressObject struct {

	// id: String - Hex-encoded 20 byte address.
	ID string `json:"id,omitempty"`

	// address: String - User friendly address (NQ-address).
	Address string `json:"address,omitempty"`
}

// Block holds the details on a block
type Block struct {

	// number: Integer - Height of the block.
	Number int `json:"number,omitempty"`

	// hash: String - Hex-encoded 32-byte hash of the block.
	Hash string `json:"hash,omitempty"`

	// pow: String - Hex-encoded 32-byte Proof-of-Work hash of the block.
	POW string `json:"pow,omitempty"`

	// parentHash: String - Hex-encoded 32-byte hash of the predecessor block.
	ParentHash string `json:"ParentHash,omitempty"`

	// nonce: Integer - The nonce of the block used to fulfill the Proof-of-Work.
	Nonce int `json:"nonce,omitempty"`

	// bodyHash: String - Hex-encoded 32-byte hash of the block body merkel root.
	BodyHash string `json:"bodyHash,omitempty"`

	// accountHash: String - Hex-encoded 32-byte hash of the accounts tree root.
	AccountHash string `json:"accountHash,omitempty"`

	// miner: String - Hex-encoded 20 byte address of the miner of the block.
	Miner string `json:"miner,omitempty"`

	// minerAddress: String - User friendly address (NQ-address) of the miner of the block.
	MinerAddress string `json:"minerAddress,omitempty"`

	// difficulty: String - Block difficulty, encoded as decimal number in string. (TODO)
	Difficulty string `json:"difficulty,omitempty"`

	// extraData: String - Hex-encoded value of the extra data field, maximum of 255 bytes.
	ExtraData string `json:"extraData,omitempty"`

	// size: Integer - Block size in byte.
	Size int `json:"size,omitempty"`

	// timestamp: Integer - UNIX timestamp of the block
	Timestamp int `json:"timestamp,omitempty"`

	// transactions: Array - Array of transactions. Either represented by the transaction hash or a Transaction object.
	// Because Transactions can either be a slice of strings or a slice of transactions,
	// transactions will be unmarshalled in a raw JSON message so that the message can be
	// unmarshalled into TransactionsHashes or TransactionsObjects depending on the data type.
	// The functions that return a Block, will automatically fill either TransactionsHashes or TransactionsObjects
	Transactions json.RawMessage `json:"transactions,omitempty"`

	// transactionsHashes contains a slice of transactions represented by the transaction hash
	TransactionsHashes []string

	// transactionsObjects contains a slice of transactions represented by transaction objects
	TransactionsObjects []Transaction
}

// BlockTemplate contains details on a block template
type BlockTemplate struct {

	// Block header
	Header struct {

		// version: Integer - Version in block header.
		Version int `json:"version,omitempty"`

		// prevHash: String - 32-byte hex-encoded hash of the previous block.
		PrevHash string `json:"prevHash,omitempty"`

		// interlinkHash: String - 32-byte hex-encoded hash of the interlink.
		InterlinkHash string `json:"interlinkHash,omitempty"`

		// accountsHash: String - 32-byte hex-encoded hash of the accounts tree.
		AccountHash string `json:"accountHash,omitempty"`

		// nBits: Integer - Compact form of the hash target for this block.
		NBits int `json:"nBits,omitempty"`

		// height: Integer - Height of the block in the block chain (also known as block number).
		Height int `json:"height,omitempty"`
	} `json:"header,omitempty"`

	// interlink: String - Hex-encoded interlink
	Interlink string `json:"interlink,omitempty"`

	// Block template body
	Body struct {

		// hash: String - 32-byte hex-encoded hash of the block body.
		Hash string `json:"hash,omitempty"`

		// minerAddr: String - 20-byte hex-encoded miner address.
		MinerAddr string `json:"minerAddr,omitempty"`

		// extraData: String - Hex-encoded value of the extra data field.
		ExtraData string `json:"extraData,omitempty"`

		// transactions: String[] - Array of hex-encoded transactions for this block.
		Transactions []string `json:"transactions,omitempty"`

		// prunedAccounts: String[] - Array of hex-encoded pruned accounts for this block.
		PrunedAccounts []string `json:"prunedAccounts,omitempty"`

		// merkleHashes: String[] - Array of hex-encoded hashes that verify the path of the miner
		// address in the merkle tree. This can be used to change the miner address easily.
		MerkleHashes []string `json:"merkleHashes,omitempty"`
	} `json:"body,omitempty"`

	// target: Integer - Compact form of the hash target to submit a block to this client.
	Target int `json:"target,omitempty"`
}

// Mempool holds the details on a mempool.
type Mempool struct {

	//total: Integer - Total number of pending transactions in mempool.
	Total int `json:"total,omitempty"`

	// buckets Integer[] - Array containing a subset of fee per byte buckets from
	// [10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1, 0]
	// that currently have more than one transaction.
	Buckets []int `json:"buckets,omitempty"`

	// any of the numbers present in buckets: Integer - Number of transaction in the bucket.
	// A transaction is assigned to the highest bucket of a value lower than its fee per byte value.
	Bucket0    int `json:"0,omitempty"`
	Bucket1    int `json:"1,omitempty"`
	Bucket2    int `json:"2,omitempty"`
	Bucket5    int `json:"5,omitempty"`
	Bucket10   int `json:"10,omitempty"`
	Bucket20   int `json:"20,omitempty"`
	Bucket50   int `json:"50,omitempty"`
	Bucket100  int `json:"100,omitempty"`
	Bucket200  int `json:"200,omitempty"`
	Bucket500  int `json:"500,omitempty"`
	Bucket1000 int `json:"1000,omitempty"`
	Bucket2000 int `json:"2000,omitempty"`
	Bucket5000 int `json:"5000,omitempty"`
}

// Peer holds the details of a peer
type Peer struct {
	ID              string `json:"id,omitempty"`
	Address         string `json:"address,omitempty"`
	AddressState    int    `json:"addressState,omitempty"`
	ConnectionState int    `json:"connectionState,omitempty"`
	Version         int    `json:"version,omitempty"`
	TimeOffset      int    `json:"timeOffset,omitempty"`
	HeadHash        string `json:"headHash,omitempty"`
	Latency         int    `json:"latency,omitempty"`
	RX              int    `json:"rx,omitempty"`
	TX              int    `json:"tx,omitempty"`
}

// Transaction holds the details on a transaction
type Transaction struct {

	// hash: String - Hex-encoded hash of the transaction.
	Hash string `json:"hash,omitempty"`

	// blockHash: String (optional) - Hex-encoded hash of the block containing the transaction.
	BlockHash string `json:"blockHash,omitempty"`

	// blockNumber: Integer (optional) - Height of the block containing the transaction.
	BlockNumber int `json:"blockNumber,omitempty"`

	// timestamp: Integer (optional) - UNIX timestamp of the block containing the transaction.
	Timestamp int `json:"timestamp,omitempty"`

	// confirmations: Integer (optional, default: 0) - Number of confirmations of the block containing the transaction.
	Confirmations int `json:"confirmations,omitempty"`

	// transactionIndex: Integer (optional) - Index of the transaction in the block.
	TransactionIndex int `json:"transactionIndex,omitempty"`

	// from: String - Hex-encoded address of the sending account.
	From string `json:"from,omitempty"`

	// fromAddress: String - Nimiq user friendly address (NQ-address) of the sending account.
	FromAddress string `json:"FromAddress,omitempty"`

	// to: String - Hex-encoded address of the recipient account.
	To string `json:"to,omitempty"`

	// toAddress: String - Nimiq user friendly address (NQ-address) of the recipient account.
	ToAddress string `json:"toAddress,omitempty"`

	// value: Integer - Integer of the value (in smallest unit) sent with this transaction.
	Value int `json:"value,omitempty"`

	// fee: Integer - Integer of the fee (in smallest unit) for this transaction.
	Fee int `json:"fee,omitempty"`

	// data: String - (optional, default: null) Hex-encoded contract parameters or a message.
	Data string `json:"data,omitempty"`

	// flags: Integer - Bit-encoded transaction flags.
	Flags int `json:"flags,omitempty"`
}

// TransactionReceipt holds the details on a transaction receipt.
type TransactionReceipt struct {

	// transactionHash : String - Hex-encoded hash of the transaction.
	TransactionHash string `json:"transactionHash,omitempty"`

	// transactionIndex: Integer - Integer of the transactions index position in the block.
	TransactionIndex int `json:"transactionIndex,omitempty"`

	// blockHash: String - Hex-encoded hash of the block where this transaction was in.
	BlockHash string `json:"blockHash,omitempty"`

	// blockNumber: Integer - Block number where this transaction was in.
	BlockNumber int `json:"blockNumber,omitempty"`

	// confirmations: Integer - Number of confirmations for this transaction (number of blocks on top of the block where this transaction was in).
	Confirmations int `json:"confirmations,omitempty"`

	// timestamp: Integer - Timestamp of the block where this transaction was in.
	Timestamp int `json:"timestamp,omitempty"`
}

// OutgoingTransaction holds the details on a transaction that is not yet sent.
type OutgoingTransaction struct {

	// from: Address - The address the transaction is send from.
	From string `json:"from"`

	// fromType: Integer - (optional, default: 0, Account.Type.BASIC) The account type at the given address (BASIC: 0, VESTING: 1, HTLC: 2).
	FromType int `json:"fromType,omitempty"`

	// to: Address - The address the transaction is directed to.
	To string `json:"to"`

	// toType: Integer - (optional, default: 0, Account.Type.BASIC) The account type at the given address (BASIC: 0, VESTING: 1, HTLC: 2).
	ToType int `json:"toType,omitempty"`

	// value: Integer - Integer of the value (in smallest unit) sent with this transaction.
	Value int `json:"value"`

	// fee: Integer - Integer of the fee (in smallest unit) for this transaction.
	Fee int `json:"fee"`

	// data: String - (optional, default: null) Hex-encoded contract parameters or a message.
	Data string `json:"data,omitempty"`
}

// SyncStatus holds information about the sync status.
type SyncStatus struct {

	// startingBlock: Integer - The block at which the import started (will only be reset, after the sync reached his head)
	StartingBlock int `json:"startingBlock,omitempty"`

	// currentBlock: Integer - The current block, same as blockNumber
	CurrentBlock int `json:"currentBlock,omitempty"`

	// highestBlock: Integer - The estimated highest block
	HighestBlock int `json:"highestBlock,omitempty"`
}

// Wallet holds the details on a wallet.
type Wallet struct {

	// id: String - Hex-encoded 20 byte address.
	ID string `json:"id,omitempty"`

	// address: String - User friendly address (NQ-address).
	Address string `json:"address,omitempty"`

	// publicKey: String - Hex-encoded 32 byte Ed25519 public key.
	PublicKey string `json:"publicKey,omitempty"`

	// privateKey: String (optional) - Hex-encoded 32 byte Ed25519 private key.
	PrivateKey string `json:"privateKey,omitempty"`
}

// Work holds the instructions to mine the next block
type Work struct {

	// data: String - Hex-encoded block header. This is what should be passed through the hash function. The last 4 bytes describe the nonce, the 4 bytes before are the current timestamp. Most implementations allow the miner to arbitrarily choose the nonce and to update the timestamp without requesting new work instructions.
	Data string `json:"data,omitempty"`

	// suffix: String - Hex-encoded block without the header. When passing a mining result to submitBlock, append the suffix to the data string with selected nonce.
	Suffix string `json:"suffix,omitempty"`

	// target: Integer - Compact form of the hash target to submit a block to this client.
	Target int `json:"target,omitempty"`

	// algorithm String - Field to describe the algorithm used to mine the block. Always nimiq-argon2 for now.
	Algorithm string `json:"algorithm,omitempty"`
}

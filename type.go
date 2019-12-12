package nimiqrpc

// Account holds the details on an account
type Account struct {

	// ID - Hex-encoded 20 byte address
	ID string `json:"id,omitempty"`

	// address: String - User friendly address (NQ-address).
	Address string `json:"address,omitempty"`

	// balance: Integer - Balance of the account (in smallest unit).
	Balance int `json:"balance,omitempty"`

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
	Transactions []interface{} `json:"transactions,omitempty"`
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

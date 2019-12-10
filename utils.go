package nimiqrpc

import "encoding/hex"

// AddressToHex converts a Nimiq address to a 20 byte hex encoded address
func AddressToHex(address string) string {
	return hex.EncodeToString([]byte(address))
}

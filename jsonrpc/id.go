package jsonrpc

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

var hashSum = md5.New()

func NewID() string {
	return hex.EncodeToString(hashSum.Sum([]byte(time.Now().String())))
}

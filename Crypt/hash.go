package Crypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func HashMD5(data string) string {
	/*
		The data sent as parameter will be hashed in MD5 algorithm and sent back to caller
	*/

	// Creating a md5 io stream
	h := md5.New()
	// Writing the data to be hashed in the io channel
	_, _ = io.WriteString(h, data)
	// Hex Encode the hashed data
	hashData := hex.EncodeToString([]byte(h.Sum(nil)))
	return hashData
}

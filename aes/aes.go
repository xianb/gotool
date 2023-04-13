package aes

import (
	"encoding/hex"
	"github.com/base-crypto/aescrypto"
)

func Encrypt(msg, prikey string) string {
	encrypted := aescrypto.AesEncryptCBC([]byte(msg), []byte(prikey))
	return hex.EncodeToString(encrypted)
}

func Decrypt(msg, prikey string) string {
	encrypted, _ := hex.DecodeString(msg)
	decrypted := aescrypto.AesDecryptCBC(encrypted, []byte(prikey))
	return string(decrypted)
}

package encryptionutils

import (
	"crypto/sha256"
	"encoding/hex"
)

// 使用SHA-256加密
func EncryptWithSha256(original string) string {
	s := sha256.New()
	s.Write([]byte(original))
	e := hex.EncodeToString(s.Sum(nil))
	return e
}

package encryptionutils

import (
	"fmt"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	GenerateRSAKeyPair(2048)
}

func TestEncryptWithRSAPublicKeyAndEncodeWithBase64(t *testing.T) {
	encrypted := EncryptWithRSAPublicKeyAndEncodeWithBase64("Hello World!")
	fmt.Println("RSA公钥加密后:\n" + encrypted)
	decrypted := DecodeWithBase64AndDecryptWithRSAPrivateKey(encrypted)
	fmt.Println("RSA私钥钥解密后:\n" + decrypted)
	if decrypted != "Hello World!" {
		t.Errorf("Hello World!使用RSA加密再解密后数据不一致!")
	}
}

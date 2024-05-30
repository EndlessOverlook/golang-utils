package encryptionutils

import (
	"fmt"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	GenerateRSAKeyPair(2048)
}

func TestEncryptWithRSAPublicKeyAndEncodeWithBase64(t *testing.T) {
	originals := "Hello World!李凯伦$#%"
	fmt.Println("加密前原始内容:\n" + originals)
	encrypted := EncryptWithRSAPublicKeyAndEncodeWithBase64(originals)
	fmt.Println("RSA公钥加密后:\n" + encrypted)
	decrypted := DecodeWithBase64AndDecryptWithRSAPrivateKey(encrypted)
	fmt.Println("RSA私钥钥解密后:\n" + decrypted)
	if decrypted != originals {
		t.Errorf("Hello World!使用RSA加密再解密后数据不一致!")
	}
}

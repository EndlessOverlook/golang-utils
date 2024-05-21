package encryptionutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

// EncryptWithSha256 使用SHA-256加密
func EncryptWithSha256(original string) string {
	s := sha256.New()
	s.Write([]byte(original))
	e := hex.EncodeToString(s.Sum(nil))
	return e
}

var commonRSAPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA3ppFQBHKpC9LJFz8kEH/IfunujiFNTpU4u/2mwuNyD+ZtBvr
fQjl2a9A2zTeOPB/Ayv9hFSKRdNervkb8pf8xG9IDjVRYeSWCJDsPfZTa2sz9pgc
yYXl/5TDozFVStbn/4Qk3DIK+KbIVUvAlMmostSZWAQkUcVVGtRgNjdIoSw+IHb7
PyzmFf1SiriFdmG4/JzezgfO4lp2LGmo/rxpZjBVCbUjXR2CzSmx34y/TFwyWsY8
gTS69N9BwlP0diC8H9WhdL4u+AHIs1jLcCYpwIoeMFji8MPXZYHxtzhzgfLIlmGH
345aF/dfJEfc3FTmraBxC9/OfVysiwwz1Ym3WwIDAQABAoIBAATvbKuf+VsshUD4
Is1b+TkaqDxTtnhYo7soUNJdLjaLWnAI0m7CPL5XyVtSz755bod2eTdu3GXx6r58
XWdCgWMliHZwLi66b60e8W3qQS03gj5sXu8E3fWNqcy7ckWngBtFJcAQB/g9UmG5
vTpPQkM8vW2bIecCNgHTRZLOJPkVxHA/8A/nmX4LFvBPnOyRMlUpi9v0iLa1mTID
f0mOOJVA4hvd41xBw13Dfvnag1X/Lwg9/Wjkbxk7TI/oiNd0PgWIpeg0htjn1Zn4
3o+AJoGEI9Yp/+byieSNANBvgFImg8xgwQiZEXE60MOq9zNJ0N294SMYQaZ5p3V6
sxDKlPkCgYEA6Q/NSGMxnwRQx9hSPze1OuLwWWPA7gKaYEZcF2BKno4BoUi6WiG6
qXtw/xQF8J6Ddor7Bj7/4kQN9jotOVV7+DkeEYbyErR4TQWO8rhx61Q0djUlXQhP
BPKDp1d8Xx40tvxmJslnazB+PyCn26fyQL+0xIK0vbsk+2aqn8OQdjUCgYEA9ILx
LPwbcy3C67SXvdgckLaAeHcRumwd1XuDj6vr3ZzmwQMM5LEOsqQzi/jlkAr1RVa/
0jR9yiBGJLvu7HZBfO/v3U0JxEUtxsTJOWESJoU9R7q+4CANx9y5VkcDTVwE57AV
jwJmtLW+MQTn6Y0BMH2O2Dwa5rJ7tjpgbN7D6U8CgYANxtNZrhNm1Ph5d74U4pQh
8PxILG9hImJBwqb68sQ2byqFgCkT3t6L2jIpweM///7BaSHBksPJekBhB7tYokb7
u1SdsJ5qoKRj2SpECschJA8d1bJpt24EWtmRcMh3Tu/KyUy4ijtEyM8NqD6D63+U
eECVRSyt8nwTKFR5sTO+6QKBgGmXequgnAIbiwfbT4Rupbdg+Q4M7Z0CP8gHtQMA
jbWAoq5LsO+5/szq2E/kRs7zzIQLOH+Ap6n4UhJfDokTcKiRvZZmRId8EI68QOIE
p5/2/kXRwchsT0c2bNtzg1uPF5culcDSmqQDsITex4A2ooLcVinPmxZxkbdacnXR
pe1FAoGASO42C5f9qAFIdzGJou6UyXI9d1wvr9dsWGlxeBKV89Z//6UtOjYz+HgY
+czIE0MUFOekgw9usmfqDsvlPF4LMBCDVc3Csiy5zR3jIPBgwTzvlRvprnEF7XVu
X3qo9pHqXC+DvoAk1h/fJvGU34vJmNQeTqHuN6qjLwouU2oQbQ8=

-----END RSA PRIVATE KEY-----`

var commonRSAPublicKey = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA3ppFQBHKpC9LJFz8kEH/IfunujiFNTpU4u/2mwuNyD+ZtBvrfQjl
2a9A2zTeOPB/Ayv9hFSKRdNervkb8pf8xG9IDjVRYeSWCJDsPfZTa2sz9pgcyYXl
/5TDozFVStbn/4Qk3DIK+KbIVUvAlMmostSZWAQkUcVVGtRgNjdIoSw+IHb7Pyzm
Ff1SiriFdmG4/JzezgfO4lp2LGmo/rxpZjBVCbUjXR2CzSmx34y/TFwyWsY8gTS6
9N9BwlP0diC8H9WhdL4u+AHIs1jLcCYpwIoeMFji8MPXZYHxtzhzgfLIlmGH345a
F/dfJEfc3FTmraBxC9/OfVysiwwz1Ym3WwIDAQAB
-----END RSA PUBLIC KEY-----`

var rsaOAEPLabel = []byte("OAEP")

func GenerateRSAKeyPair(keyLength int) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		panic("Error occurred when generating rsa key pair!")
	}
	pkcs1derPrivateKey := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	pemPrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1derPrivateKey,
	})
	pkcs1derPublicKey := x509.MarshalPKCS1PublicKey(&rsaPrivateKey.PublicKey)
	pemPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pkcs1derPublicKey,
	})
	fmt.Println("私钥: ")
	fmt.Println(string(pemPrivateKey))
	fmt.Println("公钥: ")
	fmt.Println(string(pemPublicKey))
}

func EncryptWithRSAPublicKeyAndEncodeWithBase64(plainData string) string {
	rsaPublicKeyBlock, _ := pem.Decode([]byte(commonRSAPublicKey))
	if rsaPublicKeyBlock == nil {
		fmt.Println("Error occurred when parsing PEM block inside rsaPublicKey")
	}
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(rsaPublicKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error occurred when parsing rsaPublicKey", err)
	}
	plainDataBytes := []byte(plainData)
	rsaEncryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPublicKey, plainDataBytes, rsaOAEPLabel)
	if err != nil {
		fmt.Printf("Error occurred when encrypting plainData data [%s] with rsaPublicKey", plainData)
	}
	return base64.StdEncoding.EncodeToString(rsaEncryptedData)
}

func DecodeWithBase64AndDecryptWithRSAPrivateKey(base64EncodedRSAEncrpytedData string) string {
	privateKeyBlock, _ := pem.Decode([]byte(commonRSAPrivateKey))
	if privateKeyBlock == nil {
		fmt.Println("Error occurred when parsing PEM block inside rsaPrivateKey")
	}
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Error occurred when parsing rsaPrivateKey", err)
	}
	rsaEncryptedData, err := base64.StdEncoding.DecodeString(base64EncodedRSAEncrpytedData)
	if err != nil {
		fmt.Println("Error occurred when decoding base64EncodedRSAEncryptedData with Base64", err)
	}
	plainData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, rsaPrivateKey, rsaEncryptedData, rsaOAEPLabel)
	if err != nil {
		fmt.Printf("Error occurred when decrypting rsa encrypted data [%s] with rsaPrivateKey", rsaEncryptedData)
	}
	return string(plainData)
}

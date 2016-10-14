package utils

import (
	"encoding/pem"
	"crypto/x509"
	"os"
	"crypto/rsa"
	"crypto/rand"
)

func GenRSAKey(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type: "RSA_PRIVATE_KEY",
		Bytes: derStream,
	}

	file, err := os.Create("private.pem");

	if err != nil {
		return err;
	}

	err = pem.Encode(file, block)

	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return privateKey

}

func PrivateDecrypt() {

}

func PublicEncrypt() {

}
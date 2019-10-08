package hsrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

/**
  RSA Demo[配置你可以提出来]
  @author Bill
*/

var decrypted string

func init() {
}

func usage() {
	var data []byte
	var err error
	if decrypted != "" {
		data, err = base64.StdEncoding.DecodeString(decrypted)
		if err != nil {
			panic(err)
		}
	} else {
		data, err = RsaEncrypt([]byte("polaris@studygolang.com"))
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCiAf7dzSYjXJQ28cqslEWj0Xj1C6zm1ccRAsYDesR/6y/Lbi2v
IBdMLLin+jD+ogfdWN/4MUsY1wTf/WVRyQQdpZoUQgKSPBVES1tUSOEVVgEq8iAT
k+5yycO2A00MnfqzFU9URYrReFleTfCIqUoVjSIhjtF+7uN9dO3AblAViwIDAQAB
AoGAJgqr6y9J/lG15/T872tdsur0KS5VqlqWhPMOxjBkxdjW/0De71lsvxFdRuxj
5tbrW9mLtf71MF8FIQeibAL0uRW1OFSS8PzCIeAlily74vfai1diZSnSScd91v1x
daW8yrf9Jzn0fRTarN2Ivx2ELCu/vhAB+fdDXxoozmZdiXECQQDVG+EPX20bsJU+
ukrLbMzjwQkzzv+HT9SPgOYzmj7JCBFrRYIwOKqJgRSrOFY/QIYR5jLKgmEjnPeg
PCbc/U8ZAkEAwp01HUGnL5zRNqjQXCm2SxDiage9sSsGu+OmawfxGGV9iTxa34s6
H9b6Zf2a5VmIB7QkZy8pFdzEMAXZIM2yQwJBAJBP3MxoNp61qZtc1CGFgAoLQowO
9QNQkATNqRXwseu4xvL5kvvMo8+R1clM0TrsGt4sIOD6AYX4Wcfsc/TapKECQQC2
laN1VVkzVF78dxk0vnWmq6qga67Of12L/aLum77Ycq96h8OCIy9fG+REt9ARCRSU
la/pltdnCX6Ox43UNufdAkBb5Xb2HHiudJi2+HW347ohMu8j5io1xBxt1ICuVXFc
tifJrsCvDQN96GxnPr+/koMYa6w1CUhjog68dVrZUm8q
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCiAf7dzSYjXJQ28cqslEWj0Xj1
C6zm1ccRAsYDesR/6y/Lbi2vIBdMLLin+jD+ogfdWN/4MUsY1wTf/WVRyQQdpZoU
QgKSPBVES1tUSOEVVgEq8iATk+5yycO2A00MnfqzFU9URYrReFleTfCIqUoVjSIh
jtF+7uN9dO3AblAViwIDAQAB
-----END PUBLIC KEY-----
`)

// RsaEncrypt 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func HttpRawRsaDecrypt(ciphertext string) (string, bool) {
	dText, _ := base64.StdEncoding.DecodeString(ciphertext)
	origData, err := RsaDecrypt(dText)
	if err != nil {
		return "", false
	}
	return string(origData), true
}

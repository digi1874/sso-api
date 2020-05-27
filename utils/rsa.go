/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 17:23:02
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-12 17:24:48
 */

package utils

import (
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"encoding/base64"
	"encoding/json"
	"bytes"
	"errors"
)

var prvKey *rsa.PrivateKey

// RsaPubPemEnc rsa公KEY
var RsaPubPemEnc string

// 生成rsa key
func init() {
	bits := int(1024)
	var err error

	prvKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	pkix, err := x509.MarshalPKIXPublicKey(&prvKey.PublicKey)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	block := pem.Block{
		Type: "PUBLIC KEY",
		Bytes: pkix,
	}

	pubPem := pem.EncodeToMemory(&block)
	RsaPubPemEnc = base64.StdEncoding.EncodeToString(pubPem)
}

// RsaDecrypt 解密
func RsaDecrypt(s string) ([]byte, error) {
	var ciphertext []byte
	limitLen := 172
	textLen := len(s)
	if textLen % limitLen != 0 {
		return ciphertext, errors.New("加密数据非法")
	}
	for i, j := 0, limitLen; i < textLen; i, j = i + limitLen, j + limitLen {
		decode, err := base64.StdEncoding.DecodeString(s[i:j])
		if err != nil {
			return ciphertext, err
		}
		c, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, decode)
		if err != nil {
			if err.Error() == "crypto/rsa: decryption error" {
				return ciphertext, errors.New("页面已过期，请刷新重试")
			}
			return ciphertext, err
		}
		var buffer bytes.Buffer
		buffer.Write(ciphertext)
		buffer.Write(c)
		ciphertext = buffer.Bytes()
	}
	return ciphertext, nil
}

// RsaDecryptUnmarshal 解析加密json
func RsaDecryptUnmarshal(s string, resValue interface{}) error {
	b, err := RsaDecrypt(s)
	if err == nil {
		err = json.Unmarshal(b, &resValue)
	}
	return err
}

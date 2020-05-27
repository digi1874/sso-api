/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 17:26:32
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 14:00:14
 */

package utils

import (
	"time"
	"strconv"
	"math/rand"
	"crypto/sha256"
	"crypto/hmac"
	"encoding/json"
)

var jwtHeader = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
var jwtSecret = []byte(strconv.FormatInt(rand.Int63n(1572120575) + 604661760, 36))

// JWTEncode 自定义secret生成jwt
func JWTEncode(claims map[string]interface{}, secret []byte) (string, string, error) {
	bClaims, err := json.Marshal(claims)
	if err != nil {
		return "", "", err
	}
	payload := Base64Url(bClaims)
	return payload, JWTSignature(payload, secret), err
}

// JWTEncodeDefault 生成jwt
func JWTEncodeDefault(claims map[string]interface{}) (string, string, error) {
	return JWTEncode(claims, jwtSecret)
}

// JWTVerify 验证jwt。非完整验证，只验证签名和过期时间，如果后面需要其它验证再添加
func JWTVerify(payload string, signature string) bool {
	if JWTSignature(payload, jwtSecret) != signature {
		return false
	}

	var claims map[string]interface{}
	err := json.Unmarshal(Base64UrlDecode(payload), &claims)
	if err != nil {
		return false
	}

	exp := claims["exp"]
	if exp != nil && exp.(float64) < float64(time.Now().Unix()) {
		return false
	}

	return true
}

// JWTSignature jwt签名
func JWTSignature(payload string, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(jwtHeader + "." + payload))
	return Base64Url(mac.Sum(nil))
}

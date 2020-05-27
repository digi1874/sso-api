/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 17:36:36
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 15:37:35
 */

package utils

import (
	"regexp"
	"strings"
	"encoding/base64"
)

var equalSignRE = regexp.MustCompile(`=+$`)
var plusSignRE  = regexp.MustCompile(`\+`)
var slashRE     = regexp.MustCompile(`\/`)

var minusSignRE  = regexp.MustCompile(`\-`)
var underlineRE     = regexp.MustCompile(`\_`)

// Base64Url []byte转Base64Url
func Base64Url(b []byte) string {
	encoded := base64.StdEncoding.EncodeToString(b)
	encoded = equalSignRE.ReplaceAllString(encoded, "")
	encoded = plusSignRE.ReplaceAllString(encoded, "-")
	encoded = slashRE.ReplaceAllString(encoded, "_")
	return encoded
}

// Base64UrlDecode Base64Url解[]byte
func Base64UrlDecode(encoded string) []byte {
	pad := len(encoded) % 4
	if pad > 1 {
		var ar [4]string
		encoded += strings.Join(ar[:pad+1], "=")
	}
	encoded = minusSignRE.ReplaceAllString(encoded, "+")
	encoded = underlineRE.ReplaceAllString(encoded, "/")
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return decoded
}

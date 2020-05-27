/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 21:06:17
 * @Last Modified by:   lin.zhenhui
 * @Last Modified time: 2020-03-12 21:06:17
 */

package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5Password 加密密码
func MD5Password(s1 string, s2 string) string {
	has := md5.Sum([]byte(s1 + "CicNxGuyP6QX7V1D" + s2))
	return fmt.Sprintf("%x", has)
}

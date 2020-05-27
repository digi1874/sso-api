/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 17:25:05
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-12 21:43:32
 */

package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"sso-api/utils"
)

// GetRsaHander 获取公KEY
func GetRsaHander(c *gin.Context) {
	now := time.Now()
	c.Header("Date", now.UTC().Format(http.TimeFormat))

	secret := strconv.FormatInt(now.Unix(), 10)
	payload, signature, err := utils.JWTEncode(gin.H{ "pub": utils.RsaPubPemEnc }, []byte(secret))

	if err != nil {
		cJSONBadRequest(c, "抱歉，出错了！")
		return
	}

	cJSONOk(c, payload + "." + signature)
}

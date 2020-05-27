/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-18 11:34:09
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-18 11:37:56
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	IP "sso-api/ip"
)

// IPHandle IP info
func IPHandle(c *gin.Context) {
	ip := IP.NewQQwry()
	cJSONOk(c, ip.Find(c.Param("ip")))
}


// GetIPInfoHandle IP info
func GetIPInfoHandle(c *gin.Context) {
	ip := IP.NewQQwry()
	cJSONOk(c, ip.Find(c.ClientIP()))
}

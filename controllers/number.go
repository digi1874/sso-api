/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-05 16:50:19
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-13 15:50:59
 */

package controllers

import (
	"regexp"

	"github.com/gin-gonic/gin"

	"sso-api/model"
)

// NumberRE 账号正则
var NumberRE = regexp.MustCompile(`^\w+$`)

// NumberExistHandle 检查账号是否存在
func NumberExistHandle(c *gin.Context) {
	number := c.Param("number")

	if number == "" {
		cJSONBadRequest(c, "账号不能为空")
		return
	}

	if NumberRE.MatchString(number) != true {
		cJSONBadRequest(c, "账号只能由英文、数字、下划线组成")
		return
	}

	var user model.User
	user.Filter.Number = number

	cJSONOk(c, user.NumberExist())
}

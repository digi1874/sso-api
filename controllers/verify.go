/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-14 12:19:36
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 16:29:07
 */

package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"

	"sso-api/model"
	"sso-api/utils"
)

// VerifyHandle 验证token
func VerifyHandle(c *gin.Context) {
	token := c.Param("token")

	errStr, _ := verifyToken(token)

	if errStr != "" {
		cJSONUnauthorized(c, errStr)
		return
	}

	cJSONOk(c, true)
}

func verifyToken(token string) (errStr string, userLogin model.UserLogin) {
	if token == "" {
		return "token不能为空", userLogin
	}

	tk := strings.Split(token, ".")
	if len(tk) != 2 {
		return "token不合法", userLogin
	}

	userLogin.Filter.Signature = tk[1]
	userLogin.Detail()

	if userLogin.ID == 0 {
		return "签名不存在", userLogin
	}

	if userLogin.State != 1 {
		return "token已登出", userLogin
	}

	if utils.JWTVerify(tk[0], tk[1]) == false {
		return "token无效", userLogin
	}

	return errStr, userLogin
}

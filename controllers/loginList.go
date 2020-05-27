/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-14 15:44:40
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-18 18:24:44
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"sso-api/model"
)

// LoginListHandle 登录列表
func LoginListHandle(c *gin.Context) {
	token := c.Param("token")

	errStr, userLogin := verifyToken(token)

	if errStr != "" {
		cJSONUnauthorized(c, errStr)
		return
	}

	var err error
	var userLoginList model.UserLoginList
	userLoginList.Page, userLoginList.Size, err = listHandle(c, &userLoginList.Filter)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	userLoginList.Filter.UserID = userLogin.UserID
	userLoginList.Find()

	cJSONOk(c, userLoginList)
}

/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-13 18:00:33
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-13 21:39:00
 */

package controllers

import (
	"time"

	"github.com/gin-gonic/gin"

	"sso-api/model"
	"sso-api/utils"
)

type passwordData struct {
	Data string `json:"data" binding:"required"`
}

type passwordFormData struct {
	Signature   string `json:"signature"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// PasswordUpdateHandle 更新密码
func PasswordUpdateHandle(c *gin.Context) {
	var formJSON passwordData
	// 验证接收数据
	err := c.ShouldBindJSON(&formJSON)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	var formData passwordFormData

	err = utils.RsaDecryptUnmarshal(formJSON.Data, &formData)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	if formData.Signature == "" {
		cJSONBadRequest(c, "签名不能为空")
		return
	}

	if formData.Password == "" {
		cJSONBadRequest(c, "旧密码不能为空")
		return
	}

	if formData.NewPassword == "" {
		cJSONBadRequest(c, "新密码不能为空")
		return
	}

	var userLogin model.UserLogin
	userLogin.Filter.Signature = formData.Signature
	userLogin.Detail()

	if userLogin.ID == 0 {
		cJSONBadRequest(c, "签名不存在")
		return
	}

	if userLogin.State != 1 {
		cJSONUnauthorized(c, "登录无效，请重新登录")
		return
	}

	if uint(time.Now().Unix()) > userLogin.Exp {
		cJSONUnauthorized(c, "登录过期，请重新登录")
		return
	}

	var user model.User
	user.Filter.ID = userLogin.UserID
	user.Detail()

	if user.ID == 0 {
		cJSONBadRequest(c, "获取个人信息出错")
		return
	}

	if user.State != 1 {
		cJSONBadRequest(c, "账号已禁用")
		return
	}

	if utils.MD5Password(formData.Password, user.Number) != user.Password {
		cJSONBadRequest(c, "旧密码错误")
		return
	}

	var u model.User
	u.Filter.ID = user.ID
	u.Filter.Password = utils.MD5Password(formData.NewPassword, user.Number)
	u.Update()

	cJSONOk(c, "密码修改成功")
}

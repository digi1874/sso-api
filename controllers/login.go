/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 20:15:05
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-18 12:58:31
 */


package controllers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	IP "sso-api/ip"
	"sso-api/model"
	"sso-api/utils"
)

type loginFormData struct {
	Data      string `json:"data" binding:"required"`
	Host      string `json:"host"`
}

// RegisterHandle 注册
func RegisterHandle(c *gin.Context) {
	user, formData, err := getLoginForm(c)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	if NumberRE.MatchString(user.Filter.Number) != true {
		cJSONBadRequest(c, "账号只能由英文、数字、下划线组成")
		return
	}

	if user.NumberExist() {
		cJSONBadRequest(c, "账号已被注册")
		return
	}

	if len(user.Filter.Password) < 6 {
		cJSONBadRequest(c, "密码至少6个字符")
	}

	user.Create()

	token, err := createToken(user.Filter.ID, c.ClientIP(), formData.Host)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	cJSONOk(c, token)
}

// LoginHandle 登录
func LoginHandle(c *gin.Context) {
	user, formData, err := getLoginForm(c)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	user.Detail()

	if user.ID == 0 {
		cJSONBadRequest(c, "账号或密码错误")
		go addErrorLoginLog(user.Filter.Number, c.ClientIP(), formData.Host, "密码错误")
		return
	}

	if user.State != 1 {
		cJSONBadRequest(c, "账号已禁用")
		return
	}

	token, err := createToken(user.ID, c.ClientIP(), formData.Host)

	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	cJSONOk(c, token)
}

// LogoutHandle 退出
func LogoutHandle(c *gin.Context) {
	signature := c.Param("signature")

	if signature == "" {
		cJSONUnauthorized(c, "signature 不能为空")
		return
	}

	var ul model.UserLogin
	ul.Filter.Signature = signature
	ul.Detail()

	if ul.ID != 0 {
		ul.Filter.ID = ul.ID
		ul.Filter.State = 2
		ul.Update()
	}
	cJSONOk(c, "退出成功")
}


// 获取登录或注册表单的解析数据
func getLoginForm(c *gin.Context) (model.User, loginFormData, error) {
	var user model.User
	var formJSON loginFormData
	// 验证接收数据
	err := c.ShouldBindJSON(&formJSON)
	if err != nil {
		return user, formJSON, err
	}

	err = utils.RsaDecryptUnmarshal(formJSON.Data, &user.Filter)

	if err != nil {
		return user, formJSON, err
	}

	if user.Filter.Number == "" {
		return user, formJSON, errors.New("账号不能为空")
	}

	if user.Filter.Password == "" {
		return user, formJSON, errors.New("密码不能为空")
	}

	user.Filter.Password = utils.MD5Password(user.Filter.Password, user.Filter.Number)

	return user, formJSON, err
}

// 添加登录记录
func addLoginLog(userID uint, signature string, exp uint, ip string, host string) {
	qqWry  := IP.NewQQwry()
	ipInfo := qqWry.Find(ip)

	var userLogin model.UserLogin
	userLogin.Filter.UserID    = userID
	userLogin.Filter.Signature = signature
	userLogin.Filter.Exp       = exp
	userLogin.Filter.IP        = ip
	userLogin.Filter.Country   = ipInfo.Country + " " + ipInfo.Area
	userLogin.Filter.WebsiteID = getWebsiteID(host)
	userLogin.Create()
}

func addErrorLoginLog(number string, ip string, host string, errMsg string) {
	var user model.User
	user.Filter.Number = number
	user.Detail()
	if user.ID == 0 {
		return
	}

	qqWry  := IP.NewQQwry()
	ipInfo := qqWry.Find(ip)

	var userLogin model.UserLogin
	userLogin.Filter.UserID    = user.ID
	userLogin.Filter.IP        = ip
	userLogin.Filter.Country   = ipInfo.Country + " " + ipInfo.Area
	userLogin.Filter.WebsiteID = getWebsiteID(host)
	userLogin.Filter.Message   = errMsg
	userLogin.Filter.State     = 3
	userLogin.Create()
}

func createToken(userID uint, ip string, host string) (string, error) {
	exp := getExp()

	payload, signature, err := utils.JWTEncodeDefault(gin.H{
		"id"  : userID,
		"ip"  : ip,
		"exp" : exp,
	})

	if err == nil {
		addLoginLog(userID, signature, exp, ip, host)
	}

	return payload + "." + signature, err
}

func getExp() uint {
	return uint(time.Now().Add(time.Hour * 24 * 7).Unix())
}

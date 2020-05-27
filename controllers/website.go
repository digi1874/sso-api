/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-13 14:26:05
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 18:06:34
 */

package controllers

import (
	"github.com/gin-gonic/gin"

	"sso-api/model"
)

// GetWebsiteIDHandle 获取网站id
func GetWebsiteIDHandle(c *gin.Context) {
	host := c.Param("host")
	cJSONOk(c, getWebsiteID(host))
}

func getWebsiteID(host string) uint {
	var website model.Website
	website.Filter.Host = host
	website.FirstOrCreate()
	return website.Filter.ID
}

// WebsiteListHandle 获取网站列表
func WebsiteListHandle(c *gin.Context) {
	var err error
	var websiteList model.WebsiteList
	websiteList.Page, websiteList.Size, err = listHandle(c, &websiteList.Filter)
	if err != nil {
		cJSONBadRequest(c, err.Error())
		return
	}

	websiteList.Find()

	cJSONOk(c, websiteList)
}

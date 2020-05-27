/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 15:50:01
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-04 20:31:16
 */

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sso-api/controllers"
	"sso-api/process"
)

// Run run Router
func Run() {
	if process.IsDev == false {
		gin.SetMode(gin.ReleaseMode)
	}

	var Router = gin.Default()
	Router.Use(middleware())

	login := Router.Group("/login")

	login.GET("/number/:number/exist", controllers.NumberExistHandle)
	login.GET("/website/:host/id", controllers.GetWebsiteIDHandle)
	login.GET("/website", controllers.WebsiteListHandle)
	login.GET("/rsa", controllers.GetRsaHander)
	login.POST("/register", controllers.RegisterHandle)

	login.POST("/", controllers.LoginHandle)
	login.POST("/password", controllers.PasswordUpdateHandle)
	login.GET("/verify/:token", controllers.VerifyHandle)
	login.GET("/list/:token", controllers.LoginListHandle)
	login.POST("/logout/:signature", controllers.LogoutHandle)

	login.GET("/ip/info", controllers.GetIPInfoHandle)

	if process.IsDev == true {
		Router.GET("/ip/:ip", controllers.IPHandle)

		Router.Run("127.0.0.1:8021")
	} else {
		Router.Run("127.0.0.1:8020")
	}
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Auth")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Set("content-type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

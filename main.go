/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 15:43:52
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-13 17:09:45
 */

package main

import (
	"runtime"

	"sso-api/database"
	"sso-api/ip"
	"sso-api/routers"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ip.Init()
	defer database.DB.Close()
	routers.Run()
}

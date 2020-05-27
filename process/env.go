/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 12:05:56
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-12 15:48:27
 */

package process

import (
	"flag"
)

// IsDev 是否开发环境。通过运行参数env判断。开发环境：```go run main.go -env=dev```
var IsDev bool

func init() {
	var env string
	flag.StringVar(&env, "env", "", "")
	flag.Parse()
	IsDev = env == "dev"
}

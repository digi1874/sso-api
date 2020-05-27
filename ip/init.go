package ip

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// Init 初始化
func Init() {
	datFile := flag.String("qqwry", "./qqwry.dat", "纯真 IP 库的地址")
	flag.Parse()

	IPData.FilePath = *datFile

	startTime := time.Now().UnixNano()

	res := IPData.InitIPData()

	if v, ok := res.(error); ok {
		log.Panic(v)
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("IP 库加载完成 共加载:%d 条 IP 记录, 所花时间:%.1f ms\n", IPData.IPNum, float64(endTime-startTime)/1000000)
}

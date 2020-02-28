package main

import (
	"fmt"
	"goStudy/日志库/mylogger"
	"time"
)

// 测试我自己写的日志库
func main() {
	// 测试打印到控制台
	// log := mylogger.NewLog("DEBUG")
	// id := 1001
	// name := "tianxiang"
	// for {
	// 	log.Debug("这是一条debug日志,%d:%s", id, name)
	// 	log.Trace("这是一条trace日志,%d:%s", id, name)
	// 	log.Info("这是一条info日志,%d:%s", id, name)
	// 	log.Warning("这是一条warning日志,%d:%s", id, name)
	// 	log.Error("这是一条error日志,%d:%s", id, name)
	// 	log.Fatal("这是一条fatal日志,%d:%s", id, name)
	// 	fmt.Println("***************************************")
	// 	time.Sleep(time.Second * 1)
	// }
	// 测试打印到文件中
	log := mylogger.NewFileLog("DEBUG", "./", "tianxiang", 1*1024*1024)
	id := 1001
	name := "tianxiang"
	for {
		log.Debug("这是一条debug日志,%d:%s", id, name)
		log.Trace("这是一条trace日志,%d:%s", id, name)
		log.Info("这是一条info日志,%d:%s", id, name)
		log.Warning("这是一条warning日志,%d:%s", id, name)
		log.Error("这是一条error日志,%d:%s", id, name)
		log.Fatal("这是一条fatal日志,%d:%s", id, name)
		fmt.Println("***************************************")
		time.Sleep(time.Second * 1)
	}
}

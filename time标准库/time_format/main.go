package main

import (
	"fmt"
	"time"
)

// 格式化时间
func main() {
	// time := time.Now()
	// fmt.Println(time)
	// fmt.Println(time.Format("2006-01-02 15:04:05"))
	// fmt.Println(time.Format("2006-01-02"))
	// fmt.Println(time.Format("15:04:05"))
	// fmt.Println(time.Format("2006-01-02 03:04:05PM"))
	// fmt.Println(time.Format("2006-01-02 15:04:05.000"))
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2019-01-01")
	if err != nil {
		fmt.Println("传入的时间格式不对！", err)
		return
	}
	fmt.Println(timeObj.Unix())
	// Sleep
	n := 5
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("觉醒了！")
	// time.Sleep(5)
}

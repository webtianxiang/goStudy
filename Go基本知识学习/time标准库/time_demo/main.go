package main

import (
	"fmt"
	"time"
)

// 时间类型

func main() {
	now := time.Now() // 获取当前时间对象
	fmt.Printf("%T", now)
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	res := time.Unix(1582535171, 0)
	fmt.Println(res)
	fmt.Println(res.Year())
	fmt.Println(res.Month())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 24h
	fmt.Println(now.Add(24 * time.Hour))
	// 定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Printf("%v\n", t) // 1s执行一次
	}
}

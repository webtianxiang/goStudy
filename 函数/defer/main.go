package main

import (
	"fmt"
)

// defer语句
// Go语言中，defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时
// 按照defer定义语句的顺序，逆序执行。
// 也就是说，先被defer定义的语句最后被执行，最后defer定义的语句最先被执行
func main() {
	// 在一个命名函数中，无法再声明一个命名函数
	deferDemo()
	testMain()
}
func deferDemo() {
	// defer常用于释放资源
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿01") // 将后面的语句延迟执行，在函数即将执行完毕时再执行defer定义的语句
	defer fmt.Println("嘿嘿嘿02") // 一个函数中可以存在多个defer语句
	defer fmt.Println("嘿嘿嘿03") // 多个defer语句按照先声明后执行（后声明先执行）的顺序执行
	fmt.Println("end")
}

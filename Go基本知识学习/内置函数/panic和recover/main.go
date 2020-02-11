package main

import "fmt"

// panic和recover：用来做错误处理
// panic可以在任何地方引发，但recover只能在defer调用的函数中有效
// defer一定要在可能发生panic的语句之前定义
func main() {
	f1()
	f2()
	f3()
}
func f1() {
	fmt.Println("panic in f1")
}
func f2() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("recover in f2")
		}
		fmt.Println("释放资源")
	}()
	panic("出现了一个严重的错误") // 程序崩溃退出
	fmt.Println("panic in f2")
}
func f3() {
	fmt.Println("panic in f3")
}

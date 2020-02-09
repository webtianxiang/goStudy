package main

import "fmt"

func main() {
	// var a *int // nil pointer
	// *a = 100
	// new函数申请一个内存地址
	var a = new(int)
	*a = 100
	fmt.Println(&a, *a)
}
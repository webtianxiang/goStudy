package main

import "fmt"

// 递归recursion：函数自己调用自己
// Go语言支持递归，但是开发者应在使用时设置递归结束条件，否则将陷入无限循环中
func f1() {
	f1()
}
func factorial(n uint64) (res uint64) {
	if n > 1 {
		res = n * factorial(n-1)
		return
	}
	return 1
}

// n个台阶，一次可以走一个，一次也可以走两个，问有多少个走法
func upSteps(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return upSteps(n-1) + upSteps(n-2)
	}
}

func main() {
	fmt.Println("阶乘：", factorial(5))
	fmt.Println("上台阶：", upSteps(5))
}

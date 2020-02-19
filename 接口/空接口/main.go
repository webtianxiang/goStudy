package main

import "fmt"

// 空接口：通常定义成下面的格式
// 特点：所有的类型都实现了空接口，也就是任意类型的变量都能保存到空接口类型当中
type nullInter interface{}

func show(x interface{}) {
	fmt.Printf("v: %v  t:%T\n", x, x)
}
func main() {
	// interface{}空接口类型
	var m1 map[int64]interface{}
	m1 = make(map[int64]interface{}, 10)
	m1[0] = "qwe"
	m1[1] = 123
	m1[2] = false
	m1[3] = [...]int{1, 2, 3}
	fmt.Println(m1)
	// 空接口的参数
	// 1.作为函数的参数
	show(true)
	show(nil)
	show([]string{"1", "2", "3"})
}

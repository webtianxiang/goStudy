package main

import "fmt"

// 结构体的匿名字段：没有名字的字段
// 用于变量比较少的场景
// 原本
type person struct {
	name string
	age  int
}

// 匿名
type person1 struct {
	string
	int
}

func main() {
	p1 := person1{"周林", 21}
	// 用变量类型取值
	fmt.Println(p1.string)
}

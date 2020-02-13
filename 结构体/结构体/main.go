package main

import "fmt"

// 结构体
// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。
// 结构体的定义
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p1 person
	p1.name = "tianxiang"
	p1.age = 23
	p1.gender = "男"
	p1.hobby = []string{"篮球", "足球", "羽毛球"}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	// 匿名结构体，多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "呵呵呵"
	s.y = 100
	fmt.Printf("type:%T val:%v\n", s, s)
}

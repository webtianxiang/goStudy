package main

import "fmt"

// 构造函数：返回一个结构体变量的函数
type person struct {
	name string
	age  int
}

func newPerson1(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

// 当结构比较大时，尽量返回结构体指针，来减少系统开销
func newPerson2(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
func main() {
	p1 := newPerson1("田翔", 23)
	p2 := newPerson1("王紫冰", 23)
	fmt.Println(p1, p2)
	p3 := newPerson2("田翔", 23)
	p4 := newPerson2("王紫冰", 23)
	fmt.Println(p3, p4)
	// 结构体是值类型，赋值的时候是拷贝
}

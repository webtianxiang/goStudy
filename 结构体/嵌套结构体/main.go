package main

import "fmt"

// 嵌套结构体
type address struct {
	proviences string
	city       string
}
type person struct {
	name string
	age  int
	addr address
}
type workPlace struct {
	proviences string
}

// 匿名嵌套结构体
type person1 struct {
	name string
	age  int
	address
	workPlace
}
type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周丽",
		age:  22,
		addr: address{
			proviences: "河北省",
			city:       "唐山市",
		},
	}
	p2 := person1{
		name: "周丽1",
		age:  23,
		address: address{
			proviences: "河北省1",
			city:       "唐山市1",
		},
		workPlace: workPlace{
			proviences: "河北省2",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)
	// fmt.Println(p1.city) // 现在无法访问
	// 查找顺序：现在自己的结构体里面找该字段，找不到的话再找匿名结构体中的该字段
	fmt.Println(p2.city) // 现在可以访问，把匿名嵌套结构体当做自己的变量。
	// 如果出现匿名结构体冲突，则需要写全
	fmt.Println(p2.workPlace.proviences)
}

package main

import "fmt"

type animals struct {
	name string
}

// 给animals实现动的办法
func (a animals) move() {
	fmt.Printf("谁会动？我是：%s，我会动\n", a.name)
}

type dog struct {
	feet    uint
	animals // animals拥有的方法，dog也拥有了。变向实现继承
}

func (d dog) wang() {
	fmt.Printf("谁会叫？我是：%s，汪汪汪~", d.name)
}

func main() {
	d1 := dog{
		feet:    2,
		animals: animals{name: "狗"},
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}

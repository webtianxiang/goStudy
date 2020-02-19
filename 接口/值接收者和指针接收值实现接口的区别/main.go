package main

import "fmt"

type animal interface {
	move()
	eat(something string)
}
type chicken struct {
	name string
	feet int8
}
type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("我是一只猫，我走猫步~")
}

func (c cat) eat(str string) {
	fmt.Printf("我是一只猫，我吃%s~\n", str)
}
func main() {
	var a1 animal
	fmt.Printf("v: %v  t:%T\n", a1, a1)
	c1 := cat{
		name: "茗贤",
		feet: 4,
	}
	fmt.Printf("v: %v  t:%T\n", c1, c1)
	a1 = c1
	fmt.Printf("v: %v  t:%T\n", a1, a1)
	a1.move()
	a1.eat("鱼")
	c2 := &cat{
		name: "小翔",
		feet: 2,
	}
	a1 = c2
	fmt.Printf("v: %v  t:%T\n", a1, a1)
	// 使用值接收者和使用指针接收值实现接口的区别
	// 1.使用值接收者，指针和值类型的变量都可以作为接收者传进去
	// 2.使用指针接收者，只能传入指针类型的变量进去
}

package main

import "fmt"

// 接口和类型之间的关系
// 1.多个类型可以实现同一个接口
// 2.同一个结构体可以实现多个接口
// 3.接口可以嵌套
type animal interface {
	mover
	eaten
}
type mover interface {
	move()
}
type eaten interface {
	eat(something string)
}

// cat这个结构体实现了eaten和mover接口
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

}

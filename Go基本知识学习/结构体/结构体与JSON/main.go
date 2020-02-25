package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON

// 1.序列化：把go语言中的结构体转化成js（或者其他语言）能识别的JSON
// 2.反序列化：把js（或者其他语言）的JSON转化成go语言可以识别的结构体
type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周丽",
		Age:  30,
	}
	b, error := json.Marshal(p1)
	if error != nil {
		fmt.Printf("marshal failed, error: %v\n", error)
		return
	}
	fmt.Printf("%#v\n", string(b))
	str := `{"name": "田翔","age": 23}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Printf("%#v\n", p2)
}

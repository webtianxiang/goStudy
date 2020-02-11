package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i == 5 {
			goto breakTag
		}
	}
breakTag:
	fmt.Print("结束for循环")
}

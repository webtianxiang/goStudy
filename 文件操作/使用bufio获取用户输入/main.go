package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScanln() {
	var str string
	fmt.Print("请输入内容：")
	fmt.Scan(&str)
	fmt.Printf("您输入的内容是：%s\n", str)
}
func useBufio() {
	var str string
	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	fmt.Printf("您输入的内容是：%s\n", str)
}
func main() {
	// useScanln()
	useBufio()
}

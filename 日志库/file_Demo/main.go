package main

import (
	"fmt"
	"os"
)

// 1.文件对象的类型
// 2.获取文件对象的详细信息
func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("Open File Failed, Error:", err)
		return
	}
	fmt.Printf("%T\n", fileObj)
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("取文件信息失败。")
		return
	}
	fmt.Printf("文件的大小是：%dB\n", fileInfo.Size())
	fmt.Printf("文件的大小是：%s\n", fileInfo.Name())
}

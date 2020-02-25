package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入
func method1() {
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	// Write
	fileObj.Write([]byte("wo shi shui~~~\n"))
	// WriteString
	fileObj.WriteString("我是谁~~~\n")
	fileObj.Close()
}
func method2() {
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("Hello World!") // 写到缓存中
	wr.Flush()                     // 将缓存中的内容写入文件
}

// 方法三
func method3() {
	str := "Hello 田翔！"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
}
func main() {
	method1()
	method2()
	// method3()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 利用bufio包读取文件
func readFromFileByBufio() {
	// 打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来读文件内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, e := reader.ReadString('\n')
		if e == io.EOF {
			fmt.Println("读完了！")
			return
		}
		if e != nil {
			fmt.Println("read from file failed: ", e)
			return
		}
		fmt.Print(line)
	}
}

// ioutil包读取文件
func readFromFileByIoutil() {
	res, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	fmt.Println(string(res))
	fmt.Println("读完了！")
}
func main() {
	// readFromFileByBufio()
	readFromFileByIoutil()
}

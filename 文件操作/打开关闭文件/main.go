package main

import (
	"fmt"
	"io"
	"os"
)

// 打开关闭文件
func main() {
	// 打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	for {
		// var temp = make([]byte, 128)  // 指定读的长度
		var temp [128]byte
		n, e := fileObj.Read(temp[:])
		if e == io.EOF {
			fmt.Println("读完了！")
			return
		}
		if e != nil {
			fmt.Println("read from file failed: ", e)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(temp[:n]))
		if n == 0 {
			return
		}
	}
}

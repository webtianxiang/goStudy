package main

import (
	"fmt"
	"os"
)

// 在文件中插入内容
func fileOperation() {
	fileObj, err := os.OpenFile("./test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open file failed: ", err)
		return
	}
	defer fileObj.Close()
	fileObj.Seek(2, 0)
	fileObj.Write([]byte("w\nw"))
	// var result [1]byte
	// n, err := fileObj.Read(result[:])
	// if err != nil {
	// 	fmt.Println("Read from file failed: ", err)
	// 	return
	// }
	// fmt.Println(string(result[:n]))
}
func main() {
	fileOperation()
}

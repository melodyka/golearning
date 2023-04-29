package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty: pair <type:*os.File, value:"/9"文件描述符>
	tty, err := os.OpenFile("D:\\go first\\9\\1.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Open file error", err)
		return
	}

	//r: pair
	var r io.Reader
	//r: pair <type:*os.File, value:"/9"文件描述符>
	r = tty

	//w: pair<type: , value: >
	var w io.Writer
	//w: pair <type:*os.File, value:"/9"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("halo \n"))

}

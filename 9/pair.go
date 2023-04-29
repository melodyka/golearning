package main

import (
	"fmt"
)

func main() {
	var a int
	//pair<statictype:string, value:"aceld">
	// a = "aceld"
	a = 123
	//pair<type:string, value:"aceld">
	var allType interface{}
	allType = a

	str, st := allType.(string)
	fmt.Println(str)
	fmt.Println(st)
}

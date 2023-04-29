package main

import (
	"fmt"
)

// 声明变量四种方式
var gA int = 100

func main() {

	//声明全局变量前三种没问题，第四种只能在函数体内声明
	var a int //默认值0
	fmt.Println("a = ", a)
	fmt.Printf("a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("b = %T\n", b)

	//不声明类型直接赋值
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("c = %T\n", c)

	var d string = "abcd"
	fmt.Println("d = ", d)
	fmt.Printf("d = %T\n", d)

	//先初始化在赋值
	e := 100
	//fmt.Printf("e = %s, type of e = %T\n", e, e)
	fmt.Println("e = ", e)
	fmt.Printf("e = %T\n", e)

	fmt.Println("gA = ", gA)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)

	var (
		aa = 100
		bb = 200
	)
	fmt.Println("aa = ", aa, ", bb = ", bb)

}

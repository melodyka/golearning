package main

import (
	"fmt"
	_ "fmt"
)

// 匿名形参
func foo1(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

// 有名形参，可以写成（r1, r2 int)
func foo2(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	ret1, ret2 := foo1("haha", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo2("foo2", 222)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
}

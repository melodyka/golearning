package main

import "fmt"

func myFunc(arg interface{}) { //通用万能类型 空接口
	fmt.Println("myFunc is called")
	fmt.Println(arg)
	//interface{}区分底层数据类型，可以用类型断言arg.(string)
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
}

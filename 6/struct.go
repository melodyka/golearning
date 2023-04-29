package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//指针传递
	book.auth = "666"
}

func changeBook2(book *Book) {
	//指针传递
	book.auth = "777"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhuang3"

	fmt.Printf("%v\n", book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}

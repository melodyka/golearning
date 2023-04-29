package main

import (
	"fmt"
)

const (
	//const里通过iota，每行累加1，第一行默认0
	Beijing = iota //0
	Shanghai
	Shenzhen
)

func main() {
	//常量(只读)
	const length int = 10

	fmt.Println("length = ", length)
	fmt.Println("Beijing = ", Beijing)
	fmt.Println("Shanghai = ", Shanghai)
	fmt.Println("Shenzhen = ", Shenzhen)
}

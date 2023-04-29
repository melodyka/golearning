package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()

	fmt.Println("goroutine exit")
}

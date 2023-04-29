package main

import "fmt"

// func changeValue(p *int) {
// 	*p = 10
// }

func main() {
	// var a int = 1
	// changeValue(&a)
	// fmt.Println("a = ", a)
	s := []int{1, 2, 3}
	s1 := s[0:2]
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
	s2[0] = 1000
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
}

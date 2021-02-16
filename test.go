package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 3)
	fmt.Println(s)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
}

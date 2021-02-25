package main //entrypoint

import (
	"fmt"
)

// 2 return value
func foo2(x int) (int, int) { //two return value both int
	return x, x + 1
}

func main() {
	a, b := foo2(3)
	fmt.Print(a)
	fmt.Print(b)
}

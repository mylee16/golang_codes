package main

import (
  "fmt"
  "time"
)

/*
Raca condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, add_one and sub_one are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/


func add_one(num *int) {
	for z := 0; z < 50; z++ {
		*num = *num + 1
		fmt.Println("plus", *num)
	}
}

func sub_one(num *int) {
	for z := 0; z < 50; z++ {
		*num = *num - 1
		fmt.Println("minus", *num)
	}
}

func main() {
  i := 0

  go add_one(&i)
  go sub_one(&i)

  time.Sleep(time.Second)
  fmt.Println("end")

}
/*Race condition occurs when more than one thread is accessing the same variable
and the result depends on the order of the execution of threads i.e. interleaving. As interleaving is non determinstic
if the program depends on interleaving, the result will be non determinstic.
In this example both sayHellow() and increment() function are acessing the same variable counter and
are dependent on the interleaving so the result is non deterministic. If you run the program multiple times
you will see different results.*/

package main

import (
	"fmt"
)

func prod(v1 int, v2 int, c chan int) {
	c <- v1 * v2}

func main() {
	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)

	a := <- c
	b := <- c
	fmt.Println(a*b)
}
/*Race condition occurs when more than one thread is accessing the same variable
and the result depends on the order of the execution of threads i.e. interleaving. As interleaving is non determinstic
if the program depends on interleaving, the result will be non determinstic.
In this example both sayHellow() and increment() function are acessing the same variable counter and
are dependent on the interleaving so the result is non deterministic. If you run the program multiple times
you will see different results.*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0

func sayHellow() {
	fmt.Printf("Hello #%v\n", counter)
	defer wg.Done()
}
func increment() {
	counter++
	defer wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHellow()
		go increment()
	}
	wg.Wait()

}
/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, 
each of which is sorted by a different goroutine. Each partition should be of approximately equal size. 
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the 
array should print the subarray that it will sort. When sorting is complete, the main goroutine should 
print the entire sorted list.

*/

package main

import (
	"fmt"
)

func swap(sli []int, i1 int, i2 int) {
	if sli[i2] < sli[i1] {
		sli[i2], sli[i1] = sli[i1], sli[i2]
	}
}

func bubblesort(sli []int, c chan []int){
	// go through all first values
	for i1 := range sli {
		// go through all second values
		for i2 := range sli[i1:] {
			swap(sli, i1, i2+i1)
		}
	}
	c <- sli
}

func input(sli []int, err error) []int {
	if err != nil {
		return sli
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		sli = append(sli, d)
	}
	return input(sli, err)
}

func main() {
	// get inputs
	fmt.Printf("Please enter a list of more than 4 integers:")
	sli := input([]int{}, nil)

	// before sorting
	fmt.Println("Before sorting:", sli)

	// dividing into 4 slices
	slice_len := len(sli)
	cut1 := slice_len / 2
	cut0 := cut1 / 2
	cut2 := cut1 + cut0

	sli0 := sli[:cut0]
	sli1 := sli[cut0:cut1]
	sli2 := sli[cut1:cut2]
	sli3 := sli[cut2:]

	// building the channels
	c0 := make(chan []int)
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)

	// creating the goroutines
	go bubblesort(sli0, c0)
	go bubblesort(sli1, c1)
	go bubblesort(sli2, c2)
	go bubblesort(sli3, c3)

	// receiving data from the channels
	sliSorted0 := <- c0
	sliSorted1 := <- c1
	sliSorted2 := <- c2
	sliSorted3 := <- c3

	// Printing the results from 4 channels
	fmt.Println("Goroutine 1 sorted:", sliSorted0)
	fmt.Println("Goroutine 2 sorted:", sliSorted1)
	fmt.Println("Goroutine 3 sorted:", sliSorted2)
	fmt.Println("Goroutine 4 sorted:", sliSorted3)

	// Appending to combine into 1 list
	sli_combined := append(sliSorted0, sliSorted1...)
	sli_combined = append(sli_combined, sliSorted2...)
	sli_combined = append(sli_combined, sliSorted3...)

	// resorting
	c := make(chan []int)
	go bubblesort(sli_combined, c)
	sliSortedCombined := <- c

	// Output
	fmt.Println("After sorting:", sliSortedCombined)
}

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.
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

func bubblesort(sli []int) []int {
	// go through all first values
	for i1 := range sli {
		// go through all second values
		for i2 := range sli[i1:] {
			swap(sli, i1, i2+i1)
		}
	}
	return sli
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
	fmt.Printf("Please enter a list of integers:")
	sli := input([]int{}, nil)

	// sorting
	fmt.Println("Before sorting:", sli)
	sliSorted := bubblesort(sli)
	fmt.Println("After sorting:", sliSorted)
}

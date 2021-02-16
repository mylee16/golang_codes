package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int, 0, 3)
	i := 1
	var num string

	for i == 1 {
		fmt.Printf("Please an integer: ")
		fmt.Scan(&num)

		switch {
		case num == "X" || num == "x":
			i = 0
			{
				break
			}
		default:
			numInt, _ := strconv.Atoi(num)
			sli = append(sli, numInt)
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}

}

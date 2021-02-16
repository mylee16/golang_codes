/*
Write a program which reads information from a file and represents it in a slice of structs. Assume
that there is a text file which contains a series of names. Each line of the text file has a first
name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for
the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read
each line of the text file and create a struct which contains the first and last names found in the
file. Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading all lines
from the file, your program should iterate through your slice of structs and print the first and last
names found in each struct.
*/

// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// https://stackoverflow.com/questions/26159416/init-array-of-structs-in-go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var filename string
	fmt.Printf("Please enter file name: ")
	fmt.Scan(&filename)

	type Name struct {
		Fname string
		Lname string
	}

	nameStruct := []Name{}

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)

	var line string
	// saving to slice of struct
	for {
		line, err = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		s := strings.Split(line, " ")

		p1 := Name{Fname: s[0],
			Lname: s[1]}

		nameStruct = append(nameStruct, p1)

		if err != nil {
			break
		}
	}

	// iterate over nameStruct to print out first and last name
	for _, names := range nameStruct {
		fmt.Println("First name:", names.Fname)
		fmt.Println("Last Name:", names.Lname)
		fmt.Println("")
	}
}

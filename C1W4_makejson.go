/*
Write a program which prompts the user to first enter a name, and then enter
an address. Your program should create a map and add the name and address to
the map using the keys “name” and “address”, respectively. Your program should
use Marshal() to create a JSON object from the map, and then your program should
print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	fmt.Printf("Please enter your name: ")
	fmt.Scan(&name)

	var address string
	fmt.Printf("Please enter your address: ")
	fmt.Scan(&address)

	type Person struct {
		Name    string
		Address string
	}

	p1 := Person{Name: name,
		Address: address}

	// convert to byte array
	barr, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(barr))

}

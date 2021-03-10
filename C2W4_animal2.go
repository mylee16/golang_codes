/*
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create
a new animal of one of the three types, or the user can request information about an animal that he/she has
already created. Each animal has a unique name, defined by the user. Note that the user can define animals
of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table
contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your
program should accept one command at a time from the user, print out a response, and print out a new prompt
on a new line. Your program should continue in this loop forever. Every command from the user must be either
a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The
second string is an arbitrary string which will be the name of the new animal. The third string is the type
of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by
creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string
is the name of the animal. The third string is the name of the information requested about the animal, either
“eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface
should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat()
method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method
should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define
methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user
creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the
user issues a query command.

*/

package main

import (
	"fmt"
)

type animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{ food, locomotive, spoken string }
type Bird struct{ food, locomotive, spoken string }
type Snake struct{ food, locomotive, spoken string }

func (c Cow) Eat() string   { return c.food }
func (c Cow) Move() string  { return c.locomotive }
func (c Cow) Speak() string { return c.spoken }

func (b Bird) Eat() string   { return b.food }
func (b Bird) Move() string  { return b.locomotive }
func (b Bird) Speak() string { return b.spoken }

func (s Snake) Eat() string   { return s.food }
func (s Snake) Move() string  { return s.locomotive }
func (s Snake) Speak() string { return s.spoken }

func main() {
	fmt.Println("Please enter a new animal, example 'newanimal coco cow'")
	fmt.Println("Or a query, example 'query coco speak' ")

	animalMap := make(map[string]animal)
	animalMap["cow"] = Cow{"grass", "walk", "moo"}
	animalMap["bird"] = Bird{"worms", "fly", "peep"}
	animalMap["snake"] = Snake{"mice", "slither", "hsss"}

	for {
		fmt.Println(">")

		var command, requestedAni, requestType string
		fmt.Scan(&command, &requestedAni, &requestType)

		if command == "newanimal" {
			animalMap[requestedAni] = animalMap[requestType]
			fmt.Print("Created it!")

		} else if command == "query" {
			ani := animalMap[requestedAni]
			switch requestType {
			case "eat":
				fmt.Println(ani.Eat())
			case "move":
				fmt.Println(ani.Move())
			case "speak":
				fmt.Println(ani.Speak())
			}
		}
	}
}

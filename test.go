package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct {
	food, locomotion, noise string
}

//Bird type
type Bird struct {
	food, locomotion, noise string
}

//Snake type
type Snake struct {
	food, locomotion, noise string
}

func (cow Cow) Eat() {

	fmt.Println(cow.food)
}

func (cow Cow) Move() {

	fmt.Println(cow.locomotion)
}

func (cow Cow) Speak() {

	fmt.Println(cow.noise)
}

func (bird Bird) Eat() {

	fmt.Println(bird.food)
}

func (bird Bird) Move() {

	fmt.Println(bird.locomotion)
}

func (bird Bird) Speak() {

	fmt.Println(bird.noise)
}

func (snake Snake) Eat() {

	fmt.Println(snake.food)
}

func (snake Snake) Move() {

	fmt.Println(snake.locomotion)
}

func (snake Snake) Speak() {

	fmt.Println(snake.noise)
}
func main() {

	var nameType map[string]string

	nameType = make(map[string]string)

	for {
		fmt.Print("Enter your request > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		a := scanner.Text()

		command := strings.Split(a, " ")[0]

		switch command {
		case "newanimal":

			name := strings.Split(a, " ")[1]
			animalType := strings.Split(a, " ")[2]

			_, ok := nameType[name]
			if ok == true {
				fmt.Printf("Animal name %s is taken\n", name)
			} else {
				nameType[name] = animalType
				fmt.Println("Created it")
			}
		case "query":
			name := strings.Split(a, " ")[1]
			info := strings.Split(a, " ")[2]

			switch nameType[name] {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				switch info {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				switch info {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				switch info {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Printf("%s is not in (eat, move, speak), please try again\n", info)
					continue
				}
			default:
				fmt.Printf("%s is not in (cow, bird, snake), please try again\n", name)
				continue
			}
		default:
			fmt.Printf("%s is not in (newanimal, query), please try again\n", command)
			continue

		}

	}
}

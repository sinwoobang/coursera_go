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

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animalMap := map[string]Animal{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		cmds := strings.Fields(scanner.Text())

		query := cmds[0]
		switch query {
		case "newanimal":
			name := strings.ToLower(cmds[1])
			species := strings.ToLower(cmds[2])

			switch species {
			case "cow":
				animalMap[name] = Cow{}
				fmt.Printf("%s is created.\n", name)
			case "bird":
				animalMap[name] = Bird{}
				fmt.Printf("%s is created.\n", name)
			case "snake":
				animalMap[name] = Snake{}
				fmt.Printf("%s is created.\n", name)
			}

		case "query":
			name := strings.ToLower(cmds[1])
			action := strings.ToLower(cmds[2])
			switch action {
			case "eat":
				animal, found := animalMap[name]
				if found == true {
					animal.Eat()
				}
			case "move":
				animal, found := animalMap[name]
				if found == true {
					animal.Move()
				}
			case "speak":
				animal, found := animalMap[name]
				if found == true {
					animal.Speak()
				}
			}
		}
	}
}

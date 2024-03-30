package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	animalMap := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		cmds := strings.Fields(scanner.Text())
		animal := cmds[0]
		action := strings.ToLower(cmds[1])

		switch action {
		case "eat":
			animalMap[animal].Eat()
		case "move":
			animalMap[animal].Move()
		case "speak":
			animalMap[animal].Speak()
		}
	}
}

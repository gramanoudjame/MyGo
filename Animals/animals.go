package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
}

func (c *Cow) Eat() {
	fmt.Println("grass")
}
func (c *Cow) Move() {
	fmt.Println("walk")
}
func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}
func (b *Bird) Move() {
	fmt.Println("fly")
}
func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}
func (s *Snake) Move() {
	fmt.Println("slither")
}
func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func createAnimal(animals map[string]Animal, name string, animalType string) {
	var a Animal
	switch animalType {
	case "cow":
		var c *Cow
		a = c
	case "bird":
		var b *Bird
		a = b
	case "snake":
		var s *Snake
		a = s
	default:
		return
	}
	animals[name] = a
	fmt.Println("Created it!")
}

func queryAnimal(animals map[string]Animal, name string, information string) {
	a, exists := animals[name]
	if exists {
		switch information {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			return
		}
	}
}

func main() {
	var animals = make(map[string]Animal)
	for { // loop forever
		fmt.Println(">")
		var command, arg1, arg2 string
		fmt.Scan(&command, &arg1, &arg2)
		switch command {
		case "newanimal":
			createAnimal(animals, arg1, arg2)
		case "query":
			queryAnimal(animals, arg1, arg2)
		}
	}
}

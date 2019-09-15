package main

import (
	"fmt"
)

type Animal struct{
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat() {
	fmt.Printf("%s\n",a.food)
}

func (a *Animal) Speak() {
	fmt.Printf("%s\n",a.noise)
}

func (a *Animal) Move() {
	fmt.Printf("%s\n",a.locomotion)
}



func main(){
	var a, v string;
	var animalType Animal;
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Printf("> ")
		fmt.Scan(&a)
		fmt.Scan(&v)
		switch a{
		case "cow":
			animalType = cow
		case "bird":
			animalType = bird
		case "snake":
			animalType = snake
		}

		switch v{
		case "eat":
			animalType.Eat()
		case "move":
			animalType.Move()
		case "noise":
			animalType.Speak()
		}
	}
}

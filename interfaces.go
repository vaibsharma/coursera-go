// Written by vaibhav

package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (cow *Cow) Eat() {
	fmt.Println("The cow eats grass")
}

func (cow *Cow) Move() {
	fmt.Println("The cow moves by walking")
}

func (cow *Cow) Speak() {
	fmt.Println("The cow says moo")
}

type Snake struct {
}

func (snake *Snake) Eat() {
	fmt.Println("The snake eats mice")
}

func (snake *Snake) Move() {
	fmt.Println("The snake moves by slithering")
}

func (snake *Snake) Speak() {
	fmt.Println("The snake says hsss")
}

type Bird struct {
}

func (bird *Bird) Eat() {
	fmt.Println("The bird eats worms")
}

func (bird *Bird) Move() {
	fmt.Println("The bird moves by flying")
}

func (bird *Bird) Speak() {
	fmt.Println("The bird says peep")
}

// Context struct
type Context struct {
	animals map[string]Animal
}

func NewContext() *Context {
	return &Context{
		animals: make(map[string]Animal),
	}
}

// Command functionality
type Command struct {
	action  string
	subject string
	param   string
}

func (command *Command) Scan() {
	fmt.Scanf("%s %s %s", &command.action, &command.subject, &command.param)
}

func (cmd *Command) HandleNew(ctx *Context) {

	var animal Animal

	switch cmd.param {
	case "cow":
		animal = &Cow{}
	case "bird":
		animal = &Bird{}
	case "snake":
		animal = &Snake{}
	default:
		fmt.Printf("Animal \"%s\" is not reconogized.\n", cmd.param)
		return
	}

	ctx.animals[cmd.subject] = animal
	fmt.Println("Created it!")
}

func (cmd *Command) HandleQuery(ctx *Context) {

	var exist bool
	var animal Animal

	animal, exist = ctx.animals[cmd.subject]
	if !exist {
		fmt.Printf("No animal named \"%s\" found.\n", cmd.subject)
		return
	}

	switch cmd.param {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("Action \"%s\" is not reconogized.\n", cmd.param)
		return
	}
}

func (cmd *Command) Handle(ctx *Context) {

	switch cmd.action {
	case "newanimal":
		cmd.HandleNew(ctx)
	case "query":
		cmd.HandleQuery(ctx)
	default:
		fmt.Printf("Command \"%s\" is not recognized.\n", cmd.action)
	}
}

func main() {

	cmd := new(Command)
	ctx := NewContext()

	for {
		fmt.Print(" > ")

		cmd.Scan()
		cmd.Handle(ctx)
	}
}

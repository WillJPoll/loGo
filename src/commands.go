package src

import "fmt"

type Command struct {
	name        string
	args        string
	subCommands []Command
}

func (c Command) ToString() {
	fmt.Println("********")
	fmt.Println("Name : ", c.name)
	fmt.Println("args : ", c.args)
	fmt.Println("Subcommands : ", c.subCommands)
	fmt.Println("********")

}

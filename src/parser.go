package src

import (
	"fmt"
	"regexp"
)

type Parser struct {
	Text   string
	Cursor int
}

func (p Parser) Has() bool {
	return len(p.Text) > p.Cursor
}

func (p *Parser) NewToken() string {
	var token string
	char := string(p.Text[p.Cursor])
	if char == " " {
		p.Cursor++
		return p.NewToken()
	}

	if char == "[" || char == "]" {
		p.Cursor++
		return char
	}

	for char != " " && p.Has() {
		token += char

		p.Cursor++
		if p.Cursor == len(p.Text) {
			break
		}

		fmt.Println("Cuurent char ln 38", char)
		char = string(p.Text[p.Cursor])
	}
	return token
}

func (p Parser) ExtractBlock() string {
	var token string
	var toParse string
	for p.Has() {
		token = p.NewToken()
		fmt.Println("CURRENT TOKEN : ", token)
		if token == "[" {
			continue
		} else if token == "]" {
			fmt.Println("end")
			break
		}
		fmt.Println("Constructing subcommand")

		toParse += " " + token
	}
	fmt.Println("New parser : ", toParse)
	return toParse
}

func (p Parser) Parse() []Command {
	movement, _ := regexp.Compile("^[fb]d|[lr]t$")
	pen, _ := regexp.Compile("^[pu]d")
	repeat, _ := regexp.Compile("^repeat$")

	var commands []Command
	for p.Has() {
		token := p.NewToken()

		if movement.MatchString(token) {
			commands = append(commands, Command{token, p.NewToken(), nil})
		} else if pen.MatchString(token) {
			commands = append(commands, Command{token, "", nil})

		} else if repeat.MatchString(token) {
			repeatCommand := Command{token, p.NewToken(), nil}
			subParser := Parser{p.ExtractBlock(), 0}
			repeatCommand.subCommands = subParser.Parse()
			commands = append(commands, repeatCommand)

		}
	}

	return commands
}

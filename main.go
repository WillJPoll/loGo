package main

import (
	"github.com/WillJPoll/loGo/src"
)

func main() {
	p := src.Parser{"repeat 3 [lt 100 rt 20]", 0}
	test := p.Parse()
	for command := range test {
		test[command].ToString()
	}
}

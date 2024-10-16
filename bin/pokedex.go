package main

import (
	"os"

	"github.com/yeldiRium/learning-go-pokedex/repl"
)

func main() {
	input := os.Stdin
	repl.StartRepl(input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("  help:  print this help text")
	fmt.Println("  exit:  exit the pokedex")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "exit":
			os.Exit(0)
		case "help":
			printHelp()
		default:
			fmt.Printf("'%s' is not a valid command.\n", input)
			fmt.Println()
			printHelp()
		}
	}
}

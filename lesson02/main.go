package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: lesson02 <name>")
		os.Exit(1)
	}

	fmt.Println("Hello,", args[0])
}

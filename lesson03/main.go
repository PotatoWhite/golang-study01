package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: lesson03 <name>")
		os.Exit(1)
	}

	printAny("Hello,", args)
}

func printAny(any ...interface{}) {
	fmt.Println(any...)
}

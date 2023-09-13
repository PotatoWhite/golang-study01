package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: lesson04 <name>")
		os.Exit(1)
	}

	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Wrong size")
		os.Exit(1)
	}

	wg := sync.WaitGroup{}

	for i := 0; i < size; i++ {
		wg.Add(1)
		go printAny(&wg, "Hello,", i)
	}

	wg.Wait()
}

func printAny(wg *sync.WaitGroup, any ...interface{}) {
	defer wg.Done()
	fmt.Println(any...)
}

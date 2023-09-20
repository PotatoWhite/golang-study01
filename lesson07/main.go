package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	mt  sync.Mutex
	msg string = "Hello World"
)

func main() {

	wg.Add(2)
	go updateMsg("Hello Potato")
	go updateMsg("Hello Tomato")
	wg.Wait()

	printMsg()
}

func updateMsg(m string) {
	mt.Lock()
	msg = m
	mt.Unlock()
	wg.Done()
}

func printMsg() {
	fmt.Println(msg)
}

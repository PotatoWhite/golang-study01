package main

import (
	"fmt"
	"sync"
)

var (
	usage uint64 = 0
	wg    sync.WaitGroup
	mt    sync.Mutex
)

type File struct {
	Name string
	Size uint64
}

func main() {

	fmt.Printf("Initial usage %d\n", usage)

	files := []File{
		{Name: "File1", Size: 10},
		{Name: "File2", Size: 20},
		{Name: "File3", Size: 30},
		{Name: "File4", Size: 40},
		{Name: "File5", Size: 50},
	}

	wg.Add(len(files))

	for _, file := range files {
		go addSize(file)
	}

	wg.Wait()
}

func addSize(file File) {
	defer wg.Done()

	mt.Lock()
	usage += file.Size
	fmt.Printf("%s(%d), current usage: %d\n", file.Name, file.Size, usage)
	mt.Unlock()
}

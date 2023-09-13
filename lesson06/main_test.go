package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printAny(t *testing.T) {
	stdOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printAny(&wg, "Test printAny")

	wg.Wait()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Test printAny") {
		t.Errorf("printAny() = %v, want %v", output, "Test printAny")
	}
}

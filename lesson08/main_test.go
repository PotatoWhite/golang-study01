package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	main()

	w.Close()

	result, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = stdout

	output := string(result)

	if !strings.Contains(output, "150") {
		t.Errorf("Unexpected output: %s", output)
	}
}

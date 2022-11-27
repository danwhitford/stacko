package main

import (
	"io"
	"os"
	"path"
	"testing"
)

func TestExamples(t *testing.T) {
	table := []struct{
		fname string
		expected string
	}{
		{
			"factorial.txt",
			"3628800",
		},
		{
			"fib-tail.txt",
			"55",
		},
		{
			"fib.txt",
			"55",
		},
	}

	for _, item := range table {
		path := path.Join("../..", "examples", item.fname)
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		input, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		run(t, string(input), item.expected)
	}
}
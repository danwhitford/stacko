package main

import (
	"io"
	"os"
	"path"
	"testing"
)

func TestExamples(t *testing.T) {
	table := []struct {
		fname    string
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
		// 		{
		// 			"fizzbuzz.txt",
		// 			`1
		// 2
		// fizz
		// 4
		// buzz
		// fizz
		// 7
		// 8
		// fizz
		// buzz
		// 11
		// fizz
		// 13
		// 14
		// fizzbuzz`,
		// 		},
	}

	for _, item := range table {
		t.Logf("Testing %s\n", item.fname)
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

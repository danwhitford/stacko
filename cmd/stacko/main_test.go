package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/danwhitford/stacko/vm"
	"github.com/google/go-cmp/cmp"
)

func run(t *testing.T, in string, out string) {
	var w bytes.Buffer
	runner := Runner{
		vm: vm.NewVM(&w),
	}
	err := runner.doLine(in)
	if err != nil {
		t.Error(err)
	}
	b, err := io.ReadAll(&w)
	if err != nil {
		t.Error(err)
	}
	diff := cmp.Diff(out, strings.TrimSpace(string(b)))
	if diff != "" {
		t.Errorf("failed input '%s'\nmismatch (-want +got):\n%s\n", in, diff)
	}
}

func TestMain(t *testing.T) {
	table := []struct {
		in, out string
	}{
		{
			in:  `"hello world" v`,
			out: "hello world",
		},
		{
			in: `( dup 5 = 'dup 'drop if ) 
							'foo 
							def 5 foo 6 foo v`,
			out: "5\n5",
		},
		{
			in:  `( dup 1 = ( dup * ) ( dup 1 - fact * ) if ) 'fact def 5 fact v`,
			out: "120",
		},
		{
			in: `(
				dup 2 <
				( )
				( dup 1 - fib-rec swap 2 - fib-rec + )
				if 
			) 'fib-rec def
			10 fib-rec
			v`,
			out: "55",
		},
		{
			in:  `() 'foo def foo v`,
			out: "",
		},
	}

	for _, test := range table {
		run(t, test.in, test.out)
	}
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SlowEcho(w io.Writer) {
	s := ""
	sep := ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func FastEcho(w io.Writer) {
	fmt.Fprintln(w, strings.Join(os.Args, " "))
}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(w io.Writer) {
	fmt.Fprintf(w, strings.Join(os.Args, " "))
}

func main() {
	Echo(os.Stdout)
}

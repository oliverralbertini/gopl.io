package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(w io.Writer) {
	for index, arg := range os.Args {
		fmt.Fprintln(w, index, arg)
	}
}

func main() {
	Echo(os.Stdout)

}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type fileCount struct {
	count int
	files string
}

func findDupes(w io.Writer) {
	counts := make(map[string]*fileCount)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fc := range counts {
		if fc.count > 1 {
			fc.files = strings.TrimSuffix(fc.files, ", ")
			fmt.Fprintf(w, "%s:\t%d\t%s\n", fc.files, fc.count, line)
		}
	}
}

func main() {
	findDupes(os.Stdout)
}

func countLines(f *os.File, counts map[string]*fileCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fc := counts[input.Text()]
		if fc == nil {
			fc = &fileCount{0, ""}
		}
		fc.count++
		if !strings.Contains(fc.files, f.Name()) {
			fc.files += f.Name() + ", "
		}
		counts[input.Text()] = fc
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

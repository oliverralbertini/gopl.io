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
	files []string
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
			fmt.Fprintf(w, "%s:\t%d\t%s\n", strings.Join(fc.files, ", "), fc.count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*fileCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &fileCount{0, make([]string, 0)}
		}
		fc := counts[input.Text()]
		fc.count++
		if !contains(fc.files, f.Name()) {
			fc.files = append(fc.files, f.Name())
		}
	}
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

func main() {
	findDupes(os.Stdout)
}

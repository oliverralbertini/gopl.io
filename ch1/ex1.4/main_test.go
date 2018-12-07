package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFindDupes(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "fixtures/12.txt", "fixtures/23.txt"}

	want := []string{
		fmt.Sprintf("fixtures/12.txt:\t2\tone"),
		fmt.Sprintf("fixtures/12.txt, fixtures/23.txt:\t2\ttwo"),
		fmt.Sprintf("fixtures/23.txt:\t2\tthree\n")}

	var buffer bytes.Buffer
	findDupes(&buffer)
	got := buffer.String()
	for _, wantedString := range want {
		if !strings.Contains(got, wantedString) {
			t.Errorf("wanted '%s', wasn't present in '%s'", wantedString, got)
		}
	}
}

package main

import (
	"bytes"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"./echo", "hey", "there"}
	var buffer bytes.Buffer
	want := `0 ./echo
1 hey
2 there
`
	Echo(&buffer)
	got := buffer.String()

	if want != got {
		t.Errorf("got '%s' wanted '%s'", got, want)
	}
}

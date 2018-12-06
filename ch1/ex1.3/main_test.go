package main

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"./echo", "hey", "there", "this", "is", "a", "long", "set", "of", "args"}
	var buffer bytes.Buffer
	b.Run("Slow echo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SlowEcho(&buffer)
		}
	})

	b.Run("Fast echo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FastEcho(&buffer)
		}
	})
}

package main

import (
	"bytes"
	"fmt"
)

func Greet(w *bytes.Buffer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

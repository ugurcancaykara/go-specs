package main

import (
	"bytes"
	"fmt"
)

func main() {

}

// 2- Write the minimal amount of code for the test to run and check the failing test output
func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}

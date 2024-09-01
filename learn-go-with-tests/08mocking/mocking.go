package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 5- To complete matters, let's now wire up our function into a `main` so we have some working software to reassure ourselves we're making progress.
	Countdown(os.Stdout)

	// conclusion: Yes this seems trivial but this approach is what is recommended for any project. Take a thin slice of functionality and make it work end-to-end, backed by tests.
}

// 2- Write the minimal amount of code for the test to run and check the failing test output
func Countdown(out io.Writer) { // 4- we know *bytes.Buffer works, it would be better to use general purpose interface instead
	// 3- write enough code to make it pass
	fmt.Fprint(out, "3")
}

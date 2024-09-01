package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {

	// 5- To complete matters, let's now wire up our function into a `main` so we have some working software to reassure ourselves we're making progress.
	Countdown(os.Stdout)

	// conclusion: Yes this seems trivial but this approach is what is recommended for any project. Take a thin slice of functionality and make it work end-to-end, backed by tests.
}

// 2- Write the minimal amount of code for the test to run and check the failing test output
func Countdown(out io.Writer) { // 4- we know *bytes.Buffer works, it would be better to use general purpose interface instead
	// 3- write enough code to make it pass

	// NOTE: Refactoring the function from 3.(was printing just '3' at this step) to iterative. Not gonna add anymore numbers, check git history to infer what added when

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	// NOTE: Now tests still pass and the software works as intended but we have some problems:
	// FIX: Our tests take 3 seconds to run (check `summary.md section - 8 for explanation`)

	fmt.Fprint(out, finalWord)
}

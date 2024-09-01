package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Mocking-1
// Let's define our dependency as an interface. This lets us then use a real Sleeper in main and a spy sleeper in our tests. By using an interface our Countdown function is oblivious to this and adds some flexibility for the caller.
type Sleeper interface {
	Sleep()
}

// Mocking-2
// Made a design decision that our Countdown function would not be responsible for how long the sleep is. This simplifies our code a little for now at least and means a user of our function can configure that sleepiness however they like.

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Spies are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc. In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.

// Let's create a real sleeper which implements the interface we need

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {

	// 5- To complete matters, let's now wire up our function into a `main` so we have some working software to reassure ourselves we're making progress.
	// Countdown(os.Stdout)

	// Mocking3
	// We can then use it in our real application like so
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)

	// conclusion: Yes this seems trivial but this approach is what is recommended for any project. Take a thin slice of functionality and make it work end-to-end, backed by tests.
}

// 2- Write the minimal amount of code for the test to run and check the failing test output
func Countdown(out io.Writer, sleeper Sleeper) { // 4- we know *bytes.Buffer works, it would be better to use general purpose interface instead

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	// 3- write enough code to make it pass
	// NOTE: Refactoring the function from 3.(was printing just '3' at this step) to iterative. Not gonna add anymore numbers, check git history to infer what added when

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	// NOTE: Now tests still pass and the software works as intended but we have some problems:
	// FIX: Our tests take 3 seconds to run (check `summary.md section - 8 | Countdown example -> for explanation`)

	fmt.Fprint(out, finalWord)
}

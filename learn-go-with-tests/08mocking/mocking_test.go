package main

import (
	"bytes"
	"testing"
)

// 1- Write the test first

func TestCountdown(t *testing.T) {
	// Update the tests to inject a dependency on our Spy and assert that the sleep has been called 3 times.
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}

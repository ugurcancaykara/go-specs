package main

import (
	"bytes"
	"testing"
)

// 1- Write the test first

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

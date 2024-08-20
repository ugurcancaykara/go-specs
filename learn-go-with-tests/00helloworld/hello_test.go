package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World "
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// is needed to tell test suite that this method is a helper. By doing this,
	// when it fails, the line number reported will be in our func call rather than
	// inside our test helper(assertCorrectMessage() func). This helps developers to track down problems more easily
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

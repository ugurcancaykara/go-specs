package main

import "testing"

func TestSearch(t *testing.T) {
	got := Search("key")
	want := "valueofkey"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

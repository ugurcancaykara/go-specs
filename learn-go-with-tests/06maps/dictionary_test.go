package main

import "testing"

func TestSearch(t *testing.T) {

	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")

	got, err := dictionary.Search("test")
	want := "this is just a test"

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

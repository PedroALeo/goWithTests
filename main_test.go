package main

import "testing"

func TestHellojsjgksdjhkdshjkds(t *testing.T) {
	got := Hello("Pedro")
	want := "Hello, Pedro"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Trenton")
	want := "Hello, Trenton"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

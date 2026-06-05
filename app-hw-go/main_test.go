package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting()
	want := "Hello, World!"
	if got != want {
		t.Fatalf("greeting() = %q, want %q", got, want)
	}
}

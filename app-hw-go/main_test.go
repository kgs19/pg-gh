package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting()
	want := "Hello, World!"
	if got != want {
		t.Fatalf("greeting() = %q, want %q", got, want)
	}
}

func TestGreetingDim(t *testing.T) {
    got := greetingDim()
    want := "Hello, Dim!"

    if got != want {
        t.Fatalf("greetingDim() = %q, want %q", got, want)
    }

    if got == "" {
        t.Fatalf("greetingDim() returned empty string")
    }

    if got == greeting() {
        t.Fatalf("greetingDim() should differ from greeting(), both returned %q", got)
    }
}
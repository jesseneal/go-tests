package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jesse")
	want := "Hello, Jesse"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

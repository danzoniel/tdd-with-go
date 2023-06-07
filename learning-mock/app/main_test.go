package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Daniel")
	want := "Olá, Daniel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

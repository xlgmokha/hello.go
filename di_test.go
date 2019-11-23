package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "mo")

	got := buffer.String()
	want := "Hello, mo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

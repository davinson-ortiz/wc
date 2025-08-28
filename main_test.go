package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words.
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Hello there my frieds\n")
	expected := 4
	got := count(b, false, false)

	if expected != got {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("First line\nSecond line\nThird line")
	expected := 3
	got := count(b, true, false)

	if expected != got {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}
func TestCountBytes(t *testing.T) {
	text := "How many bytes need to represent me?"
	b := bytes.NewBufferString(text)
	expected := len([]byte(text))
	got := count(b, false, true)

	if expected != got {
		t.Errorf("Expected %d, got %d instead.\n", expected, got)
	}
}

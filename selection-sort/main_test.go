package main

import "testing"

func TestBubblesort(t *testing.T) {
	// got := Bubblesort()
	// want := 40.0

	got := 0
	want := 1
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

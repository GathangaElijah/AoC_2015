package main

import "testing"

func TestPresentWrap(t *testing.T) {
	l := 1
	w := 1
	h := 10

	expected := 43

	got := PresentWrap(l, w, h)
	if got != expected {
		t.Errorf("Got  %v, expected %v . Error!", got, expected)
	}
}

package main

import "testing"

func TestFloorcalc(t *testing.T) {
	str := ")())())"
	expected := -3

	got := FloorCalc(str)
	if got != expected {
		t.Errorf("Got %v, expected %v. Error", got, expected)
		return
	}
}

func TestBasement(t *testing.T) {
	str := "()())"
	expected := 5

	got := Basement(str)
	if got != expected {
		t.Errorf("Got  %v, expected %v. Error", got, expected)
		return
	}
}
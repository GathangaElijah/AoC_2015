package main

import "testing"

func TestPresentDelivery(t *testing.T) {
	input := "^v^v^v^v^v"
	expected := 2

	got := PresentDeliver(input)
	if got != expected {
		t.Errorf("Got %v, expected %v. Error!", got, expected)
		return
	}
}

func TestSantaAndRobo(t *testing.T) {
	str := "^v^v^v^v^v"
	expected := 11

	got := SantaAndRobo(str)
	if got != expected{
		t.Errorf("Got %v, Expected %v. Error!", got, expected)
		return
	}
}
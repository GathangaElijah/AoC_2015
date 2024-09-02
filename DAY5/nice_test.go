package main

import (
	"testing"
)

func TestNice(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{"Testcase1", "ugknbfddgicrmopn", true},
		{"Testcase2", "jchzalrnumimnmhp", false},
		{"Testcase3", "aaa", true},
		{"Testcase4", "haegwjzuvuyypxyu", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nice(tt.args); got != tt.want {
				t.Errorf("Nice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNice2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{"Testcase1", "qjhvhtzxzqqjkmpb", true},
		{"Testcase2", "xxyxx", true},
		{"Testcase3", "uurcxstgmygtbstg", false},
		{"Testcase4", "ieodomkazucvgmuy", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nice2(tt.args); got != tt.want {
				t.Errorf("Nice2() = %v, want %v", got, tt.want)
			}
		})
	}
}

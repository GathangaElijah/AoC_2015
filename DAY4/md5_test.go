package main

import "testing"

func TestMd5Hash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{"TestCase1", "abcdef", 609043},
		{"TestCase2", "pqrstuv", 1048970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Hash(tt.args); got != tt.want {
				t.Errorf("Md5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func TestLights(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lights(tt.args.s); got != tt.want {
				t.Errorf("Lights() = %v, want %v", got, tt.want)
			}
		})
	}
}

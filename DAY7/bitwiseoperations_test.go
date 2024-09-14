package main

import (
	"reflect"
	"testing"
)

func TestBitwiseAnd(t *testing.T) {
	type args struct {
		a []bool
		b []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
		{"AND TestCase1", args{[]}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitwiseAnd(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

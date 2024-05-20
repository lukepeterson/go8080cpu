package gates

import (
	"fmt"
	"testing"
)

func TestAnd(t *testing.T) {
	type test struct {
		a    bool
		b    bool
		want bool
	}
	tests := []test{
		{a: false, b: false, want: false},
		{a: false, b: true, want: false},
		{a: true, b: false, want: false},
		{a: true, b: true, want: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("And(%t, %t)", tt.a, tt.b), func(t *testing.T) {
			result := And(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("And(%t, %t) returned %t, expected %t", tt.a, tt.b, result, tt.want)
			}
		})
	}
}

func TestOr(t *testing.T) {
	type test struct {
		a    bool
		b    bool
		want bool
	}
	tests := []test{
		{a: false, b: false, want: false},
		{a: false, b: true, want: true},
		{a: true, b: false, want: true},
		{a: true, b: true, want: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Or(%t, %t)", tt.a, tt.b), func(t *testing.T) {
			result := Or(tt.a, tt.b)
			if result != tt.want {
				t.Errorf("Or(%t, %t) returned %t, expected %t", tt.a, tt.b, result, tt.want)
			}
		})
	}
}

func TestNot(t *testing.T) {
	type test struct {
		a    bool
		want bool
	}
	tests := []test{
		{a: false, want: true},
		{a: true, want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Not(%t)", tt.a), func(t *testing.T) {
			result := Not(tt.a)
			if result != tt.want {
				t.Errorf("Not(%t) returned %t, expected %t", tt.a, result, tt.want)
			}
		})
	}
}

func TestNand(t *testing.T) {
	type test struct {
		a    bool
		b    bool
		want bool
	}
	tests := []test{
		{a: false, b: false, want: true},
		{a: false, b: true, want: true},
		{a: true, b: false, want: true},
		{a: true, b: true, want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Nand(%t, %t)", tt.a, tt.b), func(t *testing.T) {
			result := Not(And(tt.a, tt.b))
			if result != tt.want {
				t.Errorf("Nand(%t, %t) returned %t, expected %t", tt.a, tt.b, result, tt.want)
			}
		})
	}
}

package cpu

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	tests := []struct {
		a         bool
		b         bool
		wantSum   bool
		wantCarry bool
	}{
		{a: false, b: false, wantSum: false, wantCarry: false},
		{a: false, b: true, wantSum: true, wantCarry: false},
		{a: true, b: false, wantSum: true, wantCarry: false},
		{a: true, b: true, wantSum: false, wantCarry: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("halfAdder(%t, %t)", tt.a, tt.b), func(t *testing.T) {
			gotSum, gotCarry := halfAdder(tt.a, tt.b)
			if gotSum != tt.wantSum {
				t.Errorf("halfAdder() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotCarry != tt.wantCarry {
				t.Errorf("halfAdder() gotCarry = %v, want %v", gotCarry, tt.wantCarry)
			}
		})
	}
}

func TestFullAdder(t *testing.T) {
	tests := []struct {
		a            bool
		b            bool
		carryIn      bool
		wantSum      bool
		wantCarryOut bool
	}{
		{a: false, b: false, carryIn: false, wantSum: false, wantCarryOut: false},
		{a: false, b: false, carryIn: true, wantSum: true, wantCarryOut: false},
		{a: false, b: true, carryIn: false, wantSum: true, wantCarryOut: false},
		{a: false, b: true, carryIn: true, wantSum: false, wantCarryOut: true},
		{a: true, b: false, carryIn: false, wantSum: true, wantCarryOut: false},
		{a: true, b: false, carryIn: true, wantSum: false, wantCarryOut: true},
		{a: true, b: true, carryIn: false, wantSum: false, wantCarryOut: true},
		{a: true, b: true, carryIn: true, wantSum: true, wantCarryOut: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("fullAdder(%t, %t, %t)", tt.a, tt.b, tt.carryIn), func(t *testing.T) {
			gotSum, gotCarryOut := fullAdder(tt.a, tt.b, tt.carryIn)
			if gotSum != tt.wantSum {
				t.Errorf("fullAdder() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotCarryOut != tt.wantCarryOut {
				t.Errorf("fullAdder() gotCarryOut = %v, want %v", gotCarryOut, tt.wantCarryOut)
			}
		})
	}
}

func TestALUAdd(t *testing.T) {
	tests := []struct {
		name         string
		a            [8]bool
		b            [8]bool
		wantResult   [8]bool
		wantCarryOut bool
	}{
		{
			name:         "add 0b0000_0001 + 0b0000_0001 (no carry)",
			a:            [8]bool{false, false, false, false, false, false, false, true},
			b:            [8]bool{false, false, false, false, false, false, false, true},
			wantResult:   [8]bool{false, false, false, false, false, false, true, false},
			wantCarryOut: false,
		},
		{
			name:         "add 0b0100_0000 + 0b0000_0001 (no carry)",
			a:            [8]bool{false, true, false, false, false, false, false, false},
			b:            [8]bool{false, false, false, false, false, false, false, true},
			wantResult:   [8]bool{false, true, false, false, false, false, false, true},
			wantCarryOut: false,
		},
		{
			name:         "add 0b1000_0000 + 0b1000_0000 (with carry)",
			a:            [8]bool{true, false, false, false, false, false, false, false},
			b:            [8]bool{true, false, false, false, false, false, false, false},
			wantResult:   [8]bool{false, false, false, false, false, false, false, false},
			wantCarryOut: true,
		},
		{
			name:         "add 0b1001_0101 + 0b0101_0100 (no carry)",
			a:            [8]bool{true, false, false, true, false, true, false, true},
			b:            [8]bool{false, true, false, true, false, true, false, false},
			wantResult:   [8]bool{true, true, true, false, true, false, false, true},
			wantCarryOut: false,
		},
		{
			name:         "add 0b1111_1111 + 0b0000_0000 (no carry)",
			a:            [8]bool{true, true, true, true, true, true, true, true},
			b:            [8]bool{false, false, false, false, false, false, false, false},
			wantResult:   [8]bool{true, true, true, true, true, true, true, true},
			wantCarryOut: false,
		},
		{
			name:         "add 0b1111_1111 + 0b0000_0001 (carry)",
			a:            [8]bool{true, true, true, true, true, true, true, true},
			b:            [8]bool{false, false, false, false, false, false, false, true},
			wantResult:   [8]bool{false, false, false, false, false, false, false, false},
			wantCarryOut: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ALU.Add(%t, %t)", tt.a, tt.b), func(t *testing.T) {
			alu := &ALU{
				a: tt.a,
				b: tt.b,
			}

			gotResult, gotCarryOut := alu.Add()
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ALU.Add() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotCarryOut != tt.wantCarryOut {
				t.Errorf("ALU.Add() gotCarryOut = %v, want %v", gotCarryOut, tt.wantCarryOut)
			}
		})
	}
}

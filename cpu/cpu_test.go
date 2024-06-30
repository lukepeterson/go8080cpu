package cpu

import "testing"

func TestCPUGetFlags(t *testing.T) {
	tests := []struct {
		name  string
		flags Flags
		want  byte
	}{
		{
			name:  "all flags unset",
			flags: Flags{Sign: false, Zero: false, AuxCarry: false, Parity: false, Carry: false},
			want:  0b00000010,
		},
		{
			name:  "carry flag set",
			flags: Flags{Sign: false, Zero: false, AuxCarry: false, Parity: false, Carry: true},
			want:  0b00000011,
		},
		{
			name:  "parity flag set",
			flags: Flags{Sign: false, Zero: false, AuxCarry: false, Parity: true, Carry: false},
			want:  0b00000110,
		},
		{
			name:  "auxcarry flag set",
			flags: Flags{Sign: false, Zero: false, AuxCarry: true, Parity: false, Carry: false},
			want:  0b00010010,
		},
		{
			name:  "zero flag set",
			flags: Flags{Sign: false, Zero: true, AuxCarry: false, Parity: false, Carry: false},
			want:  0b01000010,
		},
		{
			name:  "sign flag set",
			flags: Flags{Sign: true, Zero: false, AuxCarry: false, Parity: false, Carry: false},
			want:  0b10000010,
		},
		{
			name:  "all flags set",
			flags: Flags{Sign: true, Zero: true, AuxCarry: true, Parity: true, Carry: true},
			want:  0b11010111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := CPU{
				flags: tt.flags,
			}
			if got := cpu.getFlags(); got != tt.want {
				t.Errorf("CPU.setFlags() = 0b%08b, want 0b%08b", got, tt.want)
			}
		})
	}
}

func TestCPUSetFlags(t *testing.T) {
	tests := []struct {
		name  string
		flags byte
	}{
		{
			name:  "all flags unset",
			flags: 0b00000010,
		},
		{
			name:  "carry flag set",
			flags: 0b00000011,
		},
		{
			name:  "parity flag set",
			flags: 0b00000110,
		},
		{
			name:  "auxcarry flag set",
			flags: 0b00010010,
		},
		{
			name:  "zero flag set",
			flags: 0b01000010,
		},
		{
			name:  "sign flag set",
			flags: 0b10000010,
		},
		{
			name:  "all flags set",
			flags: 0b11010111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu := &CPU{}
			cpu.setFlags(tt.flags) // setFlags() has no return, so use getFlags to confirm they were set correctly.
			if got := cpu.getFlags(); got != tt.flags {
				t.Errorf("CPU.setFlags() = 0b%08b, want 0b%08b", got, tt.flags)
			}
		})
	}
}

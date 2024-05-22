package cpu

import (
	"testing"
	"time"
)

func TestCPU_LDA(t *testing.T) {
	testCases := []struct {
		name   string
		memory map[uint16]byte
		want   byte
	}{
		{
			name: "LDA from address 0x0010 (low order byte only)",
			memory: map[uint16]byte{
				0x00: 0x3A, // Program
				0x01: 0x10,
				0x02: 0x00,
				0x03: 0x4C, // HLT
				0x10: 0x69, // Data
			},
			want: 0x69,
		},
		{
			name: "LDA from address 0x0100 (high order byte only)",
			memory: map[uint16]byte{
				0x00:   0x3A, // Program
				0x01:   0x00,
				0x02:   0x01,
				0x03:   0x4C, // HLT
				0x0100: 0xFE, // Data
			},
			want: 0xFE,
		},
		{
			name: "LDA from address 0x0100 (both high and low order bytes)",
			memory: map[uint16]byte{
				0x00:   0x3A, // Program
				0x01:   0x5A,
				0x02:   0x01,
				0x03:   0x4C, // HLT
				0x015A: 0xAA, // Data
			},
			want: 0xAA,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			goCPU := NewCPU(1*time.Millisecond, 512)
			for addr, value := range tc.memory {
				goCPU.Memory[addr] = value
			}
			goCPU.Run()

			got := goCPU.A
			if got != tc.want {
				t.Errorf("LDA() returned 0x%02X, but expected 0x%02X", got, tc.want)
			}
		})
	}
}

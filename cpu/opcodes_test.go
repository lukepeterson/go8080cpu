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
				0x00: 0x3A, // LDA
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
				0x00:   0x3A, // LDA
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
				0x00:   0x3A, // LDA
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
			goCPU := NewCPU(
				1*time.Millisecond,
				NewMemory(512),
			)

			goCPU.Load(tc.memory)
			goCPU.Run()
			got := goCPU.A
			if got != tc.want {
				t.Errorf("LDA() returned 0x%02X, but expected 0x%02X", got, tc.want)
			}
		})
	}
}

func TestCPU_MVI_r(t *testing.T) {
	testCases := []struct {
		name     string
		memory   map[uint16]byte
		want     byte
		register func(cpu *CPU) byte
	}{
		{
			name: "MOV A with value 0x11",
			memory: map[uint16]byte{
				0x00: 0x3E, // MVI A
				0x01: 0x11,
				0x03: 0x4C, // HLT
			},
			want:     0x11,
			register: func(cpu *CPU) byte { return cpu.A },
		},
		{
			name: "MOV B with value 0x22",
			memory: map[uint16]byte{
				0x00: 0x06, // MVI B
				0x01: 0x22,
				0x03: 0x4C, // HLT
			},
			want:     0x22,
			register: func(cpu *CPU) byte { return cpu.B },
		},
		{
			name: "MOV C with value 0x33",
			memory: map[uint16]byte{
				0x00: 0x0E, // MVI C
				0x01: 0x33,
				0x03: 0x4C, // HLT
			},
			want:     0x33,
			register: func(cpu *CPU) byte { return cpu.C },
		},
		{
			name: "MOV D with value 0x44",
			memory: map[uint16]byte{
				0x00: 0x10, // MVI D
				0x01: 0x44,
				0x03: 0x4C, // HLT
			},
			want:     0x44,
			register: func(cpu *CPU) byte { return cpu.D },
		},
		{
			name: "MOV E with value 0x55",
			memory: map[uint16]byte{
				0x00: 0x1E, // MVI E
				0x01: 0x55,
				0x03: 0x4C, // HLT
			},
			want:     0x55,
			register: func(cpu *CPU) byte { return cpu.E },
		},
		{
			name: "MOV H with value 0x66",
			memory: map[uint16]byte{
				0x00: 0x1A, // MVI H
				0x01: 0x66,
				0x03: 0x4C, // HLT
			},
			want:     0x66,
			register: func(cpu *CPU) byte { return cpu.H },
		},
		{
			name: "MOV L with value 0x77",
			memory: map[uint16]byte{
				0x00: 0x2E, // MVI L
				0x01: 0x77,
				0x03: 0x4C, // HLT
			},
			want:     0x77,
			register: func(cpu *CPU) byte { return cpu.L },
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			goCPU := NewCPU(1*time.Millisecond, NewMemory(4))
			goCPU.Load(tc.memory)
			goCPU.Run()
			got := tc.register(goCPU)
			if got != tc.want {
				t.Errorf("%s returned 0x%02X, but expected 0x%02X", tc.name, got, tc.want)
			}
		})
	}
}

func TestCPU_MVI_M(t *testing.T) {
	testCases := []struct {
		name   string
		memory map[uint16]byte
		want   byte
		L      byte
		H      byte
	}{
		{
			name: "MOV M with value of data stored at HL (low order byte only)",
			memory: map[uint16]byte{
				0x00: 0x24, // MVI M
				0x01: 0x11,
				0x02: 0x4C, // HLT
			},
			want: 0x11,
			L:    0x22,
			H:    0x00,
		},
		{
			name: "MOV M with value of data stored at HL (high order byte only)",
			memory: map[uint16]byte{
				0x00: 0x24, // MVI M
				0x01: 0x33,
				0x02: 0x4C, // HLT
			},
			want: 0x33,
			L:    0x00,
			H:    0x44,
		},
		{
			name: "MOV M with value of data stored at HL (both high and low order bytes)",
			memory: map[uint16]byte{
				0x00: 0x24, // MVI M
				0x01: 0x66,
				0x02: 0x4C, // HLT
			},
			want: 0x66,
			L:    0x77,
			H:    0x88,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			goCPU := NewCPU(1*time.Millisecond, NewMemory(65535))
			goCPU.Load(tc.memory)
			goCPU.L = tc.L
			goCPU.H = tc.H
			goCPU.Run()
			got := goCPU.bus.ReadByte(joinBytes(tc.L, tc.H))
			if got != tc.want {
				t.Errorf("%s returned 0x%02X, but expected 0x%02X", tc.name, got, tc.want)
			}
		})
	}
}

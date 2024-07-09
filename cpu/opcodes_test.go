package cpu

import (
	"testing"

	"github.com/lukepeterson/go8080assembler/assembler"
)

func TestCPUInstructions(t *testing.T) {
	testCases := []struct {
		name    string
		code    string
		initCPU *CPU
		wantCPU *CPU
		wantErr bool
	}{
		{
			name: "MOV B, B",
			code: `
				MOV B, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01},
		},
		{
			name: "MOV B, C",
			code: `
				MOV B, C
				HLT
				`,
			initCPU: &CPU{C: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x02, C: 0x02},
		},
		{
			name: "MOV B, D",
			code: `
				MOV B, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, D: 0x01},
		},
		{
			name: "MOV B, E",
			code: `
				MOV B, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, E: 0x01},
		},
		{
			name: "MOV B, H",
			code: `
				MOV B, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, H: 0x01},
		},
		{
			name: "MOV B, L",
			code: `
				MOV B, L
				HLT
				`,
			initCPU: &CPU{L: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x02, L: 0x02},
		},
		{
			name: "MOV B, M",
			code: `
				MVI M, 0x55
				MOV B, M
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x01, Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{B: 0x055, H: 0x01, L: 0x01},
		},
		{
			name: "MOV B, A",
			code: `
				MOV B, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, A: 0x01},
		},
		{
			name: "MOV C, B",
			code: `
				MOV C, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, C: 0x01},
		},
		{
			name: "MOV C, C",
			code: `
				MOV C, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01},
		},
		{
			name: "MOV C, D",
			code: `
				MOV C, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, D: 0x01},
		},
		{
			name: "MOV C, E",
			code: `
				MOV C, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, E: 0x01},
		},
		{
			name: "MOV C, H",
			code: `
				MOV C, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, H: 0x01},
		},
		{
			name: "MOV C, L",
			code: `
				MOV C, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, L: 0x01},
		},
		{
			name: "MOV C, M",
			code: `
				MVI M, 0x55
				MOV C, M
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x01, Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{C: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV C, A",
			code: `
				MOV C, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, C: 0x01},
		},
		{
			name: "MOV D, B",
			code: `
				MOV D, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, D: 0x01},
		},
		{
			name: "MOV D, C",
			code: `
				MOV D, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, D: 0x01},
		},
		{
			name: "MOV D, D",
			code: `
				MOV D, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01},
		},
		{
			name: "MOV D, E",
			code: `
				MOV D, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, E: 0x01},
		},
		{
			name: "MOV D, H",
			code: `
				MOV D, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, H: 0x01},
		},
		{
			name: "MOV D, L",
			code: `
				MOV D, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, L: 0x01},
		},
		{
			name: "MOV D, M",
			code: `
				MVI M, 0x55
				MOV D, M
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x01, Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{D: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV D, A",
			code: `
				MOV D, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, D: 0x01},
		},
		{
			name: "MOV E, B",
			code: `
				MOV E, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, E: 0x01},
		},
		{
			name: "MOV E, C",
			code: `
				MOV E, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, E: 0x01},
		},
		{
			name: "MOV E, D",
			code: `
				MOV E, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, E: 0x01},
		},
		{
			name: "MOV E, E",
			code: `
				MOV E, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x01},
		},
		{
			name: "MOV E, H",
			code: `
				MOV E, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x01, H: 0x01},
		},
		{
			name: "MOV E, L",
			code: `
				MOV E, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x01, L: 0x01},
		},
		{
			name: "MOV E, M",
			code: `
				MVI M, 0x55
				MOV E, M
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x01, Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{E: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV E, A",
			code: `
				MOV E, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, E: 0x01},
		},
		{
			name: "MOV H, B",
			code: `
				MOV H, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, H: 0x01},
		},
		{
			name: "MOV H, C",
			code: `
				MOV H, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, H: 0x01},
		},
		{
			name: "MOV H, D",
			code: `
				MOV H, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, H: 0x01},
		},
		{
			name: "MOV H, E",
			code: `
				MOV H, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x01, H: 0x01},
		},
		{
			name: "MOV H, H",
			code: `
				MOV H, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x01},
		},
		{
			name: "MOV H, L",
			code: `
				MOV H, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x01, L: 0x01},
		},
		{
			name: "MOV H, M",
			code: `
				LXI H, 0x0101
				MVI M, 0x02
				MOV H, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{H: 0x02, L: 0x01},
		},
		{
			name: "MOV H, A",
			code: `
				MOV H, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, H: 0x01},
		},
		{
			name: "MOV L, B",
			code: `
				MOV L, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, L: 0x01},
		},
		{
			name: "MOV L, C",
			code: `
				MOV L, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x01, L: 0x01},
		},
		{
			name: "MOV L, D",
			code: `
				MOV L, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, L: 0x01},
		},
		{
			name: "MOV L, E",
			code: `
				MOV L, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x01, L: 0x01},
		},
		{
			name: "MOV L, H",
			code: `
				MOV L, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x01, L: 0x01},
		},
		{
			name: "MOV L, L",
			code: `
				MOV L, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{L: 0x01},
		},
		{
			name: "MOV L, M",
			code: `
				LXI H, 0x0101
				MVI M, 0x02
				MOV L, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{H: 0x01, L: 0x02},
		},
		{
			name: "MOV L, A",
			code: `
				MOV L, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, L: 0x01},
		},
		{
			name: "MOV M, B",
			code: `
				LXI H, 0x0101
				MVI B, 0x55
				MOV M, B
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x55, B: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV M, C",
			code: `
				LXI H, 0x0101
				MVI C, 0x55
				MOV M, C
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x55, C: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV M, D",
			code: `
				LXI H, 0x0101
				MVI D, 0x55
				MOV M, D
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x55, D: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV M, E",
			code: `
				LXI H, 0x0101
				MVI E, 0x55
				MOV M, E
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x55, E: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV M, H",
			code: `
				LXI H, 0x0101
				MVI H, 0x02
				MOV M, H
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFF+2)}},
			wantCPU: &CPU{H: 0x02, L: 0x01},
		},
		{
			name: "MOV M, L",
			code: `
				LXI H, 0x0101
				MVI L, 0x02
				MOV M, L
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{H: 0x01, L: 0x02},
		},
		{
			name: "MOV M, A",
			code: `
				LXI H, 0x0101
				MVI A, 0x02
				MOV M, A
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x02, H: 0x01, L: 0x01},
		},
		{
			name: "MOV A, B",
			code: `
				MOV A, B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, A: 0x01},
		},
		{
			name: "MOV A, C",
			code: `
				MOV A, C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, C: 0x01},
		},
		{
			name: "MOV A, D",
			code: `
				MOV A, D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, D: 0x01},
		},
		{
			name: "MOV A, E",
			code: `
				MOV A, E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, E: 0x01},
		},
		{
			name: "MOV A, H",
			code: `
				MOV A, H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, H: 0x01},
		},
		{
			name: "MOV A, L",
			code: `
				MOV A, L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, L: 0x01},
		},
		{
			name: "MOV A, M",
			code: `
				MVI M, 0x55
				MOV A, M
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x01, Bus: &Memory{Data: make([]byte, 0xFF+3)}},
			wantCPU: &CPU{A: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "MOV A, A",
			code: `
				MOV A, A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},

		{
			name: "LXI B",
			code: `
				LXI B, 0x3344
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{B: 0x33, C: 0x44},
		},
		{
			name: "LXI D",
			code: `
				LXI D, 0x3344
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{D: 0x33, E: 0x44},
		},
		{
			name: "LXI H",
			code: `
				LXI H, 0x3344
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{H: 0x33, L: 0x44},
		},
		{
			name: "STAX B",
			code: `
				LXI B, 0x0101
				LXI H, 0x0101
				MVI A, 0x55
				STAX B
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, B: 0x01, C: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "STAX D",
			code: `
				LXI D, 0x0101
				LXI H, 0x0101
				MVI A, 0x55
				STAX D
				MOV A, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, D: 0x01, E: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "LDAX B",
			code: `
				LXI B, 0x0101
				LXI H, 0x0101
				MVI M, 0x55
				LDAX B
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, B: 0x01, C: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "LDAX D",
			code: `
				LXI D, 0x0101
				LXI H, 0x0101
				MVI M, 0x55
				LDAX D
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, D: 0x01, E: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "STA",
			code: `
				MVI A, 0x55
				STA 0x3344
				LXI H 0x3344
				MOV B, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, B: 0x55, H: 0x33, L: 0x44},
		},
		{
			name: "LDA",
			code: `
				LXI H 0x0101
				MVI M, 0x55
				LDA 0x0101
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "SHLD",
			code: `
				LXI H, 0x4455
				SHLD 0x2000
				LXI H, 0x2000
				MOV A, M
				MOV C, A
				INX H
				MOV A, M
				MOV D, A
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x44, C: 0x55, D: 0x44, H: 0x20, L: 0x01},
		},
		{
			name: "LHLD",
			code: `
				MVI A, 0x33
				STA 0x2000
				MVI A, 0x44
				STA 0x2001
				LHLD 0x2000
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x44, H: 0x44, L: 0x33},
		},
		{
			name: "XCHG",
			code: `
				XCHG
				HLT
			`,
			initCPU: &CPU{D: 0x33, E: 0x44, H: 0x55, L: 0x66, Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{D: 0x55, E: 0x66, H: 0x33, L: 0x44},
		},
		{
			name: "PUSH B",
			code: `
				LXI SP, 0x1000
				MVI B, 0x12
				MVI C, 0x34
				PUSH B
				LXI H, 0x0FFE
				MOV A, M
				LXI H, 0x0FFF
				MOV B, M
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x34, B: 0x12, C: 0x34, H: 0x0F, L: 0xFF, stackPointer: 0x0FFE},
		},
		{
			name: "PUSH D",
			code: `
				LXI SP, 0x1000
				MVI D, 0x12
				MVI E, 0x34
				PUSH D
				LXI H, 0x0FFE
				MOV A, M
				LXI H, 0x0FFF
				MOV B, M
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x34, B: 0x12, D: 0x12, E: 0x34, H: 0x0F, L: 0xFF, stackPointer: 0x0FFE},
		},
		{
			name: "PUSH H",
			code: `
				LXI SP, 0x1000
				MVI H, 0x12
				MVI L, 0x34
				PUSH H
				LXI H, 0x0FFE
				MOV A, M
				LXI H, 0x0FFF
				MOV B, M
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x34, B: 0x12, H: 0x0F, L: 0xFF, stackPointer: 0x0FFE},
		},
		{
			name: "PUSH PSW",
			code: `
				LXI SP, 0x1000
				MVI A, 0x12
				PUSH PSW
				LXI H, 0x0FFE
				MOV A, M
				LXI H, 0x0FFF
				MOV B, M
				HLT
			`,
			initCPU: &CPU{flags: Flags{Sign: true, Parity: true}, Bus: &Memory{Data: make([]byte, 0x10000)}},
			wantCPU: &CPU{A: 0x86, B: 0x12, H: 0x0F, L: 0xFF, flags: Flags{Sign: true, Parity: true}, stackPointer: 0x0FFE},
		},
		{
			name: "POP B",
			code: `
				LXI H, 0xFFFE
				MVI M, 0x12
				INX H
				MVI M, 0x34
				LXI SP, 0xFFFE
				POP B
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF+1)}},
			wantCPU: &CPU{B: 0x34, C: 0x12, H: 0xFF, L: 0xFF, stackPointer: 0},
		},
		{
			name: "POP D",
			code: `
				LXI H, 0xFFFE
				MVI M, 0x12
				INX H
				MVI M, 0x34
				LXI SP, 0xFFFE
				POP D
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF+1)}},
			wantCPU: &CPU{D: 0x34, E: 0x12, H: 0xFF, L: 0xFF, stackPointer: 0},
		},
		{
			name: "POP H",
			code: `
				LXI H, 0xFFFE
				MVI M, 0x12
				INX H
				MVI M, 0x34
				LXI SP, 0xFFFE
				POP H
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF+1)}},
			wantCPU: &CPU{H: 0x34, L: 0x12, stackPointer: 0},
		},
		{
			name: "POP PSW",
			code: `
				LXI H, 0xFFFE
				MVI M, 0x86
				INX H
				MVI M, 0x34
				LXI SP, 0xFFFE
				POP PSW
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF+1)}},
			wantCPU: &CPU{A: 0x34, H: 0xFF, L: 0xFF, flags: Flags{Sign: true, Parity: true}, stackPointer: 0},
		},
		{
			name: "XTHL",
			code: `
				LXI SP, 0xFFFF
				LXI H, 0x1122
				LXI B, 0x3344
				PUSH B
				LXI B, 0x5566
				PUSH B
				XTHL
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{B: 0x55, C: 0x66, H: 0x55, L: 0x66, stackPointer: 0xFFFB},
		},
		{
			name: "SPHL",
			code: `
				LXI H, 0xFFEE
				SPHL
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{H: 0xFF, L: 0xEE, stackPointer: 0xFFEE},
		},
		{
			name: "LXI SP",
			code: `
				LXI SP, 0x1234
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{stackPointer: 0x1234},
		},
		{
			name: "INX SP from 0x1000",
			code: `
				INX SP
				HLT
			`,
			initCPU: &CPU{stackPointer: 0x1000, Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{stackPointer: 0x1001},
		},
		{
			name: "INX SP from 0xFFFF (test overflow works)",
			code: `
				INX SP
				HLT
			`,
			initCPU: &CPU{stackPointer: 0xFFFF, Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{stackPointer: 0x0000},
		},
		{
			name: "DCX SP from 0x1000",
			code: `
				DCX SP
				HLT
			`,
			initCPU: &CPU{stackPointer: 0x1000, Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{stackPointer: 0x0FFF},
		},
		{
			name: "DCX SP from 0x0000 (test underflow works)",
			code: `
				DCX SP
				HLT
			`,
			initCPU: &CPU{stackPointer: 0x0000, Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{stackPointer: 0xFFFF},
		},
		{
			name: "JMP",
			code: `
				JMP 0x04
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{programCounter: 0x0005},
		},
		{
			name: "JC (carry set - jump)",
			code: `
				STC
				JC 0x05
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{flags: Flags{Carry: true}, programCounter: 0x0006},
		},
		{
			name: "JC (carry not set - don't jump)",
			code: `
				JC 0x04
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{programCounter: 0x0004},
		},
		{
			name: "JNC (carry set - don't jump)",
			code: `
				STC
				JNC 0x05
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{flags: Flags{Carry: true}, programCounter: 0x0005},
		},
		{
			name: "JNC (carry not set - jump)",
			code: `
				JNC 0x04
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{programCounter: 0x0005},
		},
		{
			name: "JZ (zero set - jump)",
			code: `
				CMP A
				JZ 0x05
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{flags: Flags{Zero: true, Parity: true}, programCounter: 0x0006},
		},
		{
			name: "JZ (zero not set - don't jump)",
			code: `
				JZ 0x04
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{programCounter: 0x0004},
		},
		{
			name: "JNZ (zero set - jump)",
			code: `
				CMP A
				JNZ 0x05
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{flags: Flags{Zero: true, Parity: true}, programCounter: 0x0005},
		},
		{
			name: "JNZ (zero not set - jump)",
			code: `
				JNZ 0x04
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{programCounter: 0x0005},
		},
		{
			name: "JP (sign flag set - don't jump)",
			code: `
				MVI A, 0x7F
				INR A
				JP 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x80, flags: Flags{Sign: true, AuxCarry: true}, programCounter: 0x0007},
		},
		{
			name: "JP (sign flag not set - jump)",
			code: `
				MVI A, 0x7E
				INR A
				JP 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x7F, programCounter: 0x0008},
		},
		{
			name: "JM (sign flag set - jump)",
			code: `
				MVI A, 0x7F
				INR A
				JM 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x80, flags: Flags{Sign: true, AuxCarry: true}, programCounter: 0x0008},
		},
		{
			name: "JM (sign flag not set - don't jump)",
			code: `
				MVI A, 0x7E
				INR A
				JM 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x7F, programCounter: 0x0007},
		},
		{
			name: "JPE (parity even - jump)",
			code: `
				MVI A, 0x02
				INR A
				JPE 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}, programCounter: 0x0008},
		},
		{
			name: "JPE (parity odd - don't jump)",
			code: `
				MVI A, 0x01
				INR A
				JPE 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x02, programCounter: 0x0007},
		},
		{
			name: "JPO (parity even - don't jump)",
			code: `
				MVI A, 0x02
				INR A
				JPO 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}, programCounter: 0x0007},
		},
		{
			name: "JPO (parity odd - jump)",
			code: `
				MVI A, 0x01
				INR A
				JPO 0x07
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x02, programCounter: 0x0008},
		},
		{
			name: "PCHL",
			code: `
				LXI H, 0x0005
				PCHL
				HLT
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{L: 0x05, programCounter: 0x0006},
		},

		// {
		// 	name: "CALL",
		// 	code: `
		// 		CALL 0x0005      ; Call subroutine at address 0x0006
		// 		NOP              ; This instruction should be executed after the subroutine returns
		// 		HLT              ; Halt here after subroutine returns
		// 		POP B
		// 		HLT              ; Halt within subroutine (to check the jump)
		// 		`,
		// 	initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF+2)}},
		// 	wantCPU: &CPU{B: 0x00, C: 0x03, programCounter: 0x0007},
		// },

		// CALL
		// CC
		// CNC
		// CZ
		// CNZ
		// CP
		// CM
		// CPE
		// CPO

		// RET
		// RC
		// RNC
		// RZ
		// RNZ
		// RP
		// RM
		// RPE
		// RPO

		// RST 0
		// RST 1
		// RST 2
		// RST 3
		// RST 4
		// RST 5
		// RST 6
		// RST 7

		{
			name: "INR A from 0x01",
			code: `
				INR A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02},
		},
		{
			name: "DCR A from 0x03",
			code: `
				DCR A
				HLT
				`,
			initCPU: &CPU{A: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02},
		},
		{
			name: "INR A from 0x02 (test parity flag set)",
			code: `
				INR A
				HLT
				`,
			initCPU: &CPU{A: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}},
		},
		{
			name: "DCR A from 0x04 (test parity flag set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: &CPU{A: 0x04, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}},
		},
		{
			name: "INR A from 0x7F (test sign flag set)",
			code: `
				INR A
				HLT
				`,
			initCPU: &CPU{A: 0x7F, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x80, flags: Flags{AuxCarry: true, Sign: true}},
		},
		{
			name: "DCR A from 0x81 (test sign flag set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: &CPU{A: 0x81, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x80, flags: Flags{Sign: true}},
		},
		{
			name: "INR A from 0x80 (test sign and parity flags set)",
			code: `
				INR A
				HLT
				`,
			initCPU: &CPU{A: 0x80, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x81, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "DCR A from 0x82 (test sign and parity flags set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: &CPU{A: 0x82, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x81, flags: Flags{Sign: true, Parity: true}},
		},

		{
			name: "INR A from 0xFF (test zero and parity flags set)",
			code: `
				INR A
				HLT
				`,
			initCPU: &CPU{A: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, AuxCarry: true, Parity: true}},
		},
		{
			name: "DCR A from 0x01 (test zero and parity flags set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "INR B from 0x01",
			code: `
				INR B
				HLT
				`,
			initCPU: &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x02},
		},
		{
			name: "DCR B from 0x03",
			code: `
				DCR B
				HLT
				`,
			initCPU: &CPU{B: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x02},
		},
		{
			name: "INR C from 0x01",
			code: `
				INR C
				HLT
				`,
			initCPU: &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x02},
		},
		{
			name: "DCR C from 0x03",
			code: `
				DCR C
				HLT
				`,
			initCPU: &CPU{C: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{C: 0x02},
		},
		{
			name: "INR D from 0x01",
			code: `
				INR D
				HLT
				`,
			initCPU: &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x02},
		},
		{
			name: "DCR D from 0x03",
			code: `
				DCR D
				HLT
				`,
			initCPU: &CPU{D: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x02},
		},
		{
			name: "INR E from 0x01",
			code: `
				INR E
				HLT
				`,
			initCPU: &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x02},
		},
		{
			name: "DCR E from 0x03",
			code: `
				DCR E
				HLT
				`,
			initCPU: &CPU{E: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{E: 0x02},
		},
		{
			name: "INR H from 0x01",
			code: ` INR H
				HLT
				`,
			initCPU: &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x02},
		},
		{
			name: "DCR H from 0x03",
			code: `
				DCR H
				HLT
				`,
			initCPU: &CPU{H: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x02},
		},
		{
			name: "INR L from 0x01",
			code: `
				INR L
				HLT
				`,
			initCPU: &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{L: 0x02},
		},
		{
			name: "DCR L from 0x03",
			code: `
				DCR L
				HLT
				`,
			initCPU: &CPU{L: 0x03, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{L: 0x02},
		},
		{
			name: "INR M from 0x01",
			code: `
				MVI M, 0x01
				INR M
				MOV A, M
				HLT
				`,
			initCPU: &CPU{L: 0x16, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02, L: 0x16},
		},
		{
			name: "DCR M from 0x03",
			code: `
				MVI M, 0x03
				DCR M
				MOV A, M
				HLT
				`,
			initCPU: &CPU{L: 0x16, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02, L: 0x16},
		},
		{
			name: "INX B from 0x00FF",
			code: `
				INX B
				HLT
				`,
			initCPU: &CPU{B: 0x00, C: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x01, C: 0x00},
		},
		{
			name: "INX B from 0xFFFF",
			code: `
				INX B
				HLT
				`,
			initCPU: &CPU{B: 0xFF, C: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x00, C: 0x00},
		},
		{
			name: "INX D from 0x00FF",
			code: `
				INX D
				HLT
				`,
			initCPU: &CPU{D: 0x00, E: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x01, E: 0x00},
		},
		{
			name: "INX D from 0xFFFF",
			code: `
				INX D
				HLT
				`,
			initCPU: &CPU{D: 0xFF, E: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x00, E: 0x00},
		},
		{
			name: "INX H from 0x00FF",
			code: `
				INX H
				HLT
				`,
			initCPU: &CPU{H: 0x00, L: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x01, L: 0x00},
		},
		{
			name: "INX H from 0xFFFF",
			code: `
				INX H
				HLT
				`,
			initCPU: &CPU{H: 0xFF, L: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x00, L: 0x00},
		},
		{
			name: "DCX B from 0x0100",
			code: `
				DCX B
				HLT
				`,
			initCPU: &CPU{B: 0x01, C: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0x00, C: 0xFF},
		},
		{
			name: "DCX B from 0x0000",
			code: `
				DCX B
				HLT
				`,
			initCPU: &CPU{B: 0x00, C: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{B: 0xFF, C: 0xFF},
		},
		{
			name: "DCX D from 0x0100",
			code: `
				DCX D
				HLT
				`,
			initCPU: &CPU{D: 0x01, E: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0x00, E: 0xFF},
		},
		{
			name: "DCX D from 0x0000",
			code: `
				DCX D
				HLT
				`,
			initCPU: &CPU{D: 0x00, E: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{D: 0xFF, E: 0xFF},
		},
		{
			name: "DCX H from 0x0100",
			code: `
				DCX H
				HLT
				`,
			initCPU: &CPU{H: 0x01, L: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0x00, L: 0xFF},
		},
		{
			name: "DCX H from 0x0000",
			code: `
				DCX H
				HLT
				`,
			initCPU: &CPU{H: 0x00, L: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{H: 0xFF, L: 0xFF},
		},
		{
			name: "ADD B (no carry or overflow)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x12, B: 0x34, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x46, B: 0x34},
		},
		{
			name: "ADD B (zero result)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x00, B: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ADD B (carry out on bit eight)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0xFF, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x01, flags: Flags{Zero: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "ADD B (aux carry with sign)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x7F, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x80, B: 0x01, flags: Flags{AuxCarry: true, Sign: true}},
		},
		{
			name: "ADD B (aux carry without sign)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x0F, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x10, B: 0x01, flags: Flags{AuxCarry: true}},
		},
		{
			name: "ADD B (parity)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x02, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, B: 0x01, flags: Flags{Parity: true}},
		},
		{
			name: "ADD B (maximum result without carry)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x7E, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x7F, B: 0x01},
		},
		{
			name: "ADD B (maximum result with carry)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0xFF, B: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFE, B: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Carry: true}},
		},
		{
			name: "ADD C",
			code: `
				ADD C
				HLT
				`,
			initCPU: &CPU{A: 0x01, C: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, C: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD D",
			code: `
				ADD D
				HLT
				`,
			initCPU: &CPU{A: 0x01, D: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, D: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD E",
			code: `
				ADD E
				HLT
				`,
			initCPU: &CPU{A: 0x01, E: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, E: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD H",
			code: `
				ADD H
				HLT
				`,
			initCPU: &CPU{A: 0x01, H: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, H: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD L",
			code: `
				ADD L
				HLT
				`,
			initCPU: &CPU{A: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, L: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD M",
			code: `
				MVI M, 0x55
				ADD M
				HLT
				`,
			initCPU: &CPU{A: 0x01, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0x56, H: 0x01, L: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADD A",
			code: `
				ADD A
				HLT
				`,
			initCPU: &CPU{A: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x04},
		},
		{
			name: "ADC B (carry in with zero flag)",
			code: `
				ADC B
				HLT
			`,
			initCPU: &CPU{A: 0x00, B: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC B (carry in with carry out)",
			code: `
				ADC B
				HLT
			`,
			initCPU: &CPU{A: 0x80, B: 0x80, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, B: 0x80, flags: Flags{Carry: true}},
		},
		{
			name: "ADC B (carry in with parity flag)",
			code: `
				ADC B
				HLT
			`,
			initCPU: &CPU{A: 0x01, B: 0x02, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x04, B: 0x02},
		},
		{
			name: "ADC B (carry in with no carry out + zero result)",
			code: `
				ADC B
				HLT
			`,
			initCPU: &CPU{A: 0xFF, B: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "ADC C (carry in with zero flag)",
			code: `
				ADC C
				HLT
			`,
			initCPU: &CPU{A: 0x00, C: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC D (carry in with zero flag)",
			code: `
				ADC D
				HLT
			`,
			initCPU: &CPU{A: 0x00, D: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC E (carry in with zero flag)",
			code: `
				ADC E
				HLT
			`,
			initCPU: &CPU{A: 0x00, E: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC H (carry in with zero flag)",
			code: `
				ADC H
				HLT
			`,
			initCPU: &CPU{A: 0x00, H: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC L (carry in with zero flag)",
			code: `
				ADC L
				HLT
			`,
			initCPU: &CPU{A: 0x00, L: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADC M (carry in with zero flag)",
			code: `
				MVI M, 0x55
				ADC M
				HLT
				`,
			initCPU: &CPU{A: 0x00, H: 0x01, L: 0x02, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0x56, H: 0x01, L: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "ADC A (carry in with zero flag)",
			code: `
				ADC A
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ADI",
			code: `
				ADI 0x02
				HLT
			`,
			initCPU: &CPU{A: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02},
		},
		{
			name: "ACI",
			code: `
				ACI 0x20
				HLT
			`,
			initCPU: &CPU{A: 0x10, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x31},
		},
		{
			name: "DAD B (basic addition)",
			code: `
				LXI B, 0x1000
				LXI H, 0x2000
				DAD B
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x30, B: 0x00, H: 0x30, L: 0x00},
		},
		{
			name: "DAD B (zero result including carry)",
			code: `
				LXI B, 0x8000
				LXI H, 0x8000
				DAD B
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, H: 0x00, L: 0x00, flags: Flags{Carry: true}},
		},
		{
			name: "DAD B (negative result)",
			code: `
				LXI B, 0xFF00
				LXI H, 0x0100
				DAD B
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, H: 0x00, L: 0x00, flags: Flags{Carry: true}},
		},
		{
			name: "DAD B (large value with overflow)",
			code: `
				LXI B, 0xFFFF
				LXI H, 0xFFFF
				DAD B
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0xFE, C: 0xFF, H: 0xFF, L: 0xFE, flags: Flags{Carry: true}},
		},
		{
			name: "DAD D (basic addition)",
			code: `
				LXI D, 0x1000
				LXI H, 0x2000
				DAD D
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x30, B: 0x00, D: 0x10, H: 0x30, L: 0x00},
		},
		{
			name: "DAD H (basic addition)",
			code: `
				LXI H, 0x2000
				DAD H
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x40, B: 0x00, H: 0x40, L: 0x00},
		},
		{
			name: "DAD H (zero result including carry)",
			code: `
				LXI H, 0x8000
				DAD H
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, H: 0x00, L: 0x00, flags: Flags{Carry: true}},
		},
		{
			name: "DAD H (negative result)",
			code: `
				LXI H, 0x0100
				DAD H
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02, B: 0x00, H: 0x02, L: 0x00},
		},
		{
			name: "DAD H (large value with overflow)",
			code: `
				LXI H, 0xFFFF
				DAD H
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0xFE, H: 0xFF, L: 0xFE, flags: Flags{Carry: true}},
		},
		{
			name: "DAD SP (basic addition)",
			code: `
				LXI SP, 0x1000
				LXI H, 0x2000
				DAD SP
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x30, B: 0x00, H: 0x30, L: 0x00, stackPointer: 0x1000},
		},
		{
			name: "DAD SP (zero result including carry)",
			code: `
				LXI SP, 0x8000
				LXI H, 0x8000
				DAD SP
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, H: 0x00, L: 0x00, flags: Flags{Carry: true}, stackPointer: 0x8000},
		},
		{
			name: "DAD SP (negative result)",
			code: `
				LXI SP, 0xFF00
				LXI H, 0x0100
				DAD SP
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, H: 0x00, L: 0x00, flags: Flags{Carry: true}, stackPointer: 0xFF00},
		},
		{
			name: "DAD SP (large value with overflow)",
			code: `
				LXI SP, 0xFFFF
				LXI H, 0xFFFF
				DAD SP
				MOV A, H
				MOV B, L
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0xFE, H: 0xFF, L: 0xFE, flags: Flags{Carry: true}, stackPointer: 0xFFFF},
		},
		{
			name: "SUB B (no carry or overflow)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x02, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, B: 0x01},
		},
		{
			name: "SUB B (zero result)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x00, B: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "SUB B (carry in on bit eight)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x00, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SUB B (aux carry with sign)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x80, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x7F, B: 0x01, flags: Flags{AuxCarry: true}},
		},
		{
			name: "SUB B (aux carry without sign)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x10, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x0F, B: 0x01, flags: Flags{AuxCarry: true, Parity: true}},
		},
		{
			name: "SUB B (parity)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0x04, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x03, B: 0x01, flags: Flags{Parity: true}},
		},
		{
			name: "SUB B (maximum result without carry)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0xFF, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFE, B: 0x01, flags: Flags{Sign: true}},
		},
		{
			name: "SUB B (maximum result with carry)",
			code: `
				SUB B
				HLT
				`,
			initCPU: &CPU{A: 0xFE, B: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SUB C",
			code: `
				SUB C
				HLT
				`,
			initCPU: &CPU{A: 0x02, C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, C: 0x01},
		},
		{
			name: "SUB D",
			code: `
				SUB D
				HLT
				`,
			initCPU: &CPU{A: 0x02, D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, D: 0x01},
		},
		{
			name: "SUB E",
			code: `
				SUB E
				HLT
				`,
			initCPU: &CPU{A: 0x02, E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, E: 0x01},
		},
		{
			name: "SUB H",
			code: `
				SUB H
				HLT
				`,
			initCPU: &CPU{A: 0x02, H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, H: 0x01},
		},
		{
			name: "SUB L",
			code: `
				SUB L
				HLT
				`,
			initCPU: &CPU{A: 0x02, L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, L: 0x01},
		},
		{
			name: "SUB M",
			code: `
				MVI M, 0x55
				SUB M
				HLT
				`,
			initCPU: &CPU{A: 0x99, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0x44, H: 0x01, L: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "SUB A",
			code: `
				SUB A
				HLT
				`,
			initCPU: &CPU{A: 0x3E, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "SBB B (borrow in with zero flag)",
			code: `
				SBB B
				HLT
			`,
			initCPU: &CPU{A: 0x00, B: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB B (borrow in with carry out)",
			code: `
				SBB B
				HLT
			`,
			initCPU: &CPU{A: 0x80, B: 0x80, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0x80, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB B (borrow in with no carry out + zero result)",
			code: `
				SBB B
				HLT
			`,
			initCPU: &CPU{A: 0x01, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "SBB C (borrow in with zero flag)",
			code: `
				SBB C
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB D (borrow in with zero flag)",
			code: `
				SBB D
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB E (borrow in with zero flag)",
			code: `
				SBB E
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB H (borrow in with zero flag)",
			code: `
				SBB H
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB L (borrow in with zero flag)",
			code: `
				SBB L
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SBB M",
			code: `
				MVI M, 0x55
				SBB M
				HLT
				`,
			initCPU: &CPU{A: 0x01, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0xAB, H: 0x01, L: 0x02, flags: Flags{Sign: true, AuxCarry: true, Carry: true}},
		},
		{
			name: "SBB A (borrow in with zero flag)",
			code: `
				SBB A
				HLT
			`,
			initCPU: &CPU{A: 0x00, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "SUI",
			code: `
				SUI 0x01
				HLT
			`,
			initCPU: &CPU{A: 0x02, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "SBI",
			code: `
				SBI 0x02
				HLT
			`,
			initCPU: &CPU{A: 0x04, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01},
		},
		{
			name: "ANA B (zero result)",
			code: `
				ANA B
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0101, B: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000, B: 0b1010_1010, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ANA B (one bit set)",
			code: `
				ANA B
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, B: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, B: 0b1010_1010},
		},
		{
			name: "ANA B (no bits set)",
			code: `
				ANA B
				HLT
			`,
			initCPU: &CPU{A: 0b1111_1111, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ANA B (all bits set)",
			code: `
				ANA B
				HLT
			`,
			initCPU: &CPU{A: 0b0000_1111, B: 0b1111_1111, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1111, B: 0b1111_1111, flags: Flags{Parity: true}},
		},
		{
			name: "ANA C",
			code: `
				ANA C
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, C: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, C: 0b1010_1010},
		},
		{
			name: "ANA D",
			code: `
				ANA D
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, D: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, D: 0b1010_1010},
		},
		{
			name: "ANA E",
			code: `
				ANA E
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, E: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, E: 0b1010_1010},
		},
		{
			name: "ANA H",
			code: `
				ANA H
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, H: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, H: 0b1010_1010},
		},
		{
			name: "ANA L",
			code: `
				ANA L
				HLT
			`,
			initCPU: &CPU{A: 0b0101_1101, L: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1000, L: 0b1010_1010},
		},
		{
			name: "ANA M",
			code: `
				MVI M, 0x55
				ANA M
				HLT
				`,
			initCPU: &CPU{A: 0xA9, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0b0000_0001, H: 0x01, L: 0x02},
		},
		{
			name: "ANA A",
			code: `
				ANA A
				HLT
			`,
			initCPU: &CPU{A: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1010_1010, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "XRA B (zero result)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0101, B: 0b0101_0101, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000, B: 0b0101_0101, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "XRA B (non-zero result)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0101, B: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, B: 0b1010_1010, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "XRA B (mixed nibbles)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{A: 0b0000_1111, B: 0b1111_0000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, B: 0b1111_0000, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "XRA B (one bit set)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, B: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, B: 0b0000_0100},
		},
		{
			name: "XRA B (no bits set)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "XRA B (all bits set)",
			code: `
				XRA B
				HLT
			`,
			initCPU: &CPU{A: 0b1111_1111, B: 0b1111_1111, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000, B: 0b1111_1111, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "XRA C",
			code: `
				XRA C
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, C: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, C: 0b0000_0100},
		},
		{
			name: "XRA D",
			code: `
				XRA D
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, D: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, D: 0b0000_0100},
		},
		{
			name: "XRA E",
			code: `
				XRA E
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, E: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, E: 0b0000_0100},
		},
		{
			name: "XRA H",
			code: `
				XRA H
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, H: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, H: 0b0000_0100},
		},
		{
			name: "XRA L",
			code: `
				XRA L
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, L: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, L: 0b0000_0100},
		},
		{
			name: "XRA M",
			code: `
				MVI M, 0x55
				XRA M
				HLT
				`,
			initCPU: &CPU{A: 0xA9, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0xFC, H: 0x01, L: 0x02, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "XRA A",
			code: `
				XRA A
				HLT
			`,
			initCPU: &CPU{A: 0b1001_1001, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ORA B (zero result)",
			code: `
				ORA B
				HLT
			`,
			initCPU: &CPU{B: 0b0000_0000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ORA B (non-zero result)",
			code: `
				ORA B
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0101, B: 0b1010_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, B: 0b1010_1010, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "ORA B (mixed nibbles)",
			code: `
				ORA B
				HLT
			`,
			initCPU: &CPU{A: 0b0000_1111, B: 0b1111_0000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, B: 0b1111_0000, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "ORA B (one bit set)",
			code: `
				ORA B
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, B: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, B: 0b0000_0100},
		},
		{
			name: "ORA B (all bits set)",
			code: `
				ORA B
				HLT
			`,
			initCPU: &CPU{A: 0b1111_1111, B: 0b1111_1111, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, B: 0b1111_1111, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "ORA C",
			code: `
				ORA C
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, C: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, C: 0b0000_0100},
		},
		{
			name: "ORA D",
			code: `
				ORA D
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, D: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, D: 0b0000_0100},
		},
		{
			name: "ORA E",
			code: `
				ORA E
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, E: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, E: 0b0000_0100},
		},
		{
			name: "ORA H",
			code: `
				ORA H
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, H: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, H: 0b0000_0100},
		},
		{
			name: "ORA L",
			code: `
				ORA L
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, L: 0b0000_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0100, L: 0b0000_0100},
		},
		{
			name: "ORA M",
			code: `
				MVI M, 0x55
				ORA M
				HLT
				`,
			initCPU: &CPU{A: 0xA9, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0xFD, H: 0x01, L: 0x02, flags: Flags{Sign: true}},
		},
		{
			name: "ORA A",
			code: `
				ORA A
				HLT
			`,
			initCPU: &CPU{A: 0b1001_1001, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1001_1001, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "CMP B (a = b, zero)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x00, B: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "CMP B (a = b, non-zero)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x01, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, B: 0x01, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "CMP B (A > B, positive)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x01, B: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, B: 0x00},
		},
		{
			name: "CMP B (A < B, negative)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x00, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, B: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP B (A = max, nonzero)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0xFF, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xFF, B: 0x01, flags: Flags{Sign: true}},
		},
		{
			name: "CMP B (B = max, nonzero)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x01, B: 0xFF, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, B: 0xFF, flags: Flags{AuxCarry: true, Carry: true}},
		},
		{
			name: "CMP B (A max negative, B max positive)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x80, B: 0x7F, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x80, B: 0x7F, flags: Flags{AuxCarry: true}},
		},
		{
			name: "CMP B (A max positive, B max negative)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x7F, B: 0x80, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x7F, B: 0x80, flags: Flags{Sign: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP B (A mixed, > B)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0xAA, B: 0x55, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xAA, B: 0x55, flags: Flags{Parity: true}},
		},
		{
			name: "CMP B (B mixed, > A)",
			code: `
				CMP B
				HLT
			`,
			initCPU: &CPU{A: 0x55, B: 0xAA, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x55, B: 0xAA, flags: Flags{Sign: true, AuxCarry: true, Carry: true}},
		},
		{
			name: "CMP C",
			code: `
				CMP C
				HLT
			`,
			initCPU: &CPU{A: 0x00, C: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, C: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP D",
			code: `
				CMP D
				HLT
			`,
			initCPU: &CPU{A: 0x00, D: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, D: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP E",
			code: `
				CMP E
				HLT
			`,
			initCPU: &CPU{A: 0x00, E: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, E: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP H",
			code: `
				CMP H
				HLT
			`,
			initCPU: &CPU{A: 0x00, H: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, H: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP L",
			code: `
				CMP L
				HLT
			`,
			initCPU: &CPU{A: 0x00, L: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, L: 0x01, flags: Flags{Sign: true, AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "CMP M",
			code: `
				MVI M, 0x55
				CMP M
				HLT
				`,
			initCPU: &CPU{A: 0xAA, H: 0x01, L: 0x02, Bus: &Memory{Data: make([]byte, 0xFF+4)}},
			wantCPU: &CPU{A: 0xAA, H: 0x01, L: 0x02, flags: Flags{Parity: true}},
		},
		{
			name: "CMP A",
			code: `
				CMP A
				HLT
			`,
			initCPU: &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x01, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ANI",
			code: `
				ANI 55H
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0101_0100},
		},
		{
			name: "XRI",
			code: `
				XRI 55H
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0001},
		},
		{
			name: "ORI",
			code: `
				ORI AAH
				HLT
			`,
			initCPU: &CPU{A: 0b0101_0101, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "CPI",
			code: `
				CPI 55H
				HLT
			`,
			initCPU: &CPU{A: 0xAA, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0xAA, flags: Flags{Parity: true}},
		},
		{
			name: "RLC (with carry)",
			code: `
				RLC
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1101_0101, flags: Flags{Carry: true}},
		},
		{
			name: "RLC (without carry)",
			code: `
				RLC
				HLT
			`,
			initCPU: &CPU{A: 0b0111_0100, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1110_1000},
		},
		{
			name: "RRC (with carry)",
			code: `
				RRC
				HLT
			`,
			initCPU: &CPU{A: 0b0111_0101, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1011_1010, flags: Flags{Carry: true}},
		},
		{
			name: "RRC (without carry)",
			code: `
				RRC
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0111_0101},
		},
		{
			name: "RAL (with carry)",
			code: `
				RAL
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1101_0101, flags: Flags{Carry: true}},
		},
		{
			name: "RAL (without carry)",
			code: `
				RAL
				HLT
			`,
			initCPU: &CPU{A: 0b0111_0100, flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1110_1000},
		},
		{
			name: "RAL (carry propagated)",
			code: `
				RAL
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1101_0101, flags: Flags{Carry: true}},
		},
		{
			name: "RAL (carry not propagated)",
			code: `
				RAL
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1101_0100, flags: Flags{Carry: true}},
		},
		{
			name: "RAR (with carry)",
			code: `
				RAR
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_0101, flags: Flags{Carry: false}},
		},
		{
			name: "RAR (without carry)",
			code: `
				RAR
				HLT
			`,
			initCPU: &CPU{A: 0b0111_0100, flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0011_1010, flags: Flags{Carry: false}},
		},
		{
			name: "RAR (carry propagated)",
			code: `
				RAR
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_0101, flags: Flags{Carry: false}},
		},
		{
			name: "RAR (carry not propagated)",
			code: `
				RAR
				HLT
			`,
			initCPU: &CPU{A: 0b1110_1010, flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0111_0101, flags: Flags{Carry: false}},
		},
		{
			name: "CMA (all zeros)",
			code: `
				CMA
				HLT
			`,
			initCPU: &CPU{A: 0b0000_0000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1111_1111},
		},
		{
			name: "CMA (all ones)",
			code: `
				CMA
				HLT
			`,
			initCPU: &CPU{A: 0b1111_1111, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_0000},
		},
		{
			name: "CMA (mixed)",
			code: `
				CMA
				HLT
			`,
			initCPU: &CPU{A: 0b1011_0101, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0100_1010},
		},
		{
			name: "STC (from 0)",
			code: `
				STC
				HLT
			`,
			initCPU: &CPU{flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Carry: true}},
		},
		{
			name: "STC (from 1)",
			code: `
				STC
				HLT
			`,
			initCPU: &CPU{flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Carry: true}},
		},
		{
			name: "CMC (from 0)",
			code: `
				CMC
				HLT
			`,
			initCPU: &CPU{flags: Flags{Carry: false}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Carry: true}},
		},
		{
			name: "CMC (from 1)",
			code: `
				CMC
				HLT
			`,
			initCPU: &CPU{flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{flags: Flags{Carry: false}},
		},
		{
			name: "DAA - lower nibble less than 9, with auxillary carry",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b0000_1000, flags: Flags{AuxCarry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0000_1110},
		},
		{
			name: "DAA - upper nibble less than 9, with carry",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1000_0000, flags: Flags{Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1110_0000, flags: Flags{Sign: true}},
		},
		{
			name: "DAA - upper and lower nibble less than 9, with auxillary carry and carry",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1000_1000, flags: Flags{AuxCarry: true, Carry: true}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1110_1110, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "DAA - upper and lower nibbles less than 9",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1000_1000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1000_1000, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "DAA - upper and lower nibbles equal to 9",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1001_1001, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b1001_1001, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "DAA - lower nibble greater than 9",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b0000_1011, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0001_0001, flags: Flags{AuxCarry: true, Parity: true}},
		},
		{
			name: "DAA - upper nibble greater than 9",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1011_0000, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0001_0000, flags: Flags{Carry: true}},
		},
		{
			name: "DAA - upper and lower nibbles greater than 9",
			code: `
				DAA
				HLT
			`,
			initCPU: &CPU{A: 0b1011_1011, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0b0010_0001, flags: Flags{AuxCarry: true, Parity: true, Carry: true}},
		},
		{
			name: "DAA - perform addition example from the 8080 manual (2985 + 4936 = 7921)",
			code: `
				MVI A, 29H
				MVI B, 49H
				ADD B
				DAA
				MOV B, A

				MVI A, 85H
				MVI C, 36H
				ADD C
				DAA
				MOV C, A

				MOV A, B
				ADI 01H

				MOV B, A
				HLT
			`,
			initCPU: &CPU{flags: Flags{}, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x79, B: 0x79, C: 0x21, flags: Flags{}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCPU := tc.initCPU
			a := assembler.Assembler{}
			err := a.Assemble(tc.code)
			if (err != nil) != tc.wantErr {
				t.Errorf("Assembler.Assemble() error = %v, wantErr %v", err, tc.wantErr)
			}

			gotCPU.Load(a.ByteCode)
			runErr := gotCPU.Run()
			if runErr != nil {
				t.Errorf("%s", runErr)
			}

			if tc.wantCPU.programCounter != 0 {
				if gotCPU.programCounter != tc.wantCPU.programCounter {
					t.Errorf("%s \ngotProgramCounter  0x%04X,\nwantProgramCounter 0x%04X", tc.name, gotCPU.programCounter, tc.wantCPU.programCounter)
				}
			}

			if !gotCPU.registersEqual(tc.wantCPU) {
				t.Errorf("%s \ngotCPU  %+v,\nwantCPU %+v", tc.name, gotCPU, tc.wantCPU)
			}
		})
	}
}

func (cpu *CPU) registersEqual(otherCPU *CPU) bool {
	if cpu.A == otherCPU.A &&
		cpu.flags == otherCPU.flags &&
		cpu.B == otherCPU.B &&
		cpu.C == otherCPU.C &&
		cpu.D == otherCPU.D &&
		cpu.E == otherCPU.E &&
		cpu.H == otherCPU.H &&
		cpu.L == otherCPU.L &&
		cpu.stackPointer == otherCPU.stackPointer {
		return true
	}

	return false
}

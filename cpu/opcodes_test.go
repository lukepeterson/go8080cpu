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
				MVI M, 55H
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
				MVI M, 55H
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
				MVI M, 55H
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
				MVI M, 55H
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
				MVI B, 55H
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
				MVI C, 55H
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
				MVI D, 55H
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
				MVI E, 55H
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
				MVI A, 02H
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
				MVI M, 55H
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
				LXI B, 0x0201
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{B: 0x01, C: 0x02},
		},
		{
			name: "LXI D",
			code: `
				LXI D, 0x0201
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{D: 0x01, E: 0x02},
		},
		{
			name: "LXI H",
			code: `
				LXI H, 0x0201
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{H: 0x02, L: 0x01},
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
				LXI B, 0101H
				LXI H, 0101H
				MVI M, 55H
				LDAX B
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, B: 0x01, C: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "LDAX D",
			code: `
				LXI D, 0101H
				LXI H, 0101H
				MVI M, 55H
				LDAX D
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, D: 0x01, E: 0x01, H: 0x01, L: 0x01},
		},
		{
			name: "STA",
			code: `
				MVI A, 55H
				STA 0101H
				LXI H 0101H
				MOV B, M
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, B: 0x55, H: 0x01, L: 0x01},
		},
		{
			name: "LDA",
			code: `
				LXI H 0101H
				MVI M, 55H
				LDA 0101H
				HLT
				`,
			initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
			wantCPU: &CPU{A: 0x55, H: 0x01, L: 0x01},
		},
		// TODO: Fix incomplete test
		// {
		// 	name: "SHLD",
		// 	code: `
		// 		LXI H, 4455H
		// 		SHLD 0101H
		// 		HLT
		// 		`,
		// 	initCPU: &CPU{Bus: &Memory{Data: make([]byte, 0xFFFF)}},
		// 	wantCPU: &CPU{H: 0x44, L: 0x55},
		// },
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
				MVI M, 01H
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
				MVI M, 03H
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
			name: "ADD B (zero flag)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x00, B: 0x00, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "ADD B (non zero)",
			code: `
				ADD B
				HLT
				`,
			initCPU: &CPU{A: 0x01, B: 0x01, Bus: &Memory{Data: make([]byte, 32)}},
			wantCPU: &CPU{A: 0x02, B: 0x01},
		},
		{
			name: "ADD B (carry)",
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
				MVI M, 55H
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
				MVI M, 55H
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

func TestCheckCarryOut(t *testing.T) {
	type args struct {
		a     byte
		b     byte
		carry bool
	}
	tests := []struct {
		name         string
		args         args
		wantCarry    bool
		wantAuxCarry bool
	}{
		{
			name:         "no carry, no aux carry (zero case)",
			args:         args{a: 0x00, b: 0x00, carry: NOCARRY},
			wantCarry:    false,
			wantAuxCarry: false,
		},
		{
			name:         "no carry, no aux carry (small numbers)",
			args:         args{a: 0x01, b: 0x01, carry: NOCARRY},
			wantCarry:    false,
			wantAuxCarry: false,
		},
		{
			name:         "no carry, aux carry (lower nibble sum 16)",
			args:         args{a: 0x0F, b: 0x01, carry: NOCARRY},
			wantCarry:    false,
			wantAuxCarry: true,
		},
		{
			name:         "no carry, aux carry (lower nibble sum 17)",
			args:         args{a: 0x08, b: 0x08, carry: WITHCARRY},
			wantCarry:    false,
			wantAuxCarry: true,
		},
		{
			name:         "carry, no aux carry (total sum 256)",
			args:         args{a: 0xF0, b: 0x10, carry: NOCARRY},
			wantCarry:    true,
			wantAuxCarry: false,
		},
		{
			name:         "carry, no aux carry (total sum 257)",
			args:         args{a: 0x80, b: 0x80, carry: WITHCARRY},
			wantCarry:    true,
			wantAuxCarry: false,
		},
		{
			name:         "carry, aux carry (total sum 256 + lower nibble sum 16)",
			args:         args{a: 0xFF, b: 0x01, carry: NOCARRY},
			wantCarry:    true,
			wantAuxCarry: true,
		},
		{
			name:         "carry, aux carry (total sum 255 + 1)",
			args:         args{a: 0x7F, b: 0x7F, carry: WITHCARRY},
			wantCarry:    false,
			wantAuxCarry: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCarry, gotAuxCarry := checkCarryOut(tt.args.a, tt.args.b, tt.args.carry)
			if gotCarry != tt.wantCarry {
				t.Errorf("carry() gotCarry = %v, want %v", gotCarry, tt.wantCarry)
			}
			if gotAuxCarry != tt.wantAuxCarry {
				t.Errorf("carry() gotAuxCarry = %v, want %v", gotAuxCarry, tt.wantAuxCarry)
			}
		})
	}
}

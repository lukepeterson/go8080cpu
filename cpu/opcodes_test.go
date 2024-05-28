package cpu

import (
	"testing"

	"github.com/lukepeterson/go8080assembler/assembler"
)

func TestCPUIncrementDecrement(t *testing.T) {
	testCases := []struct {
		name    string
		code    string
		initCPU func() *CPU
		wantCPU *CPU
		wantErr bool
	}{
		{
			name: "INR A from 0x01",
			code: `
				INR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x02},
		},
		{
			name: "DCR A from 0x03",
			code: `
				DCR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x02},
		},
		{
			name: "INR A from 0x02 (test parity flag set)",
			code: `
				INR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x02, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}},
		},
		{
			name: "DCR A from 0x04 (test parity flag set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x04, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x03, flags: Flags{Parity: true}},
		},
		{
			name: "INR A from 0x7F (test sign flag set)",
			code: `
				INR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x7F, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x80, flags: Flags{AuxCarry: true, Sign: true}},
		},
		{
			name: "DCR A from 0x81 (test sign flag set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x81, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x80, flags: Flags{Sign: true}},
		},
		{
			name: "INR A from 0x80 (test sign and parity flags set)",
			code: `
				INR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x80, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x81, flags: Flags{Sign: true, Parity: true}},
		},
		{
			name: "DCR A from 0x82 (test sign and parity flags set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x82, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x81, flags: Flags{Sign: true, Parity: true}},
		},

		{
			name: "INR A from 0xFF (test zero and parity flags set)",
			code: `
				INR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0xFF, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, AuxCarry: true, Parity: true}},
		},
		{
			name: "DCR A from 0x01 (test zero and parity flags set)",
			code: `
				DCR A
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{A: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x00, flags: Flags{Zero: true, Parity: true}},
		},
		{
			name: "INR B from 0x01",
			code: `
				INR B
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{B: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{B: 0x02},
		},
		{
			name: "DCR B from 0x03",
			code: `
				DCR B
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{B: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{B: 0x02},
		},
		{
			name: "INR C from 0x01",
			code: `
				INR C
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{C: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{C: 0x02},
		},
		{
			name: "DCR C from 0x03",
			code: `
				DCR C
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{C: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{C: 0x02},
		},
		{
			name: "INR D from 0x01",
			code: `
				INR D
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{D: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{D: 0x02},
		},
		{
			name: "DCR D from 0x03",
			code: `
				DCR D
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{D: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{D: 0x02},
		},
		{
			name: "INR E from 0x01",
			code: `
				INR E
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{E: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{E: 0x02},
		},
		{
			name: "DCR E from 0x03",
			code: `
				DCR E
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{E: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{E: 0x02},
		},
		{
			name: "INR H from 0x01",
			code: `
				INR H
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{H: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{H: 0x02},
		},
		{
			name: "DCR H from 0x03",
			code: `
				DCR H
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{H: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{H: 0x02},
		},
		{
			name: "INR L from 0x01",
			code: `
				INR L
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{L: 0x01, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{L: 0x02},
		},
		{
			name: "DCR L from 0x03",
			code: `
				DCR L
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{L: 0x03, Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{L: 0x02},
		},
		{
			name: "INR M from 0x01",
			code: `
				LXI H, 16H
				MVI M, 01H
				INR M
				MOV A, M
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x02, L: 0x16},
		},
		{
			name: "DCR M from 0x03",
			code: `
				LXI H, 16H
				MVI M, 03H
				DCR M
				MOV A, M
				HLT
				`,
			initCPU: func() *CPU {
				return &CPU{Bus: &Memory{Data: make([]byte, 32)}}
			}, wantCPU: &CPU{A: 0x02, L: 0x16},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCPU := tc.initCPU()
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

func (cpu *CPU) registersEqual(other *CPU) bool {
	if cpu.A == other.A &&
		cpu.flags == other.flags &&
		cpu.B == other.B &&
		cpu.C == other.C &&
		cpu.D == other.D &&
		cpu.E == other.E &&
		cpu.H == other.H &&
		cpu.L == other.L &&
		cpu.stackPointer == other.stackPointer {
		return true
	}

	return false
}

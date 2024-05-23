package main

import (
	"time"

	"github.com/lukepeterson/gocpu/cpu"
)

func main() {

	goCPU := cpu.NewCPU(
		1*time.Millisecond,
		cpu.NewMemory(32),
	)

	program := map[uint16]byte{
		0x00: 0x3A, // LDA
		0x01: 0x10,
		0x02: 0x00,
		0x03: 0x4C, // HLT
		0x10: 0x69, // Data
	}

	goCPU.Load(program)

	// goCPU.DumpRegisters()
	// goCPU.DumpMemory()
	goCPU.Run()
	goCPU.DumpRegisters()
	goCPU.DumpMemory()
}

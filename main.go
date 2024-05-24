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
		0x00: 0x3E, // MVI, A
		0x01: 0xFE, // DATA
		0x03: 0x3C, // INR A
		0x04: 0x4C, // HLT
	}

	goCPU.Load(program)

	goCPU.DumpRegisters()
	goCPU.DumpMemory()
	goCPU.Run()
	goCPU.DumpRegisters()
	goCPU.DumpMemory()
}

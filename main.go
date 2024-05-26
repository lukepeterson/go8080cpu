package main

import (
	"time"

	"github.com/lukepeterson/gocpu/cpu"
)

func main() {

	goCPU := cpu.NewCPU(
		1000*time.Millisecond,
		cpu.NewMemory(32),
	)

	program := map[uint16]byte{
		0x00: 0x3E, // MVI, A
		0x01: 0x02, // DATA
		0x02: 0x3C, // INR A
		0x03: 0xC3, // JMP
		0x04: 0x02,
		0x05: 0x00,
		0x06: 0x4C, // HLT
	}

	goCPU.Load(program)

	goCPU.Run()
}

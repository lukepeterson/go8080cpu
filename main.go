package main

import (
	"time"

	"github.com/lukepeterson/gocpu/cpu"
)

func main() {

	memory := map[uint16]byte{
		0x00: 0x3A, // LDA
		0x01: 0x10,
		0x02: 0x00,
		0x03: 0x4C, // HLT
		0x10: 0x69, // Data
		0x11: 0x35,
		0x12: 0x9A,
		0x13: 0x7D,
	}

	goCPU := cpu.NewCPU(1*time.Millisecond, 32)
	for addr, value := range memory {
		goCPU.Memory[addr] = value
	}

	goCPU.DumpRegisters()
	goCPU.DumpMemory()
	goCPU.Run()
	goCPU.DumpRegisters()
	goCPU.DumpMemory()
}

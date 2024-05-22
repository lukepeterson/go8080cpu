package main

import (
	"fmt"

	"github.com/lukepeterson/gocpu/cpu"
)

func main() {

	goCPU := cpu.NewCPU()

	goCPU.Memory[0x00] = 0x00 // NOP
	goCPU.Memory[0x01] = 0x00 // NOP
	goCPU.Memory[0x02] = 0x4C // HLT

	goCPU.Memory[0x10] = 0x01
	goCPU.Memory[0x11] = 0x02
	goCPU.Memory[0x12] = 0x03
	goCPU.Memory[0x13] = 0x04

	goCPU.DumpMemory()
	goCPU.Run()
	goCPU.DumpMemory()

	bytes := []byte{0x08, 0x13, 0x64}
	fmt.Printf("bytes: %02X\n", bytes)
	fmt.Printf("bytes: %08b\n", bytes)

}

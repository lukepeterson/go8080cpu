package cpu

import (
	"fmt"
	"log"
	"time"
)

type CPU struct {
	// alu            ALU
	A, flags       byte
	B, C           byte
	D, E           byte
	H, L           byte
	stackPointer   byte
	programCounter byte

	Memory        [32]byte
	fetchedOpcode byte
	halted        bool
	delay         time.Duration
}

func (cpu *CPU) DumpMemory() {
	for i := range cpu.Memory {
		fmt.Printf("%02X ", cpu.Memory[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}

func NewCPU(delay time.Duration) *CPU {
	initOpcodeMap()
	return &CPU{
		delay: delay,
	}
}

func (cpu *CPU) Fetch() {
	cpu.fetchedOpcode = cpu.Memory[cpu.programCounter]
	cpu.programCounter++
}

func (cpu *CPU) Decode() opcodeFunc {
	return opcodeMap[cpu.fetchedOpcode]
}

func (cpu *CPU) Execute(instruction opcodeFunc) {
	if instruction != nil {
		instruction(cpu)
	} else {
		cpu.halted = true
		log.Fatal("instruction not found")
	}
}

func (cpu *CPU) Run() {
	for !cpu.halted {
		time.Sleep(cpu.delay)
		cpu.Fetch()
		instruction := cpu.Decode()
		cpu.Execute(instruction)
	}
}

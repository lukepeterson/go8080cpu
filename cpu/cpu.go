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
	stackPointer   uint16
	programCounter uint16

	Memory        []byte
	fetchedOpcode byte
	halted        bool
	delay         time.Duration
}

func (cpu *CPU) DumpRegisters() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Registers:")
	fmt.Printf("     A: %08b (%02X), F: %08b (%02X)\n", cpu.A, cpu.A, cpu.flags, cpu.flags)
	fmt.Printf("     B: %08b (%02X), C: %08b (%02X)\n", cpu.B, cpu.B, cpu.C, cpu.C)
	fmt.Printf("     D: %08b (%02X), E: %08b (%02X)\n", cpu.D, cpu.D, cpu.E, cpu.E)
	fmt.Printf("     H: %08b (%02X), L: %08b (%02X)\n", cpu.H, cpu.H, cpu.L, cpu.L)
	fmt.Printf("    PC: %016b (%04X)\n", cpu.programCounter, cpu.programCounter)
	fmt.Printf("    SP: %016b (%04X)\n", cpu.stackPointer, cpu.stackPointer)
}

func (cpu *CPU) DumpMemory() {
	fmt.Println("Memory: ")
	fmt.Print("    ")
	for i := range cpu.Memory {
		fmt.Printf("%02X ", cpu.Memory[i])
		if (i+1)%8 == 0 {
			fmt.Print("\n    ")
		}
	}
	fmt.Println("-----------------------------------------")
}

func NewCPU(delay time.Duration, memorySize int) *CPU {
	initOpcodeMap()
	return &CPU{
		delay:  delay,
		Memory: make([]byte, memorySize),
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

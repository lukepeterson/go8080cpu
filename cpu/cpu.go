package cpu

import (
	"fmt"
	"log"
	"time"
)

type Register struct {
	Data [8]bool
}

func (r *Register) Load(input byte) []bool {
	data := make([]bool, 8)
	for i := 7; i >= 0; i-- {
		data[7-i] = (input>>i)&1 == 1
	}

	return data
}

type opcodeFunc func(*CPU)

var opcodeMap [256]opcodeFunc

func (cpu *CPU) NOP() {
	fmt.Println("NOP")
}

func (cpu *CPU) HLT() {
	fmt.Println("HALT!")
	cpu.halted = true
}

func initOpcodeMap() {
	opcodeMap[0x00] = (*CPU).NOP
	opcodeMap[0x4C] = (*CPU).HLT
}

// var decodeMnemonic = map[string]byte{
// 	"NOP": 0x00,
// 	"LDA": 0x3A,
// 	"HLT": 0x4C,
// }

type Bus struct {
	Data [8]bool
}

func NewBus() *Bus {
	return &Bus{
		Data: [8]bool{},
	}
}

type CPU struct {
	// alu            ALU
	A Register
	B Register

	fetchedOpcode  byte
	Memory         [32]byte
	programCounter byte
	bus            *Bus
	halted         bool
}

func (cpu *CPU) DumpMemory() {
	for i := range cpu.Memory {
		fmt.Printf("%02X ", cpu.Memory[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}

func NewCPU() *CPU {
	cpu := CPU{
		bus: NewBus(),
	}

	initOpcodeMap()

	return &cpu
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
		time.Sleep(10 * time.Millisecond)
		cpu.Fetch()
		instruction := cpu.Decode()
		cpu.Execute(instruction)
	}

}

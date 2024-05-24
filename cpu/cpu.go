package cpu

import (
	"fmt"
	"log"
	"time"
)

type word uint16

type Bus interface {
	ReadByte(address word) byte
	WriteByte(address word, data byte)
	length() uint16
}

type Memory struct {
	Data []byte
}

func (memory Memory) ReadByte(address word) byte {
	return memory.Data[address]
}

func (memory *Memory) WriteByte(address word, data byte) {
	memory.Data[address] = data
}

func (memory Memory) length() uint16 {
	return uint16(len(memory.Data))
}

func (cpu *CPU) fetchByte() byte {
	nextByte := cpu.Bus.ReadByte(cpu.programCounter)
	cpu.programCounter++
	return nextByte
}

func (cpu *CPU) fetchWord() word {
	low := cpu.fetchByte()
	high := cpu.fetchByte()
	return joinBytes(low, high) // 8080 is little endian
}

func joinBytes(low, high byte) word {
	return word(low) | word(high)<<8
}

type Flags struct {
	Sign     bool
	Zero     bool
	AuxCarry bool
	Parity   bool
	Carry    bool
}

type CPU struct {
	A    byte
	B, C byte
	D, E byte
	H, L byte

	flags Flags

	stackPointer   word
	programCounter word

	// alu    ALU
	Bus    Bus
	halted bool
	Delay  time.Duration
}

func (cpu CPU) DumpRegisters() {
	fmt.Println("-----------------------------------------")
	fmt.Println("Registers:")
	fmt.Printf("     A: %08b (%02X), S:%v Z:%v AC:%v P:%v C:%v\n", cpu.A, cpu.A, boolToInt(cpu.flags.Sign), boolToInt(cpu.flags.Zero), boolToInt(cpu.flags.AuxCarry), boolToInt(cpu.flags.AuxCarry), boolToInt(cpu.flags.Carry))
	fmt.Printf("     B: %08b (%02X), C: %08b (%02X)\n", cpu.B, cpu.B, cpu.C, cpu.C)
	fmt.Printf("     D: %08b (%02X), E: %08b (%02X)\n", cpu.D, cpu.D, cpu.E, cpu.E)
	fmt.Printf("     H: %08b (%02X), L: %08b (%02X)\n", cpu.H, cpu.H, cpu.L, cpu.L)
	fmt.Printf("    PC: %016b (%04X)\n", cpu.programCounter, cpu.programCounter)
	fmt.Printf("    SP: %016b (%04X)\n", cpu.stackPointer, cpu.stackPointer)
}

func (cpu *CPU) DumpMemory() {
	fmt.Printf("Memory: (%v bytes)\n", cpu.Bus.length())
	fmt.Print("    ")
	for i := 0; i < int(cpu.Bus.length()); i++ {
		fmt.Printf("%02X ", cpu.Bus.ReadByte(word(i)))
		if (i+1)%8 == 0 {
			fmt.Print("\n    ")
		}
	}
	fmt.Println("-------------------------------------")
}

func boolToInt(in bool) int {
	if in {
		return 1
	}

	return 0
}

func getParity(b byte) bool {
	var count int
	for b != 0 {
		count += int(b & 1)
		b >>= 1
	}
	return count%2 == 0
}

func NewMemory(size uint16) *Memory {
	return &Memory{
		Data: make([]byte, size),
	}
}

func NewCPU(delay time.Duration, memory *Memory) *CPU {
	return &CPU{
		Delay: delay,
		Bus:   memory,
	}
}

func (cpu *CPU) Execute(instruction opcodeFunc) {
	if instruction != nil {
		instruction(cpu)
	} else {
		cpu.halted = true
		log.Println("instruction not found")
	}
}

func (cpu *CPU) Run() {
	initOpcodeMap()
	for !cpu.halted {
		// cpu.DumpRegisters()
		// cpu.DumpMemory()

		opCode := cpu.fetchByte()     // FETCH
		instruction := decode[opCode] // DECODE
		cpu.Execute(instruction)      // EXECUTE
		time.Sleep(cpu.Delay)
	}
}

func (cpu *CPU) Load(data map[uint16]byte) {
	for addr, value := range data {
		cpu.Bus.WriteByte(word(addr), value)
	}
}

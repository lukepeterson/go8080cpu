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
	getLength() uint16
}

type Memory struct {
	data []byte
}

func (memory Memory) getLength() uint16 {
	return uint16(len(memory.data))
}

func (memory Memory) ReadByte(address word) byte {
	return memory.data[address]
}

func (memory *Memory) WriteByte(address word, data byte) {
	memory.data[address] = data
}

func (cpu *CPU) fetchByte() byte {
	nextByte := cpu.bus.ReadByte(cpu.programCounter)
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

type CPU struct {
	// alu            ALU
	A, flags       byte
	B, C           byte
	D, E           byte
	H, L           byte
	stackPointer   word
	programCounter word

	bus    Bus
	halted bool
	delay  time.Duration
}

func (cpu CPU) DumpRegisters() {
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
	fmt.Printf("Memory: (%v bytes)\n", cpu.bus.getLength())
	fmt.Print("    ")
	for i := 0; i < int(cpu.bus.getLength()); i++ {
		fmt.Printf("%02X ", cpu.bus.ReadByte(word(i)))
		if (i+1)%8 == 0 {
			fmt.Print("\n    ")
		}
	}
	fmt.Println("-------------------------------------")
}

func NewMemory(size uint16) *Memory {
	return &Memory{
		data: make([]byte, size),
	}
}

func NewCPU(delay time.Duration, memory *Memory) *CPU {
	initOpcodeMap()
	return &CPU{
		delay: delay,
		bus:   memory,
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
	for !cpu.halted {
		opCode := cpu.fetchByte()     // FETCH
		instruction := decode[opCode] // DECODE
		cpu.Execute(instruction)      // EXECUTE
		time.Sleep(cpu.delay)
	}
}

func (cpu *CPU) Load(data map[uint16]byte) {
	for addr, value := range data {
		cpu.bus.WriteByte(word(addr), value)
	}
}

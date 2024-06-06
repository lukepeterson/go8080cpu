package cpu

import (
	"fmt"
	"strings"
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
	low := cpu.fetchByte() // 8080 is little endian, so low byte comes first when reading from memory
	high := cpu.fetchByte()
	return joinBytes(high, low)
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

	Bus    Bus
	halted bool
}

func (cpu CPU) DumpRegisters() {
	var sb strings.Builder
	// sb.WriteString("\033[H\033[2J") // Clear the screen and move top, left
	sb.WriteString("-----------------------------------------\n")
	sb.WriteString("Registers:\n")
	sb.WriteString(fmt.Sprintf("     A: %08b (%02X), S:%v Z:%v AC:%v P:%v C:%v\n", cpu.A, cpu.A, boolToInt(cpu.flags.Sign), boolToInt(cpu.flags.Zero), boolToInt(cpu.flags.AuxCarry), boolToInt(cpu.flags.Parity), boolToInt(cpu.flags.Carry)))
	sb.WriteString(fmt.Sprintf("     B: %08b (%02X), C: %08b (%02X)\n", cpu.B, cpu.B, cpu.C, cpu.C))
	sb.WriteString(fmt.Sprintf("     D: %08b (%02X), E: %08b (%02X)\n", cpu.D, cpu.D, cpu.E, cpu.E))
	sb.WriteString(fmt.Sprintf("     H: %08b (%02X), L: %08b (%02X)\n", cpu.H, cpu.H, cpu.L, cpu.L))
	sb.WriteString(fmt.Sprintf("    PC: %016b (%04X)\n", cpu.programCounter, cpu.programCounter))
	sb.WriteString(fmt.Sprintf("    SP: %016b (%04X)\n", cpu.stackPointer, cpu.stackPointer))
	fmt.Print(sb.String())
}

func (cpu *CPU) DumpMemory(startAddress, endAddress word) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Memory: %v bytes, ", cpu.Bus.length()))
	sb.WriteString(fmt.Sprintf("Start: 0x%0004X, End: 0x%0004X\n", startAddress, endAddress))
	sb.WriteString("    ")
	for i := startAddress; i < endAddress; i++ {
		sb.WriteString(fmt.Sprintf("%02X ", cpu.Bus.ReadByte(word(i))))
		if (i+1)%16 == 0 {
			sb.WriteString("\n    ")
		}
	}
	sb.WriteString("\n-------------------------------------\n")
	fmt.Print(sb.String())
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

func NewCPU(memory *Memory) *CPU {
	return &CPU{
		Bus: memory,
	}
}

func (cpu *CPU) Run() error {
	for !cpu.halted {
		err := cpu.Execute(cpu.fetchByte())
		if err != nil {
			return err
		}

		// cpu.DumpRegisters()
		// cpu.DumpMemory(0x2000, 0x2000+32)
	}

	return nil
}

func (cpu *CPU) Load(data []byte) {
	for addr, value := range data {
		cpu.Bus.WriteByte(word(addr), value)
	}
}

package cpu

import (
	"fmt"
	"strings"
)

type word uint16

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
	sb.WriteString(fmt.Sprintf("     A: %08b (0x%02X), S:%v Z:%v AC:%v P:%v C:%v\n", cpu.A, cpu.A, boolToInt(cpu.flags.Sign), boolToInt(cpu.flags.Zero), boolToInt(cpu.flags.AuxCarry), boolToInt(cpu.flags.Parity), boolToInt(cpu.flags.Carry)))
	sb.WriteString(fmt.Sprintf("     B: %08b (0x%02X), C: %08b (0x%02X)\n", cpu.B, cpu.B, cpu.C, cpu.C))
	sb.WriteString(fmt.Sprintf("     D: %08b (0x%02X), E: %08b (0x%02X)\n", cpu.D, cpu.D, cpu.E, cpu.E))
	sb.WriteString(fmt.Sprintf("     H: %08b (0x%02X), L: %08b (0x%02X)\n", cpu.H, cpu.H, cpu.L, cpu.L))
	sb.WriteString(fmt.Sprintf("    PC: %016b (0x%04X)\n", cpu.programCounter, cpu.programCounter))
	sb.WriteString(fmt.Sprintf("    SP: %016b (0x%04X)\n", cpu.stackPointer, cpu.stackPointer))
	fmt.Print(sb.String())
}

func (cpu *CPU) DumpMemory(startAddress, endAddress word) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Memory: %v bytes, ", cpu.Bus.length()))
	sb.WriteString(fmt.Sprintf("Start: 0x%0004X, End: 0x%0004X\n", startAddress, endAddress))
	sb.WriteString("    ")
	for i := startAddress; i < endAddress; i++ {
		readByte, err := cpu.Bus.ReadByteAt(word(i))
		if err != nil {
			return fmt.Errorf("could not read byte 0x%02X at address 0x%04X: %v", readByte, word(i), err)
		}

		sb.WriteString(fmt.Sprintf("%02X ", readByte))
		if (i+1)%16 == 0 {
			sb.WriteString("\n    ")
		}
	}
	sb.WriteString("\n-------------------------------------\n")
	fmt.Print(sb.String())

	return nil
}

func NewCPU(memory *Memory) *CPU {
	return &CPU{
		Bus: memory,
	}
}

func (cpu *CPU) Load(data []byte) error {
	for addr, value := range data {
		err := cpu.Bus.WriteByteAt(word(addr), value)
		if err != nil {
			return fmt.Errorf("could not write byte 0x%02X at address 0x%04X: %v", value, word(addr), err)
		}
	}
	return nil
}

func (cpu *CPU) Run() error {
	for !cpu.halted {
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return fmt.Errorf("could not fetch byte %v", err)
		}

		err = cpu.Execute(fetchedByte)
		if err != nil {
			return fmt.Errorf("could not execute fetchedByte 0x%02X: %v", fetchedByte, err)
		}

		// cpu.DumpRegisters()
		// cpu.DumpMemory(0x00, word(cpu.Bus.length()))
	}

	return nil
}

func (cpu *CPU) fetchByte() (byte, error) {
	readByte, err := cpu.Bus.ReadByteAt(cpu.programCounter)
	if err != nil {
		return 0, fmt.Errorf("could not read byte at 0x%04X: %v", cpu.programCounter, err)
	}

	cpu.programCounter++
	return readByte, nil
}

func (cpu *CPU) fetchWord() (word, error) {
	low, err := cpu.fetchByte() // 8080 is little endian, so low byte comes first when reading from memory
	if err != nil {
		return 0, fmt.Errorf("could not fetch low byte of fetchWord: %v", err)
	}

	high, err := cpu.fetchByte()
	if err != nil {
		return 0, fmt.Errorf("could not fetch high byte of fetchWord: %v", err)
	}

	return joinBytes(high, low), nil
}

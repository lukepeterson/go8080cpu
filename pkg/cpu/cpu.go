package cpu

import (
	"fmt"

	"github.com/lukepeterson/go8080cpu/pkg/memory"
	"github.com/lukepeterson/go8080cpu/pkg/types"
)

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

	stackPointer   types.Word
	programCounter types.Word

	interruptEnabled     bool
	interruptPending     bool
	interruptInstruction byte

	Bus       Bus
	halted    bool
	DebugMode bool
}

type Bus interface {
	ReadByteAt(address types.Word) (byte, error)
	WriteByteAt(address types.Word, data byte) error
}

func New() *CPU {
	return &CPU{Bus: memory.New()}
}

func (cpu *CPU) Load(data []byte) error {
	for addr, value := range data {
		err := cpu.Bus.WriteByteAt(types.Word(addr), value)
		if err != nil {
			return fmt.Errorf("could not write byte 0x%02X at address 0x%04X: %v", value, types.Word(addr), err)
		}
	}

	return nil
}

func (cpu *CPU) Run() error {
	for !cpu.halted {
		var nextInstruction byte
		var err error
		if cpu.interruptEnabled && cpu.interruptPending {
			cpu.interruptEnabled = false
			cpu.interruptPending = false
			nextInstruction = cpu.interruptInstruction
		} else {
			nextInstruction, err = cpu.fetchByte()
			if err != nil {
				return fmt.Errorf("could not fetch byte: %v", err)
			}
		}

		err = cpu.Execute(nextInstruction)
		if err != nil {
			return fmt.Errorf("could not execute nextInstruction 0x%02X: %v", nextInstruction, err)
		}

		if cpu.DebugMode {
			cpu.DumpRegisters()
			cpu.DumpMemory(0x0000, 0x0020) // Start of program code
			cpu.DumpMemory(0xFFDF, 0xFFFF) // End of stack
		}
	}

	return nil
}

func (cpu *CPU) fetchByte() (byte, error) {
	readByte, err := cpu.Bus.ReadByteAt(cpu.programCounter)
	if err != nil {
		return 0, fmt.Errorf("could not fetch byte at 0x%04X: %v", cpu.programCounter, err)
	}

	cpu.programCounter++
	return readByte, nil
}

func (cpu *CPU) fetchWord() (types.Word, error) {
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

// getFlags returns the current state of the CPU flags packed into a single byte.
// The flags are ordered from MSB (bit 7) to LSB (bit 0).
//
// This method performs the following steps:
// 1. Generates a slice of eight bools for the flag storage
// 2. Iterates through each bit in the slice, shifting the bits to the left if set
//
// Example:
//
//	cpu := &CPU{flags: Flags{Sign: true, Parity: true}}
//	result := cpu.getFlags()
//	// result is 0b10000110 (0x86 or 134)
func (cpu CPU) getFlags() byte {
	flags := []bool{
		cpu.flags.Sign,
		cpu.flags.Zero,
		false, // Bit 5 is always false
		cpu.flags.AuxCarry,
		false, // Bit 3 is always false
		cpu.flags.Parity,
		true, // Bit 1 is always true
		cpu.flags.Carry,
	}

	// Convert the five boolean values stored in cpu.flags into an 8-bit byte
	var result byte
	for i, flag := range flags {
		if flag {
			result |= 1 << (7 - i)
		}
	}

	return result
}

// setFlags updates the CPU flags based on the provided flags byte.
//
// Example:
//
//	cpu.setFlags(0b10010110) // 0x96
//	// cpu.flags = Flags{Sign: true, Zero: false, AuxCarry: true, Parity: true, Carry: false}
func (cpu *CPU) setFlags(flags byte) {
	cpu.flags.Sign = (flags & (1 << 7)) != 0
	cpu.flags.Zero = (flags & (1 << 6)) != 0
	// Bit 5 is always false
	cpu.flags.AuxCarry = (flags & (1 << 4)) != 0
	// Bit 3 is always false
	cpu.flags.Parity = (flags & (1 << 2)) != 0
	// Bit 1 is always true
	cpu.flags.Carry = (flags & (1 << 0)) != 0
}

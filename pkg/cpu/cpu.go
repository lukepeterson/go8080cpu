package cpu

import (
	"fmt"

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

	Bus       Bus
	halted    bool
	DebugMode bool
}

type Bus interface {
	ReadByteAt(address types.Word) (byte, error)
	WriteByteAt(address types.Word, data byte) error
	Length() uint16
}

func New() *CPU {
	return &CPU{}
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
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return fmt.Errorf("could not fetch byte: %v", err)
		}

		err = cpu.Execute(fetchedByte)
		if err != nil {
			return fmt.Errorf("could not execute fetchedByte 0x%02X: %v", fetchedByte, err)
		}

		if cpu.DebugMode {
			cpu.DumpRegisters()
			cpu.DumpMemory(0x00, types.Word(cpu.Bus.Length()))
		}
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

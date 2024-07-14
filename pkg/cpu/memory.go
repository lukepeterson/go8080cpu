package cpu

import "fmt"

type Bus interface {
	ReadByteAt(address word) (byte, error)
	WriteByteAt(address word, data byte) error
	length() uint16
}

type Memory struct {
	Data []byte
}

func NewMemory(size uint16) *Memory {
	return &Memory{
		Data: make([]byte, size),
	}
}

func (cpu *CPU) Load(data []byte) error {
	for addr, value := range data {
		err := cpu.Bus.WriteByteAt(word(addr), value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (memory Memory) ReadByteAt(address word) (byte, error) {
	if int(address) >= len(memory.Data) {
		return 0, fmt.Errorf("could not read from address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}
	return memory.Data[address], nil
}

func (memory *Memory) WriteByteAt(address word, data byte) error {
	if int(address) >= len(memory.Data) {
		return fmt.Errorf("could not write to address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}
	memory.Data[address] = data
	return nil
}

func (memory Memory) length() uint16 {
	return uint16(len(memory.Data))
}

func (cpu *CPU) fetchByte() (byte, error) {
	readByte, err := cpu.Bus.ReadByteAt(cpu.programCounter)
	if err != nil {
		return 0, err
	}

	cpu.programCounter++
	return readByte, nil
}

func (cpu *CPU) fetchWord() (word, error) {
	low, err := cpu.fetchByte() // 8080 is little endian, so low byte comes first when reading from memory
	if err != nil {
		return 0, err
	}

	high, err := cpu.fetchByte()
	if err != nil {
		return 0, err
	}

	return joinBytes(high, low), nil
}

func boolToInt(in bool) int {
	if in {
		return 1
	}

	return 0
}

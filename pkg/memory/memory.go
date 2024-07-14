package memory

import (
	"fmt"

	"github.com/lukepeterson/go8080cpu/pkg/types"
)

type Memory struct {
	Data []byte
}

func New(size uint) *Memory {
	return &Memory{
		Data: make([]byte, size),
	}
}

func (memory Memory) ReadByteAt(address types.Word) (byte, error) {
	if int(address) >= len(memory.Data) {
		return 0, fmt.Errorf("could not read from address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}
	return memory.Data[address], nil
}

func (memory *Memory) WriteByteAt(address types.Word, data byte) error {
	if int(address) >= len(memory.Data) {
		return fmt.Errorf("could not write to address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}
	memory.Data[address] = data
	return nil
}

func (memory Memory) Length() uint16 {
	return uint16(len(memory.Data))
}

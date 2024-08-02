package memory

import (
	"fmt"
	"math"

	"github.com/lukepeterson/go8080cpu/pkg/types"
)

type Memory struct {
	Data []byte
}

func New() *Memory {
	return &Memory{
		// The 8080 CPU has 64KB of memory, so 65,536 bytes starting from zero,
		// accessed via memory locations 0x0000 to 0xFFFF.
		Data: make([]byte, math.MaxUint16),
	}
}

// ReadByteAt reads a byte from the specified memory location
func (memory Memory) ReadByteAt(address types.Word) (byte, error) {
	if int(address) >= len(memory.Data) {
		return 0, fmt.Errorf("could not read from address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}

	return memory.Data[address], nil
}

// WriteByteTo writes a byte to the specified memory location
func (memory *Memory) WriteByteAt(address types.Word, data byte) error {
	if int(address) >= len(memory.Data) {
		return fmt.Errorf("could not write to address 0x%04X (out of bounds as memory size is 0x%04X)", address, len(memory.Data))
	}

	memory.Data[address] = data
	return nil
}

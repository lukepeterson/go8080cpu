package cpu

import (
	"fmt"
	"math"
	"strings"

	"github.com/lukepeterson/go8080cpu/pkg/types"
)

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

func (cpu *CPU) DumpMemory(startAddress, endAddress types.Word) error {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Memory: %v bytes, ", math.MaxUint16))
	sb.WriteString(fmt.Sprintf("Start: 0x%0004X, End: 0x%0004X\n", startAddress, endAddress))
	sb.WriteString("    ")
	for i := startAddress; i < endAddress; i++ {
		readByte, err := cpu.Bus.ReadByteAt(types.Word(i))
		if err != nil {
			return fmt.Errorf("could not read byte 0x%02X at address 0x%04X: %v", readByte, types.Word(i), err)
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

func boolToInt(in bool) int {
	if in {
		return 1
	}

	return 0
}

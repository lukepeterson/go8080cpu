package cpu

import "fmt"

type opcodeFunc func(*CPU)

var decode [256]opcodeFunc

// initOpcodeMap is a mapping of Intel 8080 opcodes to their functions.
// Grouped by instruction set group as per "Table 2. Instruction Set Summary",
// in the Intel 8080A 8-BIT N-CHANNEL MICROPROCESSOR datasheet.
func initOpcodeMap() {
	// MOVE, LOAD AND STORE
	decode[0x3A] = (*CPU).LDA

	// STACK OPERATIONS

	// JUMP

	// RESTART

	// INCREMENT AND DECREMENT

	// ADD

	// SUBTRACT

	// LOGICAL

	// ROTATE

	// SPECIALS

	// INPUT/OUTPUT

	// CONTROL
	decode[0x00] = (*CPU).NOP
	decode[0x4C] = (*CPU).HLT
}

func (cpu *CPU) NOP() { // 0x00
}

func (cpu *CPU) LDA() { // 0x3A
	cpu.A = cpu.bus.ReadByte(cpu.fetchWord())
}

// func (cpu *CPU) MVI_A() { // 0x3A

// 	// cpu.A = cpu.bus.ReadByte(cpu.fetchWord())

// }

func (cpu *CPU) HLT() { // 0x4C
	fmt.Println("HALT!")
	cpu.halted = true
}

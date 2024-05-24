package cpu

import "log"

// Execute executes the 8-bit instruction passed in via opCode
// Grouped by instruction set group as per "Table 2. Instruction Set Summary",
// in the Intel 8080A 8-BIT N-CHANNEL MICROPROCESSOR datasheet.
func (cpu *CPU) Execute(opCode byte) {

	switch opCode {

	// MOVE, LOAD AND STORE
	case 0x3E: // MVI A
		cpu.A = cpu.fetchByte()
	case 0x06: // MVI B
		cpu.B = cpu.fetchByte()
	case 0x0E: // MVI C
		cpu.C = cpu.fetchByte()
	case 0x10: // MVI D
		cpu.D = cpu.fetchByte()
	case 0x1E: // MVI E
		cpu.E = cpu.fetchByte()
	case 0x1A: // MVI H
		cpu.H = cpu.fetchByte()
	case 0x2E: // MVI L
		cpu.L = cpu.fetchByte()
	case 0x24: // MVI M
		cpu.Bus.WriteByte(joinBytes(cpu.L, cpu.H), cpu.fetchByte())
	case 0x3A: // LDA
		cpu.A = cpu.Bus.ReadByte(cpu.fetchWord())
	// STACK OPERATIONS

	// JUMP

	// RESTART

	// INCREMENT AND DECREMENT
	case 0x3C: // INR A
		cpu.flags.AuxCarry = (cpu.A & 0x0F) == 0x0F
		cpu.A++
		cpu.flags.Sign = cpu.A&(1<<7) > 0
		cpu.flags.Zero = cpu.A == 0
		cpu.flags.Parity = getParity(cpu.A)

	case 0x04: // INR B
	case 0x0C: // INR_C

	// ADD

	// SUBTRACT

	// LOGICAL

	// ROTATE

	// SPECIALS

	// INPUT/OUTPUT

	// CONTROL
	case 0x00: // NOP
	case 0x4C: // HLT
		cpu.halted = true

	default:
		cpu.halted = true
		log.Printf("instruction %02X not found", opCode)
	}
}

package cpu

import "fmt"

const (
	NOCARRY   = false
	WITHCARRY = true
)

// Execute executes the 8-bit instruction passed in via opCode
// Grouped by instruction set group as per "Table 2. Instruction Set Summary",
// in the Intel 8080A 8-BIT N-CHANNEL MICROPROCESSOR datasheet.
func (cpu *CPU) Execute(opCode byte) error {
	// fmt.Printf("Executing instruction: 0x%02X\n", opCode)

	switch opCode {
	// MOVE, LOAD AND STORE
	case 0x40: // MOV B,B - Move register to register
		temp := cpu.B // Redundant, but added for completeness
		cpu.B = temp
	case 0x41: // MOV B,C - Move register to register
		cpu.B = cpu.C
	case 0x42: // MOV B,D - Move register to register
		cpu.B = cpu.D
	case 0x43: // MOV B,E - Move register to register
		cpu.B = cpu.E
	case 0x44: // MOV B,H - Move register to register
		cpu.B = cpu.H
	case 0x45: // MOV B,L - Move register to register
		cpu.B = cpu.L
	case 0x46: // MOV B,M - Move memory to register
		cpu.B = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x47: // MOV B,A - Move register to register
		cpu.B = cpu.A

	case 0x48: // MOV C,B - Move register to register
		cpu.C = cpu.B
	case 0x49: // MOV C,C - Move register to register
		temp := cpu.C
		cpu.C = temp
	case 0x4A: // MOV C,D - Move register to register
		cpu.C = cpu.D
	case 0x4B: // MOV C,E - Move register to register
		cpu.C = cpu.E
	case 0x4C: // MOV C,H - Move register to register
		cpu.C = cpu.H
	case 0x4D: // MOV C,L - Move register to register
		cpu.C = cpu.L
	case 0x4E: // MOV C,M - Move memory to register
		cpu.C = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x4F: // MOV C,A - Move register to register
		cpu.C = cpu.A

	case 0x50: // MOV D,B - Move register to register
		cpu.D = cpu.B
	case 0x51: // MOV D,C - Move register to register
		cpu.D = cpu.C
	case 0x52: // MOV D,D - Move register to register
		temp := cpu.D
		cpu.D = temp
	case 0x53: // MOV D,E - Move register to register
		cpu.D = cpu.E
	case 0x54: // MOV D,H - Move register to register
		cpu.D = cpu.H
	case 0x55: // MOV D,L - Move register to register
		cpu.D = cpu.L
	case 0x56: // MOV D,M - Move memory to register
		cpu.D = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x57: // MOV D,A
		cpu.D = cpu.A

	case 0x58: // MOV E,B - Move register to register
		cpu.E = cpu.B
	case 0x59: // MOV E,C - Move register to register
		cpu.E = cpu.C
	case 0x5A: // MOV E,D - Move register to register
		cpu.E = cpu.D
	case 0x5B: // MOV E,E - Move register to register
		temp := cpu.E
		cpu.E = temp
	case 0x5C: // MOV E,H - Move register to register
		cpu.E = cpu.H
	case 0x5D: // MOV E,L - Move register to register
		cpu.E = cpu.L
	case 0x5E: // MOV E,M - Move memory to register
		cpu.E = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x5F: // MOV E,A - Move register to register
		cpu.E = cpu.A

	case 0x60: // MOV H,B - Move register to register
		cpu.H = cpu.B
	case 0x61: // MOV H,C - Move register to register
		cpu.H = cpu.C
	case 0x62: // MOV H,D - Move register to register
		cpu.H = cpu.D
	case 0x63: // MOV H,E - Move register to register
		cpu.H = cpu.E
	case 0x64: // MOV H,H - Move register to register
		temp := cpu.H
		cpu.H = temp
	case 0x65: // MOV H,L - Move register to register
		cpu.H = cpu.L
	case 0x66: // MOV H,M - Move memory to register
		cpu.H = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x67: // MOV H,A - Move register to register
		cpu.H = cpu.A

	case 0x68: // MOV L,B - Move register to register
		cpu.L = cpu.B
	case 0x69: // MOV L,C - Move register to register
		cpu.L = cpu.C
	case 0x6A: // MOV L,D - Move register to register
		cpu.L = cpu.D
	case 0x6B: // MOV L,E - Move register to register
		cpu.L = cpu.E
	case 0x6C: // MOV L,H - Move register to register
		cpu.L = cpu.H
	case 0x6D: // MOV L,L - Move register to register
		temp := cpu.L
		cpu.L = temp
	case 0x6E: // MOV L,M - Move memory to register
		cpu.L = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x6F: // MOV L,A - Move register to register
		cpu.L = cpu.A

	case 0x70: // MOV M,B - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.B)
	case 0x71: // MOV M,C - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.C)
	case 0x72: // MOV M,D - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.D)
	case 0x73: // MOV M,E - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.E)
	case 0x74: // MOV M,H - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.H)
	case 0x75: // MOV M,L - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.L)
	// There is no MOV M, M instruction on the 8080.  0x76 is used for HLT.
	case 0x77: // MOV M,A - Move register to memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.A)

	case 0x78: // MOV A,B - Move register to register
		cpu.A = cpu.B
	case 0x79: // MOV A,C - Move register to register
		cpu.A = cpu.C
	case 0x7A: // MOV A,D - Move register to register
		cpu.A = cpu.D
	case 0x7B: // MOV A,E - Move register to register
		cpu.A = cpu.E
	case 0x7C: // MOV A,H - Move register to register
		cpu.A = cpu.H
	case 0x7D: // MOV A,L - Move register to register
		cpu.A = cpu.L
	case 0x7E: // MOV A,M - Move memory to register
		cpu.A = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x7F: // MOV A,A - Move register to register
		temp := cpu.A
		cpu.A = temp

	case 0x06: // MVI B - Move immediate register
		cpu.B = cpu.fetchByte()
	case 0x0E: // MVI C - Move immediate register
		cpu.C = cpu.fetchByte()
	case 0x16: // MVI D - Move immediate register
		cpu.D = cpu.fetchByte()
	case 0x1E: // MVI E - Move immediate register
		cpu.E = cpu.fetchByte()
	case 0x26: // MVI H - Move immediate register
		cpu.H = cpu.fetchByte()
	case 0x2E: // MVI L - Move immediate register
		cpu.L = cpu.fetchByte()
	case 0x36: // MVI M - Move immediate memory
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.fetchByte())
	case 0x3E: // MVI A - Move immediate register
		cpu.A = cpu.fetchByte()

	case 0x01: // LXI B - Load immediate register paid B&C
		cpu.B, cpu.C = splitWord(cpu.fetchWord())
	case 0x11: // LXI D - Load immediate register pair D&E
		cpu.D, cpu.E = splitWord(cpu.fetchWord())
	case 0x21: // LXI H - Load immediate register pair H&L
		cpu.H, cpu.L = splitWord(cpu.fetchWord())
	case 0x02: // STAX B - Store A indirect
		cpu.Bus.WriteByte(joinBytes(cpu.B, cpu.C), cpu.A)
	case 0x12: // STAX D - Store A indirect
		cpu.Bus.WriteByte(joinBytes(cpu.D, cpu.E), cpu.A)
	case 0x0A: // LDAX B - Load A indirect
		cpu.A = cpu.Bus.ReadByte(joinBytes(cpu.B, cpu.C))
	case 0x1A: // LDAX D - Load A indirect
		cpu.A = cpu.Bus.ReadByte(joinBytes(cpu.D, cpu.E))
	case 0x32: // STA - Store A direct
		cpu.Bus.WriteByte(cpu.fetchWord(), cpu.A)
	case 0x3A: // LDA - Load A direct
		cpu.A = cpu.Bus.ReadByte(cpu.fetchWord())
	case 0x22: // SHLD - Store H&L direct
		address := cpu.fetchWord()
		cpu.Bus.WriteByte(address, cpu.L)
		cpu.Bus.WriteByte(address+1, cpu.H)
	case 0x2A: // LHLD - Load H&L direct
		address := cpu.fetchWord()
		cpu.L = cpu.Bus.ReadByte(address)
		cpu.H = cpu.Bus.ReadByte(address + 1)
	case 0xEB: // XCHG - Exchange D&E, H&L registers
		cpu.D, cpu.E, cpu.H, cpu.L = cpu.H, cpu.L, cpu.D, cpu.E

	// STACK OPERATIONS
	case 0xC5: // PUSH B
		high, low := splitWord(joinBytes(cpu.B, cpu.C))
		cpu.Bus.WriteByte(cpu.stackPointer-1, high)
		cpu.Bus.WriteByte(cpu.stackPointer-2, low)
		cpu.stackPointer -= 2
	case 0xD5: // PUSH D
		high, low := splitWord(joinBytes(cpu.D, cpu.E))
		cpu.Bus.WriteByte(cpu.stackPointer-1, high)
		cpu.Bus.WriteByte(cpu.stackPointer-2, low)
		cpu.stackPointer -= 2
	case 0xE5: // PUSH H
		high, low := splitWord(joinBytes(cpu.H, cpu.L))
		cpu.Bus.WriteByte(cpu.stackPointer-1, high)
		cpu.Bus.WriteByte(cpu.stackPointer-2, low)
		cpu.stackPointer -= 2
	case 0xF5: // PUSH PSW

		// high, low := splitWord(joinBytes(cpu.H, cpu.L))
		// cpu.Bus.WriteByte(cpu.stackPointer-1, high)
		// cpu.Bus.WriteByte(cpu.stackPointer-2, low)
		// cpu.stackPointer -= 2

	case 0xC1: // POP B
		return ErrNotImplemented(opCode)
	case 0xD1: // POP D
		return ErrNotImplemented(opCode)
	case 0xE1: // POP H
		return ErrNotImplemented(opCode)
	case 0xF1: // POP PSW
		return ErrNotImplemented(opCode)
	case 0xE3: // XTHL
		return ErrNotImplemented(opCode)
	case 0xF9: // SPHL
		return ErrNotImplemented(opCode)
	case 0x31: // LXI SP
		cpu.stackPointer = cpu.fetchWord()
	case 0x33: // INX SP
		return ErrNotImplemented(opCode)
	case 0x3B: // DCX SP
		return ErrNotImplemented(opCode)
	case 0x39: // DAD SP
		return ErrNotImplemented(opCode)

	// JUMP
	case 0xC3: // JMP
		cpu.programCounter = cpu.fetchWord()
	case 0xDA: // JC
		return ErrNotImplemented(opCode)
	case 0xD2: // JNC
		return ErrNotImplemented(opCode)
	case 0xCA: // JZ
		return ErrNotImplemented(opCode)
	case 0xC2: // JNZ
		return ErrNotImplemented(opCode)
	case 0xF2: // JP
		return ErrNotImplemented(opCode)
	case 0xFA: // JM
		return ErrNotImplemented(opCode)
	case 0xEA: // JPE
		return ErrNotImplemented(opCode)
	case 0xE2: // JPO
		return ErrNotImplemented(opCode)
	case 0xE9: // PCHL
		return ErrNotImplemented(opCode)

	// CALL
	case 0xCD: // CALL
		return ErrNotImplemented(opCode)
	case 0xDC: // CC
		return ErrNotImplemented(opCode)
	case 0xD4: // CNC
		return ErrNotImplemented(opCode)
	case 0xCC: // CZ
		return ErrNotImplemented(opCode)
	case 0xC4: // CNZ
		return ErrNotImplemented(opCode)
	case 0xF4: // CP
		return ErrNotImplemented(opCode)
	case 0xFC: // CM
		return ErrNotImplemented(opCode)
	case 0xEC: // CPE
		return ErrNotImplemented(opCode)
	case 0xE4: // CPO
		return ErrNotImplemented(opCode)

	// RETURN
	case 0xC9: // RET
		return ErrNotImplemented(opCode)
	case 0xD8: // RC
		return ErrNotImplemented(opCode)
	case 0xD0: // RNC
		return ErrNotImplemented(opCode)
	case 0xC8: // RZ
		return ErrNotImplemented(opCode)
	case 0xC0: // RNZ
		return ErrNotImplemented(opCode)
	case 0xF0: // RP
		return ErrNotImplemented(opCode)
	case 0xF8: // RM
		return ErrNotImplemented(opCode)
	case 0xE8: // RPE
		return ErrNotImplemented(opCode)
	case 0xE0: // RPO
		return ErrNotImplemented(opCode)

	// RESTART
	case 0xC7: // RST 0
		return ErrNotImplemented(opCode)
	case 0xCF: // RST 1
		return ErrNotImplemented(opCode)
	case 0xD7: // RST 2
		return ErrNotImplemented(opCode)
	case 0xDF: // RST 3
		return ErrNotImplemented(opCode)
	case 0xE7: // RST 4
		return ErrNotImplemented(opCode)
	case 0xEF: // RST 5
		return ErrNotImplemented(opCode)
	case 0xF7: // RST 6
		return ErrNotImplemented(opCode)
	case 0xFF: // RST 7
		return ErrNotImplemented(opCode)

	// INCREMENT AND DECREMENT
	case 0x04: // INR B
		cpu.inr(&cpu.B)
	case 0x0C: // INR C
		cpu.inr(&cpu.C)
	case 0x14: // INR D
		cpu.inr(&cpu.D)
	case 0x1C: // INR E
		cpu.inr(&cpu.E)
	case 0x24: // INR H
		cpu.inr(&cpu.H)
	case 0x2C: // INR L
		cpu.inr(&cpu.L)
	case 0x34: // INR M
		tempM := cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
		cpu.inr(&tempM)
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), tempM)
	case 0x3C: // INR A
		cpu.inr(&cpu.A)

	case 0x05: // DCR B
		cpu.dcr(&cpu.B)
	case 0x0D: // DCR C
		cpu.dcr(&cpu.C)
	case 0x15: // DCR D
		cpu.dcr(&cpu.D)
	case 0x1D: // DCR E
		cpu.dcr(&cpu.E)
	case 0x25: // DCR H
		cpu.dcr(&cpu.H)
	case 0x2D: // DCR L
		cpu.dcr(&cpu.L)
	case 0x35: // DCR M
		tempM := cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
		cpu.dcr(&tempM)
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), tempM)
	case 0x3D: // DCR A
		cpu.dcr(&cpu.A)

	case 0x03: // INX B
		cpu.B, cpu.C = splitWord(joinBytes(cpu.B, cpu.C) + 1)
	case 0x13: // INX D
		cpu.D, cpu.E = splitWord(joinBytes(cpu.D, cpu.E) + 1)
	case 0x23: // INX H
		cpu.H, cpu.L = splitWord(joinBytes(cpu.H, cpu.L) + 1)
	case 0x0B: // DCX B
		cpu.B, cpu.C = splitWord(joinBytes(cpu.B, cpu.C) - 1)
	case 0x1B: // DCX D
		cpu.D, cpu.E = splitWord(joinBytes(cpu.D, cpu.E) - 1)
	case 0x2B: // DCX H
		cpu.H, cpu.L = splitWord(joinBytes(cpu.H, cpu.L) - 1)

	// ADD
	case 0x80: // ADD B
		cpu.add(cpu.B, NOCARRY)
	case 0x81: // ADD C
		cpu.add(cpu.C, NOCARRY)
	case 0x82: // ADD D
		cpu.add(cpu.D, NOCARRY)
	case 0x83: // ADD E
		cpu.add(cpu.E, NOCARRY)
	case 0x84: // ADD H
		cpu.add(cpu.H, NOCARRY)
	case 0x85: // ADD L
		cpu.add(cpu.L, NOCARRY)
	case 0x86: // ADD M
		cpu.add(cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L)), NOCARRY)
	case 0x87: // ADD A
		cpu.add(cpu.A, NOCARRY)

	case 0x88: // ADC B
		cpu.add(cpu.B, cpu.flags.Carry)
	case 0x89: // ADC C
		cpu.add(cpu.C, cpu.flags.Carry)
	case 0x8A: // ADC D
		cpu.add(cpu.D, cpu.flags.Carry)
	case 0x8B: // ADC E
		cpu.add(cpu.E, cpu.flags.Carry)
	case 0x8C: // ADC H
		cpu.add(cpu.H, cpu.flags.Carry)
	case 0x8D: // ADC L
		cpu.add(cpu.L, cpu.flags.Carry)
	case 0x8E: // ADC M
		cpu.add(cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L)), cpu.flags.Carry)
	case 0x8F: // ADC A
		cpu.add(cpu.A, cpu.flags.Carry)

	case 0xC6: // ADI
		return ErrNotImplemented(opCode)
	case 0xCE: // ACI
		return ErrNotImplemented(opCode)

	case 0x09: // DAD B
		return ErrNotImplemented(opCode)
	case 0x19: // DAD D
		return ErrNotImplemented(opCode)
	case 0x29: // DAD H
		return ErrNotImplemented(opCode)

	// SUBTRACT
	case 0x90: // SUB B
		return ErrNotImplemented(opCode)
	case 0x91: // SUB C
		return ErrNotImplemented(opCode)
	case 0x92: // SUB D
		return ErrNotImplemented(opCode)
	case 0x93: // SUB E
		return ErrNotImplemented(opCode)
	case 0x94: // SUB H
		return ErrNotImplemented(opCode)
	case 0x95: // SUB L
		return ErrNotImplemented(opCode)
	case 0x96: // SUB M
		return ErrNotImplemented(opCode)
	case 0x97: // SUB A
		return ErrNotImplemented(opCode)
	case 0x98: // SBB B
		return ErrNotImplemented(opCode)
	case 0x99: // SBB C
		return ErrNotImplemented(opCode)
	case 0x9A: // SBB D
		return ErrNotImplemented(opCode)
	case 0x9B: // SBB E
		return ErrNotImplemented(opCode)
	case 0x9C: // SBB H
		return ErrNotImplemented(opCode)
	case 0x9D: // SBB L
		return ErrNotImplemented(opCode)
	case 0x9E: // SBB M
		return ErrNotImplemented(opCode)
	case 0x9F: // SBB A
		return ErrNotImplemented(opCode)

	case 0xD6: // SUI
		return ErrNotImplemented(opCode)
	case 0xDE: // SBI
		return ErrNotImplemented(opCode)

	// LOGICAL
	case 0xA0: // ANA B
		return ErrNotImplemented(opCode)
	case 0xA1: // ANA C
		return ErrNotImplemented(opCode)
	case 0xA2: // ANA D
		return ErrNotImplemented(opCode)
	case 0xA3: // ANA E
		return ErrNotImplemented(opCode)
	case 0xA4: // ANA H
		return ErrNotImplemented(opCode)
	case 0xA5: // ANA L
		return ErrNotImplemented(opCode)
	case 0xA6: // ANA M
		return ErrNotImplemented(opCode)
	case 0xA7: // ANA A
		return ErrNotImplemented(opCode)

	case 0xA8: // XRA B
		return ErrNotImplemented(opCode)
	case 0xA9: // XRA C
		return ErrNotImplemented(opCode)
	case 0xAA: // XRA D
		return ErrNotImplemented(opCode)
	case 0xAB: // XRA E
		return ErrNotImplemented(opCode)
	case 0xAC: // XRA H
		return ErrNotImplemented(opCode)
	case 0xAD: // XRA L
		return ErrNotImplemented(opCode)
	case 0xAE: // XRA M
		return ErrNotImplemented(opCode)
	case 0xAF: // XRA A
		return ErrNotImplemented(opCode)

	case 0xB0: // ORA B
		return ErrNotImplemented(opCode)
	case 0xB1: // ORA C
		return ErrNotImplemented(opCode)
	case 0xB2: // ORA D
		return ErrNotImplemented(opCode)
	case 0xB3: // ORA E
		return ErrNotImplemented(opCode)
	case 0xB4: // ORA H
		return ErrNotImplemented(opCode)
	case 0xB5: // ORA L
		return ErrNotImplemented(opCode)
	case 0xB6: // ORA M
		return ErrNotImplemented(opCode)
	case 0xB7: // ORA A
		return ErrNotImplemented(opCode)

	case 0xB8: // CMP B
		return ErrNotImplemented(opCode)
	case 0xB9: // CMP C
		return ErrNotImplemented(opCode)
	case 0xBA: // CMP D
		return ErrNotImplemented(opCode)
	case 0xBB: // CMP E
		return ErrNotImplemented(opCode)
	case 0xBC: // CMP H
		return ErrNotImplemented(opCode)
	case 0xBD: // CMP L
		return ErrNotImplemented(opCode)
	case 0xBE: // CMP M
		return ErrNotImplemented(opCode)
	case 0xBF: // CMP A
		return ErrNotImplemented(opCode)

	case 0xE6: // ANI
		return ErrNotImplemented(opCode)
	case 0xEE: // XRI
		return ErrNotImplemented(opCode)
	case 0xF6: // ORI
		return ErrNotImplemented(opCode)
	case 0xFE: // CPI
		return ErrNotImplemented(opCode)

	// ROTATE
	case 0x07: // RLC
		return ErrNotImplemented(opCode)
	case 0x0F: // RRC
		return ErrNotImplemented(opCode)
	case 0x17: // RAL
		return ErrNotImplemented(opCode)
	case 0x1F: // RAR
		return ErrNotImplemented(opCode)

	// SPECIALS
	case 0x2F: // CMA
		return ErrNotImplemented(opCode)
	case 0x37: // STC
		return ErrNotImplemented(opCode)
	case 0x3F: // CMC
		return ErrNotImplemented(opCode)
	case 0x27: // DAA
		return ErrNotImplemented(opCode)

	// INPUT/OUTPUT
	case 0xDB: // IN
		return ErrNotImplemented(opCode)
	case 0xD3: // OUT
		return ErrNotImplemented(opCode)

	// CONTROL
	case 0xFB: // EI
		return ErrNotImplemented(opCode)
	case 0xF3: // DI
		return ErrNotImplemented(opCode)
	case 0x00: // NOP
		// Do nothing
	case 0x76: // HLT
		cpu.halted = true

	default:
		cpu.halted = true
		return fmt.Errorf("instruction %02X not found", opCode)
	}

	return nil
}

// inr increments the value of a given register by 1, updating the CPU flags accordingly.
//
// This method performs the following steps:
// 1. Checks and sets the auxiliary carry flag based on the lower nibble of the register.
// 2. Increments the value of the register by 1.
// 3. Sets the Sign, Zero, and Parity flags based on the new value of the register.
//
// Parameters:
// - register (*byte): A pointer to the byte register to be incremented.
//
// Example:
//
//	register := byte(0x0F)
//	cpu := &CPU{}
//	cpu.inr(&register)
//	// register is 0x10
//	// cpu.flags are updated based on the result
func (cpu *CPU) inr(register *byte) {
	cpu.flags.AuxCarry = (*register & 0x0F) == 0x0F
	*register++
	cpu.setSignZeroParityFlags(*register)
}

// dcr decrements the value of a given register by 1, updating the CPU flags accordingly.
//
// This method performs the following steps:
// 1. Checks and sets the auxiliary carry flag based on the lower nibble of the register.
// 2. Decrements the value of the register by 1.
// 3. Sets the Sign, Zero, and Parity flags based on the new value of the register.
//
// Parameters:
// - register (*byte): A pointer to the byte register to be decremented.
//
// Example:
//
//	register := byte(0x10)
//	cpu := &CPU{}
//	cpu.dcr(&register)
//	// register is 0x0F
//	// cpu.flags are updated based on the result
func (cpu *CPU) dcr(register *byte) {
	cpu.flags.AuxCarry = (*register & 0x0F) == 0x0F
	*register--
	cpu.setSignZeroParityFlags(*register)
}

// add adds the value of a register and an optional carry-in to the accumulator,
// updating the accumulator and the CPU flags accordingly.
//
// This method performs the following steps:
// 1. Converts the carry-in boolean to a byte value (1 if true, 0 if false).
// 2. Adds the accumulator, the register value, and the carry-in value.
// 3. Sets the Sign, Zero, and Parity flags based on the result.
// 4. Checks for carry-out and auxiliary carry-out, updating the corresponding flags.
// 5. Updates the accumulator with the result.
//
// Parameters:
// - register (byte): The value of the register to be added to the accumulator.
// - carry (bool): A boolean indicating if there is an initial carry-in.
//
// Example:
//
//	cpu := &CPU{A: 0x10, flags: Flags{}}
//	cpu.add(0x20, true)
//	// cpu.A is 0x31
//	// cpu.flags are updated based on the result
func (cpu *CPU) add(register byte, carry bool) {
	var carryValue byte // Cast the bool to a byte for our carry calculation
	if carry {
		carryValue = 1
	}

	result := cpu.A + register + carryValue
	cpu.setSignZeroParityFlags(result)
	cpu.flags.Carry, cpu.flags.AuxCarry = checkCarryOut(cpu.A, register, carry)
	cpu.A = result
}

// joinBytes combines two bytes into a 16-bit word.
//
// This function takes a high byte and a low byte and joins them to form
// a 16-bit word, with the high byte occupying the most significant 8 bits
// and the low byte occupying the least significant 8 bits.
//
// Parameters:
// - high (byte): The high byte (most significant 8 bits).
// - low (byte): The low byte (least significant 8 bits).
//
// Returns:
// - word: The 16-bit word resulting from combining the high and low bytes.
//
// Example:
//
//	word := joinBytes(0x12, 0x34)
//	// word is 0x1234
func joinBytes(high, low byte) word {
	return word(high)<<8 | word(low)
}

// splitWord splits a 16-bit word into its high and low bytes.
//
// This function takes a 16-bit word and extracts the high (most significant)
// byte and the low (least significant) byte.
//
// Parameters:
// - address (word): The 16-bit word to be split.
//
// Returns:
// - byte: The high byte (most significant 8 bits) of the input word.
// - byte: The low byte (least significant 8 bits) of the input word.
//
// Example:
//
//	high, low := splitWord(0x1234)
//	// high is 0x12
//	// low is 0x34
func splitWord(address word) (high, low byte) {
	return byte(address >> 8), byte(address)
}

// checkCarryOut calculates and returns the carry-out and auxiliary carry-out
// from the addition of two bytes and an optional carry-in.
//
// The function performs the following steps:
// 1. Converts the carry-in boolean to a byte value (1 if true, 0 if false).
// 2. Adds the two input bytes and the carry-in value, storing the result in a 16-bit word.
// 3. Determines if there is a carry-out from the 8th bit (outCarry).
// 4. Determines if there is an auxiliary carry-out from the 4th bit (auxCarry).
//
// Parameters:
// - a (byte): The first byte to be added.
// - b (byte): The second byte to be added.
// - inCarry (bool): A boolean indicating if there is an initial carry-in.
//
// Returns:
// - bool: True if there is a carry-out from the 8th bit, false otherwise.
// - bool: True if there is an auxiliary carry-out from the 4th bit, false otherwise.
//
// Example:
//
//	outCarry, auxCarry := checkCarryOut(0x1F, 0xE1, true)
//	// outCarry is false
//	// auxCarry is true
func checkCarryOut(a, b byte, inCarry bool) (bool, bool) {
	var carryValue byte // Cast the bool to a byte for our carry calculation
	if inCarry {
		carryValue = 1
	}
	sum := word(a) + word(b) + word(carryValue)
	outCarry := sum > 0xFF                             // Carry-out on bit 8?
	auxCarry := (a&0xF + b&0xF + carryValue&0xF) > 0xF // Carry-out on bit 4?

	return outCarry, auxCarry
}

// setSignZeroParityFlags sets the Sign, Zero, and Parity flags in the CPU's flag register
// based on the provided input byte.
//
// This function updates the following flags:
// - Sign flag: Set if the most significant bit (bit 7) of the input byte is 1.
// - Zero flag: Set if the input byte is 0.
// - Parity flag: Set if the input byte has an even number of 1 bits.
//
// Parameters:
// - input (byte): The byte value used to determine the flag states.
//
// Example:
//
//	cpu := &CPU{}
//	cpu.setSignZeroParityFlags(0x80)
//	// cpu.flags.Sign is true
//	// cpu.flags.Zero is false
//	// cpu.flags.Parity is true (0x80 has one 1 bit, which is odd, so Parity is false)
func (cpu *CPU) setSignZeroParityFlags(input byte) {
	cpu.flags.Sign = input&(1<<7) > 0
	cpu.flags.Zero = input == 0
	cpu.flags.Parity = getParity(input)
}

// temporary function to be removed when all instructions are implemented
func ErrNotImplemented(opCode byte) error {
	return fmt.Errorf("instruction 0x%02X not implemented", opCode)
}

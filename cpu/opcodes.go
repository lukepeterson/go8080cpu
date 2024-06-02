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
	case 0x78: // MOV A,B
		return ErrNotImplemented(opCode)
	case 0x79: // MOV A,C
		return ErrNotImplemented(opCode)
	case 0x7A: // MOV A,D
		return ErrNotImplemented(opCode)
	case 0x7B: // MOV A,E
		return ErrNotImplemented(opCode)
	case 0x7C: // MOV A,H
		return ErrNotImplemented(opCode)
	case 0x7D: // MOV A,L
		return ErrNotImplemented(opCode)
	case 0x7E: // MOV A,M
		cpu.A = cpu.Bus.ReadByte(joinBytes(cpu.H, cpu.L))
	case 0x7F: // MOV A,A
		return ErrNotImplemented(opCode)

	case 0x40: // MOV B,B
		return ErrNotImplemented(opCode)
	case 0x41: // MOV B,C
		return ErrNotImplemented(opCode)
	case 0x42: // MOV B,D
		return ErrNotImplemented(opCode)
	case 0x43: // MOV B,E
		return ErrNotImplemented(opCode)
	case 0x44: // MOV B,H
		return ErrNotImplemented(opCode)
	case 0x45: // MOV B,L
		return ErrNotImplemented(opCode)
	case 0x46: // MOV B,M
		return ErrNotImplemented(opCode)
	case 0x47: // MOV B,A
		return ErrNotImplemented(opCode)

	case 0x48: // MOV C,B
		return ErrNotImplemented(opCode)
	case 0x49: // MOV C,C
		return ErrNotImplemented(opCode)
	case 0x4A: // MOV C,D
		return ErrNotImplemented(opCode)
	case 0x4B: // MOV C,E
		return ErrNotImplemented(opCode)
	case 0x4C: // MOV C,H
		return ErrNotImplemented(opCode)
	case 0x4D: // MOV C,L
		return ErrNotImplemented(opCode)
	case 0x4E: // MOV C,M
		return ErrNotImplemented(opCode)
	case 0x4F: // MOV C,A
		return ErrNotImplemented(opCode)

	case 0x50: // MOV D,B
		return ErrNotImplemented(opCode)
	case 0x51: // MOV D,C
		return ErrNotImplemented(opCode)
	case 0x52: // MOV D,D
		return ErrNotImplemented(opCode)
	case 0x53: // MOV D,E
		return ErrNotImplemented(opCode)
	case 0x54: // MOV D,H
		return ErrNotImplemented(opCode)
	case 0x55: // MOV D,L
		return ErrNotImplemented(opCode)
	case 0x56: // MOV D,M
		return ErrNotImplemented(opCode)
	case 0x57: // MOV D,A
		return ErrNotImplemented(opCode)

	case 0x58: // MOV E,B
		return ErrNotImplemented(opCode)
	case 0x59: // MOV E,C
		return ErrNotImplemented(opCode)
	case 0x5A: // MOV E,D
		return ErrNotImplemented(opCode)
	case 0x5B: // MOV E,E
		return ErrNotImplemented(opCode)
	case 0x5C: // MOV E,H
		return ErrNotImplemented(opCode)
	case 0x5D: // MOV E,L
		return ErrNotImplemented(opCode)
	case 0x5E: // MOV E,M
		return ErrNotImplemented(opCode)
	case 0x5F: // MOV E,A
		return ErrNotImplemented(opCode)

	case 0x60: // MOV H,B
		return ErrNotImplemented(opCode)
	case 0x61: // MOV H,C
		return ErrNotImplemented(opCode)
	case 0x62: // MOV H,D
		return ErrNotImplemented(opCode)
	case 0x63: // MOV H,E
		return ErrNotImplemented(opCode)
	case 0x64: // MOV H,H
		return ErrNotImplemented(opCode)
	case 0x65: // MOV H,L
		return ErrNotImplemented(opCode)
	case 0x66: // MOV H,M
		return ErrNotImplemented(opCode)
	case 0x67: // MOV H,A
		return ErrNotImplemented(opCode)

	case 0x68: // MOV L,B
		return ErrNotImplemented(opCode)
	case 0x69: // MOV L,C
		return ErrNotImplemented(opCode)
	case 0x6A: // MOV L,D
		return ErrNotImplemented(opCode)
	case 0x6B: // MOV L,E
		return ErrNotImplemented(opCode)
	case 0x6C: // MOV L,H
		return ErrNotImplemented(opCode)
	case 0x6D: // MOV L,L
		return ErrNotImplemented(opCode)
	case 0x6E: // MOV L,M
		return ErrNotImplemented(opCode)
	case 0x6F: // MOV L,A
		return ErrNotImplemented(opCode)

	case 0x70: // MOV M,B
		return ErrNotImplemented(opCode)
	case 0x71: // MOV M,C
		return ErrNotImplemented(opCode)
	case 0x72: // MOV M,D
		return ErrNotImplemented(opCode)
	case 0x73: // MOV M,E
		return ErrNotImplemented(opCode)
	case 0x74: // MOV M,H
		return ErrNotImplemented(opCode)
	case 0x75: // MOV M,L
		return ErrNotImplemented(opCode)
	case 0x77: // MOV M,A
		return ErrNotImplemented(opCode)

	case 0x06: // MVI B
		cpu.B = cpu.fetchByte()
	case 0x0E: // MVI C
		cpu.C = cpu.fetchByte()
	case 0x16: // MVI D
		cpu.D = cpu.fetchByte()
	case 0x1E: // MVI E
		cpu.E = cpu.fetchByte()
	case 0x26: // MVI H
		cpu.H = cpu.fetchByte()
	case 0x2E: // MVI L
		cpu.L = cpu.fetchByte()
	case 0x36: // MVI M
		cpu.Bus.WriteByte(joinBytes(cpu.H, cpu.L), cpu.fetchByte())
	case 0x3E: // MVI A
		cpu.A = cpu.fetchByte()

	case 0x01: // LXI B
		return ErrNotImplemented(opCode)
	case 0x11: // LXI D
		return ErrNotImplemented(opCode)
	case 0x21: // LXI H
		cpu.L, cpu.H = splitWord(cpu.fetchWord())
	case 0x02: // STAX B
		return ErrNotImplemented(opCode)
	case 0x12: // STAX D
		return ErrNotImplemented(opCode)
	case 0x0A: // LDAX B
		return ErrNotImplemented(opCode)
	case 0x1A: // LDAX D
		return ErrNotImplemented(opCode)
	case 0x32: // STA
		return ErrNotImplemented(opCode)
	case 0x3A: // LDA
		cpu.A = cpu.Bus.ReadByte(cpu.fetchWord())
	case 0x22: // SHLD
		return ErrNotImplemented(opCode)
	case 0x2A: // LHLD
		return ErrNotImplemented(opCode)
	case 0xEB: // XCH
		return ErrNotImplemented(opCode)

	// STACK OPERATIONS
	case 0xC5: // PUSH B
		return ErrNotImplemented(opCode)
	case 0xD5: // PUSH D
		return ErrNotImplemented(opCode)
	case 0xE5: // PUSH H
		return ErrNotImplemented(opCode)
	case 0xF5: // PUSH PSW
		return ErrNotImplemented(opCode)
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
		return ErrNotImplemented(opCode)
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

func (cpu *CPU) inr(register *byte) {
	cpu.flags.AuxCarry = (*register & 0x0F) == 0x0F
	*register++
	cpu.setSignZeroParityFlags(*register)
}

func (cpu *CPU) dcr(register *byte) {
	cpu.flags.AuxCarry = (*register & 0x0F) == 0x0F
	*register--
	cpu.setSignZeroParityFlags(*register)
}

// add adds the register specified in register to the accumulator (R), as well as calculation and setting the sign, zero, parity, carry and aux carry flags where appropriate.
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

func ErrNotImplemented(opCode byte) error {
	return fmt.Errorf("instruction 0x%02X not implemented", opCode)
}

func joinBytes(high, low byte) word {
	return word(high)<<8 | word(low)
}

func splitWord(address word) (high, low byte) {
	return byte(address >> 8), byte(address)
}

// carry checks if there is a carry-out from the addition of three bytes (one of which is a carry-in), in addition to checking if there is a carry from the lower 4 bits (an aux carry).
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

func (cpu *CPU) setSignZeroParityFlags(input byte) {
	cpu.flags.Sign = input&(1<<7) > 0
	cpu.flags.Zero = input == 0
	cpu.flags.Parity = getParity(input)
}

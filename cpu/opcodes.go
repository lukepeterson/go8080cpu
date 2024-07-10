package cpu

import (
	"fmt"
)

const (
	NOCARRY   = 0
	WITHCARRY = 1
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.B = readByte
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.C = readByte
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.D = readByte
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.E = readByte
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.H = readByte
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.L = readByte
	case 0x6F: // MOV L,A - Move register to register
		cpu.L = cpu.A
	case 0x70: // MOV M,B - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.B)
	case 0x71: // MOV M,C - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.C)
	case 0x72: // MOV M,D - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.D)
	case 0x73: // MOV M,E - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.E)
	case 0x74: // MOV M,H - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.H)
	case 0x75: // MOV M,L - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.L)
	case 0x77: // MOV M,A - Move register to memory
		cpu.Bus.WriteByteAt(cpu.getHL(), cpu.A)
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
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.A = readByte
	case 0x7F: // MOV A,A - Move register to register
		temp := cpu.A
		cpu.A = temp
	case 0x06: // MVI B - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.B = fetchedByte
	case 0x0E: // MVI C - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.C = fetchedByte
	case 0x16: // MVI D - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.D = fetchedByte
	case 0x1E: // MVI E - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.E = fetchedByte
	case 0x26: // MVI H - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.H = fetchedByte
	case 0x2E: // MVI L - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.L = fetchedByte
	case 0x36: // MVI M - Move immediate memory
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.Bus.WriteByteAt(cpu.getHL(), fetchedByte)
	case 0x3E: // MVI A - Move immediate register
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.A = fetchedByte
	case 0x01: // LXI B - Load immediate register paid B&C
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.B, cpu.C = splitWord(fetchedWord)
	case 0x11: // LXI D - Load immediate register pair D&E
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.D, cpu.E = splitWord(fetchedWord)
	case 0x21: // LXI H - Load immediate register pair H&L
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.H, cpu.L = splitWord(fetchedWord)
	case 0x02: // STAX B - Store A indirect
		cpu.Bus.WriteByteAt(cpu.getBC(), cpu.A)
	case 0x12: // STAX D - Store A indirect
		cpu.Bus.WriteByteAt(cpu.getDE(), cpu.A)
	case 0x0A: // LDAX B - Load A indirect
		readByte, err := cpu.Bus.ReadByteAt(cpu.getBC())
		if err != nil {
			return err
		}
		cpu.A = readByte
	case 0x1A: // LDAX D - Load A indirect
		readByte, err := cpu.Bus.ReadByteAt(cpu.getDE())
		if err != nil {
			return err
		}
		cpu.A = readByte
	case 0x32: // STA - Store A direct
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.Bus.WriteByteAt(fetchedWord, cpu.A)
	case 0x3A: // LDA - Load A direct
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		readByte, err := cpu.Bus.ReadByteAt(address)
		if err != nil {
			return err
		}
		cpu.A = readByte
	case 0x22: // SHLD - Store H&L direct
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.Bus.WriteByteAt(fetchedWord, cpu.L)
		cpu.Bus.WriteByteAt(fetchedWord+1, cpu.H)
	case 0x2A: // LHLD - Load H&L direct
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		readByte, err := cpu.Bus.ReadByteAt(address)
		if err != nil {
			return err
		}
		cpu.L = readByte
		readByte, err = cpu.Bus.ReadByteAt(address + 1)
		if err != nil {
			return err
		}
		cpu.H = readByte
	case 0xEB: // XCHG - Exchange D&E, H&L registers
		cpu.D, cpu.E, cpu.H, cpu.L = cpu.H, cpu.L, cpu.D, cpu.E

	// STACK OPERATIONS
	case 0xC5: // PUSH B - Push register pair B&C to stack
		err := cpu.pushStack(cpu.getBC())
		if err != nil {
			return err
		}
	case 0xD5: // PUSH D - Push register pair D&E to stack
		err := cpu.pushStack(cpu.getDE())
		if err != nil {
			return err
		}
	case 0xE5: // PUSH H - Push register pair H&L to stack
		err := cpu.pushStack(cpu.getHL())
		if err != nil {
			return err
		}
	case 0xF5: // PUSH PSW - Push register A and flags to stack
		err := cpu.pushStack(cpu.getAPSW())
		if err != nil {
			return err
		}
	case 0xC1: // POP B - Pop register pair B&C off stack
		readWord, err := cpu.popStack()
		if err != nil {
			return err
		}
		cpu.B, cpu.C = splitWord(readWord)
	case 0xD1: // POP D - Pop register pair D&E off stack
		readWord, err := cpu.popStack()
		if err != nil {
			return err
		}
		cpu.D, cpu.E = splitWord(readWord)
	case 0xE1: // POP H - Pop register pair H&L off stack
		readWord, err := cpu.popStack()
		if err != nil {
			return err
		}
		cpu.H, cpu.L = splitWord(readWord)
	case 0xF1: // POP PSW - Pop register A and flags off stack
		readWord, err := cpu.popStack()
		if err != nil {
			return err
		}
		var flags byte
		cpu.A, flags = splitWord(readWord)
		cpu.setFlags(flags)
	case 0xE3: // XTHL - Exchange top of stack with H&L
		readByte, err := cpu.Bus.ReadByteAt(cpu.stackPointer + 1)
		if err != nil {
			return err
		}
		cpu.H = readByte
		readByte, err = cpu.Bus.ReadByteAt(cpu.stackPointer)
		if err != nil {
			return err
		}
		cpu.L = readByte
		cpu.Bus.WriteByteAt(cpu.stackPointer+1, cpu.H)
		cpu.Bus.WriteByteAt(cpu.stackPointer, cpu.L)
	case 0xF9: // SPHL - Load stack pointer from H&L
		cpu.stackPointer = cpu.getHL()
	case 0x31: // LXI SP - Load immediate stack pointer
		fetchedWord, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.stackPointer = fetchedWord
	case 0x33: // INX SP - Increment stack pointer
		cpu.stackPointer++
	case 0x3B: // DCX SP - Decrement stack pointer
		cpu.stackPointer--

	// JUMP
	case 0xC3: // JMP - Jump unconditional
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.programCounter = address
	case 0xDA: // JC - Jump on carry
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if cpu.flags.Carry {
			cpu.programCounter = address
		}
	case 0xD2: // JNC - Jump on no carry
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if !cpu.flags.Carry {
			cpu.programCounter = address
		}
	case 0xCA: // JZ - Jump on zero
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if cpu.flags.Zero {
			cpu.programCounter = address
		}
	case 0xC2: // JNZ - Jump on no zero
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if !cpu.flags.Zero {
			cpu.programCounter = address
		}
	case 0xF2: // JP - Jump on positive
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if !cpu.flags.Sign {
			cpu.programCounter = address
		}
	case 0xFA: // JM - Jump on minus
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if cpu.flags.Sign {
			cpu.programCounter = address
		}
	case 0xEA: // JPE - Jump on parity even
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if cpu.flags.Parity {
			cpu.programCounter = address
		}
	case 0xE2: // JPO - Jump on parity odd
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		if !cpu.flags.Parity {
			cpu.programCounter = address
		}
	case 0xE9: // PCHL - H&L to program counter
		cpu.programCounter = joinBytes(cpu.H, cpu.L)

	// CALL
	case 0xCD: // CALL
		address, err := cpu.fetchWord()
		if err != nil {
			return err
		}
		cpu.pushStack(cpu.programCounter)
		cpu.programCounter = address
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
		address, err := cpu.popStack()
		if err != nil {
			return err
		}
		cpu.programCounter = address
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
	case 0x04: // INR B - Increment register
		cpu.inr(&cpu.B)
	case 0x0C: // INR C - Increment register
		cpu.inr(&cpu.C)
	case 0x14: // INR D - Increment register
		cpu.inr(&cpu.D)
	case 0x1C: // INR E - Increment register
		cpu.inr(&cpu.E)
	case 0x24: // INR H - Increment register
		cpu.inr(&cpu.H)
	case 0x2C: // INR L - Increment register
		cpu.inr(&cpu.L)
	case 0x34: // INR M - Increment memory
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.inr(&readByte)
		cpu.Bus.WriteByteAt(cpu.getHL(), readByte)
	case 0x3C: // INR A - Increment register
		cpu.inr(&cpu.A)
	case 0x05: // DCR B - Decrement register
		cpu.dcr(&cpu.B)
	case 0x0D: // DCR C - Decrement register
		cpu.dcr(&cpu.C)
	case 0x15: // DCR D - Decrement register
		cpu.dcr(&cpu.D)
	case 0x1D: // DCR E - Decrement register
		cpu.dcr(&cpu.E)
	case 0x25: // DCR H - Decrement register
		cpu.dcr(&cpu.H)
	case 0x2D: // DCR L - Decrement register
		cpu.dcr(&cpu.L)
	case 0x35: // DCR M - Decrement memory
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.dcr(&readByte)
		cpu.Bus.WriteByteAt(cpu.getHL(), readByte)
	case 0x3D: // DCR A - Decrement register
		cpu.dcr(&cpu.A)
	case 0x03: // INX B - Increment B&C registers
		cpu.B, cpu.C = splitWord(cpu.getBC() + 1)
	case 0x13: // INX D - Increment D&E registers
		cpu.D, cpu.E = splitWord(cpu.getDE() + 1)
	case 0x23: // INX H - Increment H&L registers
		cpu.H, cpu.L = splitWord(cpu.getHL() + 1)
	case 0x0B: // DCX B - Decrement B&C registers
		cpu.B, cpu.C = splitWord(cpu.getBC() - 1)
	case 0x1B: // DCX D - Decrement D&E registers
		cpu.D, cpu.E = splitWord(cpu.getDE() - 1)
	case 0x2B: // DCX H - Decrement H&L registers
		cpu.H, cpu.L = splitWord(cpu.getHL() - 1)

	// ADD
	case 0x80: // ADD B - Add register to A
		cpu.add(cpu.B, NOCARRY)
	case 0x81: // ADD C - Add register to A
		cpu.add(cpu.C, NOCARRY)
	case 0x82: // ADD D - Add register to A
		cpu.add(cpu.D, NOCARRY)
	case 0x83: // ADD E - Add register to A
		cpu.add(cpu.E, NOCARRY)
	case 0x84: // ADD H - Add register to A
		cpu.add(cpu.H, NOCARRY)
	case 0x85: // ADD L - Add register to A
		cpu.add(cpu.L, NOCARRY)
	case 0x86: // ADD M - Add memory to A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.add(readByte, NOCARRY)
	case 0x87: // ADD A - Add register to A
		cpu.add(cpu.A, NOCARRY)
	case 0x88: // ADC B - Add register to A with carry
		cpu.add(cpu.B, WITHCARRY)
	case 0x89: // ADC C - Add register to A with carry
		cpu.add(cpu.C, WITHCARRY)
	case 0x8A: // ADC D - Add register to A with carry
		cpu.add(cpu.D, WITHCARRY)
	case 0x8B: // ADC E - Add register to A with carry
		cpu.add(cpu.E, WITHCARRY)
	case 0x8C: // ADC H - Add register to A with carry
		cpu.add(cpu.H, WITHCARRY)
	case 0x8D: // ADC L - Add register to A with carry
		cpu.add(cpu.L, WITHCARRY)
	case 0x8E: // ADC M - Add memory to A with carry
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.add(readByte, WITHCARRY)
	case 0x8F: // ADC A - Add register to A with carry
		cpu.add(cpu.A, WITHCARRY)
	case 0xC6: // ADI
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.add(fetchedByte, NOCARRY)
	case 0xCE: // ACI
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.add(fetchedByte, WITHCARRY)
	case 0x09: // DAD B
		cpu.flags.Carry = 0xFFFF-cpu.getBC() < cpu.getHL()
		cpu.H, cpu.L = splitWord(cpu.getHL() + cpu.getBC())
	case 0x19: // DAD D
		cpu.flags.Carry = 0xFFFF-cpu.getDE() < cpu.getHL()
		cpu.H, cpu.L = splitWord(cpu.getHL() + cpu.getDE())
	case 0x29: // DAD H
		cpu.flags.Carry = 0xFFFF-cpu.getHL() < cpu.getHL()
		cpu.H, cpu.L = splitWord(cpu.getHL() * 2) // Add HL to itself
	case 0x39: // DAD SP
		cpu.flags.Carry = 0xFFFF-cpu.stackPointer < cpu.getHL()
		cpu.H, cpu.L = splitWord(cpu.getHL() + cpu.stackPointer)

	// SUBTRACT
	case 0x90: // SUB B - Subtract register from A
		cpu.sub(cpu.B, NOCARRY)
	case 0x91: // SUB C - Subtract register from A
		cpu.sub(cpu.C, NOCARRY)
	case 0x92: // SUB D - Subtract register from A
		cpu.sub(cpu.D, NOCARRY)
	case 0x93: // SUB E - Subtract register from A
		cpu.sub(cpu.E, NOCARRY)
	case 0x94: // SUB H - Subtract register from A
		cpu.sub(cpu.H, NOCARRY)
	case 0x95: // SUB L - Subtract register from A
		cpu.sub(cpu.L, NOCARRY)
	case 0x96: // SUB M - Subtract memory from A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.sub(readByte, NOCARRY)
	case 0x97: // SUB A - Subtract register from A
		cpu.sub(cpu.A, NOCARRY)
	case 0x98: // SBB B - Subtract register from A with borrow
		cpu.sub(cpu.B, WITHCARRY)
	case 0x99: // SBB C - Subtract register from A with borrow
		cpu.sub(cpu.C, WITHCARRY)
	case 0x9A: // SBB D - Subtract register from A with borrow
		cpu.sub(cpu.D, WITHCARRY)
	case 0x9B: // SBB E - Subtract register from A with borrow
		cpu.sub(cpu.E, WITHCARRY)
	case 0x9C: // SBB H - Subtract register from A with borrow
		cpu.sub(cpu.H, WITHCARRY)
	case 0x9D: // SBB L - Subtract register from A with borrow
		cpu.sub(cpu.L, WITHCARRY)
	case 0x9E: // SBB M - Subtract memory from A with borrow
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.sub(readByte, WITHCARRY)
	case 0x9F: // SBB A - Subtract register from A with borrow
		cpu.sub(cpu.A, WITHCARRY)
	case 0xD6: // SUI - Subtract immediate from A
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.sub(fetchedByte, NOCARRY)
	case 0xDE: // SBI - Subtract immediate from A with borrow
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.sub(fetchedByte, WITHCARRY)

	// LOGICAL
	case 0xA0: // ANA B - AND register with A
		cpu.ana(cpu.B)
	case 0xA1: // ANA C - AND register with A
		cpu.ana(cpu.C)
	case 0xA2: // ANA D - AND register with A
		cpu.ana(cpu.D)
	case 0xA3: // ANA E - AND register with A
		cpu.ana(cpu.E)
	case 0xA4: // ANA H - AND register with A
		cpu.ana(cpu.H)
	case 0xA5: // ANA L - AND register with A
		cpu.ana(cpu.L)
	case 0xA6: // ANA M - AND memory with A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.ana(readByte)
	case 0xA7: // ANA A - AND register with A
		cpu.ana(cpu.A)
	case 0xA8: // XRA B - XOR register with A
		cpu.xra(cpu.B)
	case 0xA9: // XRA C - XOR register with A
		cpu.xra(cpu.C)
	case 0xAA: // XRA D - XOR register with A
		cpu.xra(cpu.D)
	case 0xAB: // XRA E - XOR register with A
		cpu.xra(cpu.E)
	case 0xAC: // XRA H - XOR register with A
		cpu.xra(cpu.H)
	case 0xAD: // XRA L - XOR register with A
		cpu.xra(cpu.L)
	case 0xAE: // XRA M - XOR memory with A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.xra(readByte)
	case 0xAF: // XRA A - XOR register with A
		cpu.xra(cpu.A)
	case 0xB0: // ORA B - OR register with A
		cpu.ora(cpu.B)
	case 0xB1: // ORA C - OR register with A
		cpu.ora(cpu.C)
	case 0xB2: // ORA D - OR register with A
		cpu.ora(cpu.D)
	case 0xB3: // ORA E - OR register with A
		cpu.ora(cpu.E)
	case 0xB4: // ORA H - OR register with A
		cpu.ora(cpu.H)
	case 0xB5: // ORA L - OR register with A
		cpu.ora(cpu.L)
	case 0xB6: // ORA M - OR memory with A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.ora(readByte)
	case 0xB7: // ORA A - OR register with A
		cpu.ora(cpu.A)
	case 0xB8: // CMP B - Compare register with A
		cpu.cmp(cpu.B)
	case 0xB9: // CMP C - Compare register with A
		cpu.cmp(cpu.C)
	case 0xBA: // CMP D - Compare register with A
		cpu.cmp(cpu.D)
	case 0xBB: // CMP E - Compare register with A
		cpu.cmp(cpu.E)
	case 0xBC: // CMP H - Compare register with A
		cpu.cmp(cpu.H)
	case 0xBD: // CMP L - Compare register with A
		cpu.cmp(cpu.L)
	case 0xBE: // CMP M - Compare memory with A
		readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
		if err != nil {
			return err
		}
		cpu.cmp(readByte)
	case 0xBF: // CMP A - Compare register with A
		cpu.cmp(cpu.A)
	case 0xE6: // ANI - AND immediate with A
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.ana(fetchedByte)
	case 0xEE: // XRI - XOR immediate with A
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.xra(fetchedByte)
	case 0xF6: // ORI - OR immediate with A
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.ora(fetchedByte)
	case 0xFE: // CPI - Compare immediate with A
		fetchedByte, err := cpu.fetchByte()
		if err != nil {
			return err
		}
		cpu.cmp(fetchedByte)

	// ROTATE
	case 0x07: // RLC - Rotate A left
		msb := cpu.A >> 7 // Isolate the MSB (bit 7)
		cpu.A <<= 1       // Shift everything one bit to the left
		cpu.A |= msb      // Replace the LSB (bit 0) with the MSB (bit 7)
		cpu.flags.Carry = (msb == 1)
	case 0x0F: // RRC - Rotate A right
		lsb := cpu.A & 1  // Isolate the LSB (bit 0)
		cpu.A >>= 1       // Shift everything one bit to the right
		cpu.A |= lsb << 7 // Replace the MSB (bit 7) with the LSB (bit 0)
		cpu.flags.Carry = (lsb == 1)
	case 0x17: // RAL - Rotate A left through carry
		msb := cpu.A >> 7 // Isolate the MSB (bit 7)
		cpu.A <<= 1       // Shift everything one bit to the left
		if cpu.flags.Carry {
			cpu.A |= 0b0000_0001 // Replace the LSB (bit 0) with the carry flag
		}
		cpu.flags.Carry = (msb == 1)
	case 0x1F: // RAR - Rotate A right through carry
		lsb := cpu.A & 1 // Isolate the LSB (bit 0)
		cpu.A >>= 1      // Shift everything one bit to the right
		if cpu.flags.Carry {
			cpu.A |= 0b1000_0000 // Replace the MSB (bit 7) with the carry flag
		}
		cpu.flags.Carry = (lsb == 1)

	// SPECIALS
	case 0x2F: // CMA - Complement A
		cpu.A = ^cpu.A
	case 0x37: // STC - Set carry
		cpu.flags.Carry = true
	case 0x3F: // CMC - Complement carry
		cpu.flags.Carry = !cpu.flags.Carry
	case 0x27: // DAA - Decimal adjust A
		cpu.daa()

	// INPUT/OUTPUT
	case 0xDB: // IN - Input
		return ErrNotImplemented(opCode)
	case 0xD3: // OUT - Output
		return ErrNotImplemented(opCode)

	// CONTROL
	case 0xFB: // EI - Enable interrupts
		return ErrNotImplemented(opCode)
	case 0xF3: // DI - Disable interrupts
		return ErrNotImplemented(opCode)
	case 0x00: // NOP - No-operation
		// Do nothing
	case 0x76: // HLT - Halt
		cpu.halted = true

	default:
		cpu.halted = true
		return fmt.Errorf("instruction %02X not found", opCode)
	}

	return nil
}

// getFlags returns the current state of the CPU flags packed into a single byte, for use in
// functions such as PUSH PSW.  The flags are ordered from MSB (bit 7) to LSB (bit 0).
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

	var result byte
	for i, flag := range flags {
		if flag {
			result |= 1 << (7 - i)
		}
	}
	return result
}

// setFlags updates the CPU flags based on the provided flags byte.
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

func (cpu *CPU) pushStack(address word) error {
	high, low := splitWord(address)
	err := cpu.Bus.WriteByteAt(cpu.stackPointer-1, high)
	if err != nil {
		return err
	}
	err = cpu.Bus.WriteByteAt(cpu.stackPointer-2, low)
	if err != nil {
		return err
	}
	cpu.stackPointer -= 2
	return nil
}

func (cpu *CPU) popStack() (word, error) {
	low, err := cpu.Bus.ReadByteAt(cpu.stackPointer)
	if err != nil {
		return 0, err
	}
	high, err := cpu.Bus.ReadByteAt(cpu.stackPointer + 1)
	if err != nil {
		return 0, err
	}
	cpu.stackPointer += 2
	return joinBytes(high, low), nil
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
// 1. Adds the accumulator, register value, and the carry-in value, allowing for overflow
// 3. Checks for carry-out and auxiliary carry-out, updating the corresponding flags
// 2. Sets the sign, zero, and parity flags based on the result
// 4. Updates the accumulator with the result
//
// Parameters:
// - register (byte): The value of the register to be added to the accumulator
// - carry (byte): A byte indicating if there is an initial carry-in (used by ADC opcodes)
//
// Examples:
//
//	cpu := &CPU{A: 0x10}
//	cpu.add(0x20, 1)
//	// cpu.A is 0x31
//
//	cpu := &CPU{A: 0x10}
//	cpu.add(0x10, 0)
//	// cpu.A is 0x20
func (cpu *CPU) add(register byte, carry byte) {
	// Calculate the result but capture the overflow by casting to a word (uint16).
	result := word(cpu.A) + word(register) + word(carry)

	// Set the carry flag by checking whether we've got an overflow into bit eight.
	cpu.flags.Carry = result > 0b1111_1111 // (0xFF)

	// Set the auxillary carry flag by checking whether our addition results in a carry-out from bit four.
	auxCarrySum := (cpu.A&0b1111 + register&0b1111 + carry&0b1111)
	cpu.flags.AuxCarry = auxCarrySum > 0b1111 // (0x0F)

	// Set the sign, zero and parity flags, based on the LSB only, given we've already
	// taken note of whether there was a carry-in to bit eight above.
	cpu.setSignZeroParityFlags(byte(result))

	// Return the eight least significant bits (LSB) only
	cpu.A = byte(result)
}

// sub subtracts the value of a register and an optional borrow-in from the accumulator,
// updating the accumulator and the CPU flags accordingly.
//
// This method performs the following steps:
// 1. Subtracts the accumulator, register value, and the borrow-in value, allowing for underflow
// 3. Checks for carry-out and auxiliary carry-out, updating the corresponding flags
// 2. Sets the sign, zero, and parity flags based on the result
// 4. Updates the accumulator with the result
//
// Parameters:
// - register (byte): The value of the register to be subtracted from the accumulator
// - carry (byte): A byte indicating if there is an initial borrow-in (used by ADC opcodes)
//
// Examples:
//
//	cpu := &CPU{A: 0x10}
//	cpu.sub(0x01, 1)
//	// cpu.A is 0x0A
//
//	cpu := &CPU{A: 0x10}
//	cpu.sub(0x03, 0)
//	// cpu.A is 0x0D
func (cpu *CPU) sub(register byte, borrow byte) {
	// Calculate the result but capture the underflow by casting to a word (uint16).
	result := word(cpu.A) - word(register) - word(borrow)

	// Set the carry flag by checking whether we've got an underflow from bit eight.
	cpu.flags.Carry = result > 0b1111_1111 // (0xFF)

	// Set the auxillary carry flag by checking whether our addition results in a carry-in to bit four.
	auxBorrowSum := (cpu.A&0b1111 - register&0b1111 - borrow&0b1111)
	cpu.flags.AuxCarry = auxBorrowSum&0b0001_0000 != 0 // 0b0001_0000 = 0x10

	// Set the sign, zero and parity flags, based on the LSB only, given we've already
	// taken note of whether there was a carry-in to bit eight above.
	cpu.setSignZeroParityFlags(byte(result))

	// Return the eight least significant bits (LSB) only
	cpu.A = byte(result)
}

// ana performs a logical AND with the A register.
//
// Parameters:
// - register (byte): The value of the register to be ANDed with the A register
func (cpu *CPU) ana(register byte) {
	cpu.A = cpu.A & register
	cpu.setSignZeroParityFlags(cpu.A)
	cpu.flags.AuxCarry = false
	cpu.flags.Carry = false
}

// xor performs a logical XOR (exclusive OR) with the A register.
//
// Parameters:
// - register (byte): The value of the register to be XORed with the A register
func (cpu *CPU) xra(register byte) {
	cpu.A = cpu.A ^ register
	cpu.setSignZeroParityFlags(cpu.A)
	cpu.flags.AuxCarry = false
	cpu.flags.Carry = false
}

// ora performs a logical OR with the A register.
//
// Parameters:
// - register (byte): The value of the register to be ORed with the A register
func (cpu *CPU) ora(register byte) {
	cpu.A = cpu.A | register
	cpu.setSignZeroParityFlags(cpu.A)
	cpu.flags.AuxCarry = false
	cpu.flags.Carry = false
}

// daa adjusts the eight-bit value in the accumulator to form two four-bit binary coded decimal digits.
//
// This method performs the following steps:
// 1. If the least significant four bits of the accumulator have a value greater than nine,
// or if the auxiliary carry flag is set, DAA adds six to the accumulator.
// 2. If the most significant four bits of the accumulator have a value greater than nine,
// or if the carry flag is set, DAA adds six to the most significant four bits of the accumulator.
func (cpu *CPU) daa() {
	// Adjust lower nibble
	lowerNibble := cpu.A & 0b0000_1111 // Isolate the four LSBs
	if lowerNibble > 9 || cpu.flags.AuxCarry {
		cpu.flags.AuxCarry = lowerNibble > 9
		cpu.A += 0x06 // Add 6 to the lower nibble of the accumulator
	}

	// Adjust higher nibble
	upperNibble := cpu.A >> 4 // Isolate the four MSBs
	if upperNibble > 9 || cpu.flags.Carry {
		cpu.flags.Carry = upperNibble > 9
		cpu.A += 0x60 // Add 6 to the upper nibble of the accumulator
	}

	cpu.setSignZeroParityFlags(cpu.A)
}

// cmp performs a logical comparison with the A register.
//
// The comparison is performed by internally subtracting the contents of the register
// from the accumulator (leaving both unchanged) and setting the condition bits according to the result.
// We take a temporary copy of the A register first, then perform the subtraction to extra the bits,
// then reset the A register to the temporary copy to preserve its contents.
//
// Parameters:
// - register (byte): The value of the register to be compared with the A register
func (cpu *CPU) cmp(register byte) {
	tempA := cpu.A
	cpu.sub(register, NOCARRY) // We're only interested in the flags
	cpu.A = tempA
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

func (cpu CPU) getBC() word {
	return joinBytes(cpu.B, cpu.C)
}

func (cpu CPU) getDE() word {
	return joinBytes(cpu.D, cpu.E)
}

func (cpu CPU) getHL() word {
	return joinBytes(cpu.H, cpu.L)
}

func (cpu CPU) getAPSW() word {
	return joinBytes(cpu.A, cpu.getFlags())
}

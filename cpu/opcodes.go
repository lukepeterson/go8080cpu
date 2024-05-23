package cpu

type opcodeFunc func(*CPU)

var decode [256]opcodeFunc

// initOpcodeMap is a mapping of Intel 8080 opcodes to their functions.
// Grouped by instruction set group as per "Table 2. Instruction Set Summary",
// in the Intel 8080A 8-BIT N-CHANNEL MICROPROCESSOR datasheet.
func initOpcodeMap() {
	// MOVE, LOAD AND STORE
	decode[0x3E] = (*CPU).MVI_A
	decode[0x06] = (*CPU).MVI_B
	decode[0x0E] = (*CPU).MVI_C
	decode[0x10] = (*CPU).MVI_D
	decode[0x1E] = (*CPU).MVI_E
	decode[0x1A] = (*CPU).MVI_H
	decode[0x2E] = (*CPU).MVI_L

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

func (cpu *CPU) MVI_A() { // 0x3A
	cpu.A = cpu.fetchByte()
}

func (cpu *CPU) MVI_B() { // 0x06
	cpu.B = cpu.fetchByte()
}

func (cpu *CPU) MVI_C() { // 0x0E
	cpu.C = cpu.fetchByte()
}

func (cpu *CPU) MVI_D() { // 0x10
	cpu.D = cpu.fetchByte()
}

func (cpu *CPU) MVI_E() { // 0x1E
	cpu.E = cpu.fetchByte()
}

func (cpu *CPU) MVI_H() { // 0x1A
	cpu.H = cpu.fetchByte()
}

func (cpu *CPU) MVI_L() { // 0x2E
	cpu.L = cpu.fetchByte()
}

func (cpu *CPU) HLT() { // 0x4C
	cpu.halted = true
}

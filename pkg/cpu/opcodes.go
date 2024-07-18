package cpu

import (
	"fmt"
	"math/bits"

	"github.com/lukepeterson/go8080cpu/pkg/types"
)

const (
	NoCarry   = 0
	WithCarry = 1
)

// push 'pushes' the two byte word onto the stack, before decrementing the
// stack pointer by two.
//
// Example:
//
//	// Assuming the stack pointer was 0xFFFF prior to the push instruction:
//	cpu.push(0x1234)
//	// Memory location 0xFFFE (stackPointer - 1) = 0x12
//	// Memory location 0xFFFD (stackPointer - 2) = 0x34
func (cpu *CPU) push(value types.Word) error {
	high, low := splitWord(value)

	err := cpu.Bus.WriteByteAt(cpu.stackPointer-1, high)
	if err != nil {
		return fmt.Errorf("could not write 0x%02X to cpu.stackPointer - 1 (0x%04X): %v", high, cpu.stackPointer-1, err)
	}

	err = cpu.Bus.WriteByteAt(cpu.stackPointer-2, low)
	if err != nil {
		return fmt.Errorf("could not write 0x%02X to cpu.stackPointer - 2 (0x%04X): %v", low, cpu.stackPointer-2, err)
	}

	cpu.stackPointer -= 2
	return nil
}

// pop 'pops' a two byte word from the stack, before incrementing the
// stack pointer by two.
//
// Example:
//
// // Assuming the stack pointer was 0xFFFD prior to the pop instruction, and:
// // Memory location 0xFFFE (stackPointer - 1) = 0x12
// // Memory location 0xFFFD (stackPointer - 2) = 0x34
// poppedValue, _ := cpu.pop()
// // poppedValue = 0x1234
func (cpu *CPU) pop() (types.Word, error) {
	low, err := cpu.Bus.ReadByteAt(cpu.stackPointer)
	if err != nil {
		return 0, fmt.Errorf("could not read byte from cpu.stackPointer (0x%04X): %v", cpu.stackPointer, err)
	}

	high, err := cpu.Bus.ReadByteAt(cpu.stackPointer + 1)
	if err != nil {
		return 0, fmt.Errorf("could not read byte from cpu.stackPointer + 1 (0x%04X): %v", cpu.stackPointer+1, err)
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
	result := types.Word(cpu.A) + types.Word(register) + types.Word(carry)

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

// dad performs a double add to the two byte value stores in the H and L registers
//
// Example:
// // Assume BC = 0x339F and HL = 0xA17B
// cpu.dad(cpu.getBC())
// // HL now contains 0xD51A
func (cpu *CPU) dad(value types.Word) {
	cpu.flags.Carry = 0xFFFF-value < cpu.getHL()
	cpu.H, cpu.L = splitWord(cpu.getHL() + value)
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
	result := types.Word(cpu.A) - types.Word(register) - types.Word(borrow)

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

	// Adjust upper nibble
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
// then reset the A register to the temporary copy to restore its contents.
//
// Parameters:
// - register (byte): The value of the register to be compared with the A register
func (cpu *CPU) cmp(register byte) {
	tempA := cpu.A
	cpu.sub(register, NoCarry) // We're only interested in the flags
	cpu.A = tempA
}

// jmp causes a transfer of program control depending upon the condition being met.
//
// If the condition is true, program execution will continue at the memory location formed by
// concatenating the third byte of the instruction with the second byte of the instruction (as
// instructions are stored little endian, therefore, in reverse).
// If the condition is false, program execution continues at the next instruction.
//
// Parameters:
//   - condition (bool): determines whether to jump to the address specified in the third and
//     second bytes of the instruction.
func (cpu *CPU) jmp(condition bool) error {
	address, err := cpu.fetchWord()
	if err != nil {
		return fmt.Errorf("could not jmp() to address 0x%04X: %v", address, err)
	}

	if condition {
		cpu.programCounter = address
	}

	return nil
}

// call functions similarly to the JMP instruction, however, a return address is also pushed onto
// the stack before jumping to the address specified in the third and second bytes of the instruction.
//
// Parameters:
//   - condition (bool): determines whether to jump to the address specified in the third and
//     second bytes of the instruction.
func (cpu *CPU) call(condition bool) error {
	address, err := cpu.fetchWord()
	if err != nil {
		return err
	}

	if condition {
		err = cpu.push(cpu.programCounter)
		if err != nil {
			return fmt.Errorf("could not call() to address 0x%04X: %v", address, err)
		}
		cpu.programCounter = address
	}

	return nil
}

// ret pops the last address saved on the stack into the program counter, causing a transfer of
// program control to that address.  RET is typically called to return from a subroutine initiated by
// a CALL instruction.
//
// Parameters:
//   - condition (bool): determines whether to return to the address specified in the last two bytes
//     popped off the stack.
func (cpu *CPU) ret(condition bool) error {
	address, err := cpu.pop()
	if err != nil {
		return fmt.Errorf("could not ret() from address 0x%04X: %v", address, err)
	}

	if condition {
		cpu.programCounter = address
	}

	return nil
}

// rst (restart) is a special purpose subroutine jump.
//
// The contents of the program counter are pushed onto the stack for later use by
// a RET instruction.  The program counter is then set to the address parameter.
func (cpu *CPU) rst(address types.Word) error {
	err := cpu.push(cpu.programCounter)
	if err != nil {
		return err
	}

	cpu.programCounter = address
	return nil
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
func joinBytes(high, low byte) types.Word {
	return types.Word(high)<<8 | types.Word(low)
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
func splitWord(address types.Word) (high, low byte) {
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
//	// cpu.flags.Parity is false (0x80 has an odd number of 1's)
func (cpu *CPU) setSignZeroParityFlags(input byte) {
	cpu.flags.Sign = input >= 0b1000_0000
	cpu.flags.Zero = input == 0
	cpu.flags.Parity = bits.OnesCount8(input)%2 == 0 // Check if parity is even
}

// getBC returns a two byte word by joining the B and C registers
func (cpu CPU) getBC() types.Word {
	return joinBytes(cpu.B, cpu.C)
}

// getDE returns a two byte word by joining the D and E registers
func (cpu CPU) getDE() types.Word {
	return joinBytes(cpu.D, cpu.E)
}

// getHL returns a two byte word by joining the H and L registers
func (cpu CPU) getHL() types.Word {
	return joinBytes(cpu.H, cpu.L)
}

// getAWithFlags returns a two byte word by joining the A and flag registers
func (cpu CPU) getAWithFlags() types.Word {
	return joinBytes(cpu.A, cpu.getFlags())
}

// getM returns a byte stored in memory, pointed to by the H and L registers
func (cpu CPU) getM() (byte, error) {
	readByte, err := cpu.Bus.ReadByteAt(cpu.getHL())
	if err != nil {
		return 0, err
	}

	return readByte, nil
}

// setM stores a byte stored in memory, pointed to by the H and L registers
func (cpu *CPU) setM(value byte) error {
	err := cpu.Bus.WriteByteAt(cpu.getHL(), value)
	if err != nil {
		return err
	}

	return nil
}

// ErrNotImplemented is a temporary function to be removed when all instructions are implemented
func errNotImplemented(opCode byte) error {
	return fmt.Errorf("instruction 0x%02X not implemented", opCode)
}

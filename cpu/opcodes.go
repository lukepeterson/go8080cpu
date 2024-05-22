package cpu

import "fmt"

type opcodeFunc func(*CPU)

var opcodeMap [256]opcodeFunc

func initOpcodeMap() {
	opcodeMap[0x00] = (*CPU).NOP
	opcodeMap[0x3A] = (*CPU).LDA
	opcodeMap[0x4C] = (*CPU).HLT
}

func (cpu *CPU) NOP() { // 0x00
	fmt.Println("NOP")
}

func (cpu *CPU) LDA() { // 0x3A
	lowOrderByte := uint16(cpu.Memory[cpu.programCounter])
	highOrderByte := uint16(cpu.Memory[cpu.programCounter+1])
	cpu.A = cpu.Memory[(highOrderByte<<8)|lowOrderByte]
	cpu.programCounter += 2
}

func (cpu *CPU) HLT() { // 0x4C
	fmt.Println("HALT!")
	cpu.halted = true
}

// var decodeMnemonic = map[string]byte{
// 	"NOP": 0x00,
// 	"LDA": 0x3A,
// 	"HLT": 0x4C,
// }

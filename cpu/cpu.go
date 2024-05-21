package cpu

import (
	"fmt"
	"log"
	"time"
)

const (
	FETCH = iota
	DECODE
	EXECUTE
)

const (
	NOP = iota
	LDA
	ADD
	SUB
	HLT
)

type Register struct {
	Data [8]bool
}

func (r *Register) Load(input byte) []bool {
	data := make([]bool, 8)
	for i := 7; i >= 0; i-- {
		data[7-i] = (input>>i)&1 == 1
	}

	return data
}

type Bus struct {
	Data [8]bool
}

func NewBus() *Bus {
	return &Bus{
		Data: [8]bool{},
	}
}

func (cpu *CPU) Fetch() {
	fmt.Println("do fetch")
	cpu.Controller.instruction = cpu.Memory[cpu.programCounter]
	cpu.programCounter++

	fmt.Printf("    cpu.Controller.instruction: %v\n", cpu.Controller.instruction)
	fmt.Println("--------------------")
}

func (cpu *CPU) Decode() {
	fmt.Println("do decode")

	cpu.Controller.opcode = cpu.Controller.instruction
	if cpu.Controller.opcode == LDA || cpu.Controller.opcode == ADD || cpu.Controller.opcode == SUB { // instructions that have an operand
		cpu.Controller.operand = cpu.Memory[cpu.programCounter]
		cpu.programCounter++
	}

	fmt.Printf("    cpu.Controller.opcode: %v\n", cpu.Controller.opcode)
	fmt.Printf("    cpu.Controller.operand: %v\n", cpu.Controller.operand)
	fmt.Println("--------------------")
}

func (cpu *CPU) Execute() {
	fmt.Println("do execute")

	switch cpu.Controller.opcode {
	case HLT:
		cpu.halted = true
	case LDA:
		cpu.registerA.Load(cpu.Controller.operand)
	case ADD:
	case SUB:
	default:
		log.Fatalf("invalid instruction: %v", cpu.Controller.opcode)
		cpu.halted = true
	}

	fmt.Println("--------------------")
}

func (cpu *CPU) Run() {

	for !cpu.halted {
		time.Sleep(100 * time.Millisecond)
		switch cpu.Controller.State {
		case FETCH:
			cpu.Fetch()
			cpu.Controller.Next()
		case DECODE:
			cpu.Decode()
			cpu.Controller.Next()
		case EXECUTE:
			cpu.Execute()
			cpu.Controller.Next()
		}
	}

}

type Controller struct {
	State       int
	instruction byte
	opcode      byte
	operand     byte
}

func (c *Controller) Next() {
	switch c.State {
	case FETCH:
		c.State = DECODE
	case DECODE:
		c.State = EXECUTE
	case EXECUTE:
		c.State = FETCH
	}
}

type CPU struct {
	// alu            ALU
	registerA      Register
	registerB      Register
	Controller     Controller
	Memory         [256]byte
	programCounter byte
	Bus            *Bus
	halted         bool
}

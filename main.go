package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lukepeterson/go8080assembler/assembler"
	"github.com/lukepeterson/go8080cpu/cpu"
)

func main() {

	goCPU := cpu.NewCPU(
		1*time.Millisecond,
		cpu.NewMemory(32),
	)

	code := `
		MVI A, 02h
		INR A
		IN 55h
		LDA 3400h
		INR A
		DCR H	
		NOP
		HLT
	`

	assembler := &assembler.Assembler{}
	err := assembler.Assemble(code)
	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range assembler.ByteCode {
		fmt.Printf("%02X ", instruction)
	}

	goCPU.Load(assembler.ByteCode)
	runErr := goCPU.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}

}

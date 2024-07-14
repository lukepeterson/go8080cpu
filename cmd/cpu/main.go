package main

import (
	"fmt"
	"log"

	"github.com/lukepeterson/go8080assembler/assembler"
	"github.com/lukepeterson/go8080cpu/pkg/cpu"
	"github.com/lukepeterson/go8080cpu/pkg/memory"
)

func main() {

	goCPU := cpu.New()
	goCPU.Bus = memory.New(32)
	// goCPU.DebugMode = true

	code := `
		INR A
		DCR H
		INR B
		DCR C
		INR D
		DCR E
		INR L
		HLT
	`

	assembler := assembler.New()
	err := assembler.Assemble(code)
	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range assembler.ByteCode {
		fmt.Printf("%02X ", instruction)
	}
	fmt.Println("")

	goCPU.Load(assembler.ByteCode)
	runErr := goCPU.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
}

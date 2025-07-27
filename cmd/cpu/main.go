package main

import (
	"fmt"
	"log"

	"github.com/lukepeterson/go8080assembler/pkg/assembler"
	"github.com/lukepeterson/go8080cpu/pkg/cpu"
)

func main() {

	goCPU := cpu.New()
	goCPU.DebugMode = true

	input := `
		INR A
		DCR H
		INR B
		DCR C
		INR D
		DCR E
		INR L
		HLT
		`

	asm := assembler.New(input)
	bytecode, err := asm.Assemble()
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range bytecode {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	goCPU.Load(bytecode)
	err = goCPU.Run()
	if err != nil {
		log.Fatal(err)
	}
}

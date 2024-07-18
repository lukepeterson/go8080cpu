package main

import (
	"fmt"
	"log"

	"github.com/lukepeterson/go8080assembler/assembler"
	"github.com/lukepeterson/go8080cpu/pkg/cpu"
)

func main() {

	goCPU := cpu.New()
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

	asm := assembler.New()
	err := asm.Assemble(code)
	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range asm.ByteCode {
		fmt.Printf("%02X ", instruction)
	}
	fmt.Println("")

	goCPU.Load(asm.ByteCode)
	err = goCPU.Run()
	if err != nil {
		log.Fatal(err)
	}
}

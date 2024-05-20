package main

import (
	"fmt"

	"github.com/lukepeterson/gocpu/gates"
)

func main() {
	fmt.Printf("gates.And(true, false): %v\n", gates.And(true, false))
	fmt.Printf("gates.Or(true, false): %v\n", gates.Or(true, false))
	fmt.Printf("gates.Not(true): %v\n", gates.Not(true))
	fmt.Printf("gates.Nand(true, false): %v\n", gates.Nand(true, false))
}

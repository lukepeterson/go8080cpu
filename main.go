package main

import (
	"github.com/lukepeterson/gocpu/cpu"
)

func main() {

	goCPU := cpu.CPU{
		Bus: cpu.NewBus(),
	}
	goCPU.Controller.State = 0 // TODO: Initialise in the cpu package

	goCPU.Memory[0] = cpu.LDA
	goCPU.Memory[1] = 9
	goCPU.Memory[2] = cpu.ADD
	goCPU.Memory[3] = 10
	goCPU.Memory[4] = cpu.ADD
	goCPU.Memory[5] = 11
	goCPU.Memory[6] = cpu.SUB
	goCPU.Memory[7] = 12
	goCPU.Memory[8] = cpu.HLT

	goCPU.Memory[9] = 1
	goCPU.Memory[10] = 2
	goCPU.Memory[11] = 3
	goCPU.Memory[12] = 4

	goCPU.Run()
}

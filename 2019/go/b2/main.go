package main

import (
	"fmt"

	"camille.codes/aoc/b2/pc"
)

func main() {
	const opLength = 4
	const endProgram = 99

	c := &pc.Computer{}
	c.InitializeMemory(12, 2)

	c.ProcessInstructions(opLength, endProgram)
	fmt.Println(c.Output)
}

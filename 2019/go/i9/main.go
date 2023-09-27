package main

import (
	"bufio"

	"camille.codes/aoc/i9/pc"
	"camille.codes/aoc/utils"
)

func main() {
	file := utils.GetFile("i9/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boostProgram := utils.GetProgram(scanner)

	computer := &pc.Computer{}
	inputSignal := 1
	computer.InitializeMemory(boostProgram, inputSignal)
	computer.ProcessInstructions()

	// computer.MemoryDump()
}

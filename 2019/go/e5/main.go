package main

import (
	"bufio"
	"fmt"

	"camille.codes/aoc/e5/pc"
	"camille.codes/aoc/utils"
)

func main() {
	file := utils.GetFile("e5/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	computer := &pc.Computer{}

	fmt.Println("Part 1")
	inputInstruction := 1
	computer.LoadProgram(scanner, inputInstruction)
	computer.ProcessInstructions()

	fmt.Println("Part 2")
	inputInstruction = 5
	computer.LoadProgram(scanner, inputInstruction)
	computer.ProcessInstructions()
}

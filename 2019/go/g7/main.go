package main

import (
	"bufio"
	"fmt"

	"camille.codes/aoc/g7/pc"
	"camille.codes/aoc/utils"
)

func main() {
	file := utils.GetFile("g7/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amplifierController := utils.GetProgram(scanner)

	phaseSettings := []int{0, 1, 2, 3, 4}
	permutations := utils.GetPermutations(phaseSettings)
	fmt.Println("permutations:", permutations)

	finalOutputSignal := 0
	for _, permutation := range permutations {
		amplifierA := &pc.Computer{}
		amplifierB := &pc.Computer{}
		amplifierC := &pc.Computer{}
		amplifierD := &pc.Computer{}
		amplifierE := &pc.Computer{}

		phaseSetting := permutation[0]
		inputInstruction := 0
		amplifierA.InitializeMemory(amplifierController, phaseSetting, inputInstruction)
		amplifierA.ProcessInstructions()

		phaseSetting = permutation[1]
		amplifierB.InitializeMemory(amplifierController, phaseSetting, amplifierA.OutputSignal)
		amplifierB.ProcessInstructions()

		phaseSetting = permutation[2]
		amplifierC.InitializeMemory(amplifierController, phaseSetting, amplifierB.OutputSignal)
		amplifierC.ProcessInstructions()

		phaseSetting = permutation[3]
		amplifierD.InitializeMemory(amplifierController, phaseSetting, amplifierC.OutputSignal)
		amplifierD.ProcessInstructions()

		phaseSetting = permutation[4]
		amplifierE.InitializeMemory(amplifierController, phaseSetting, amplifierD.OutputSignal)
		amplifierE.ProcessInstructions()

		if amplifierE.OutputSignal > finalOutputSignal {
			finalOutputSignal = amplifierE.OutputSignal
		}
	}

	fmt.Println("part 1:", finalOutputSignal)
}

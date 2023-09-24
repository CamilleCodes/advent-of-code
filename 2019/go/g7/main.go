package main

import (
	"bufio"
	"fmt"

	"camille.codes/aoc/g7/pc"
	"camille.codes/aoc/utils"

	"github.com/Ramshackle-Jamathon/go-quickPerm"
)

func main() {
	fmt.Println("part 1:", runAmplifiers())
	fmt.Println("part 2:", loopAmplifiers())
}

func runAmplifiers() int {
	file := utils.GetFile("g7/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amplifierController := utils.GetProgram(scanner)

	phaseSettings := []int{0, 1, 2, 3, 4}
	finalOutputSignal := 0

	amplifierA := &pc.Computer{}
	amplifierB := &pc.Computer{}
	amplifierC := &pc.Computer{}
	amplifierD := &pc.Computer{}
	amplifierE := &pc.Computer{}

	// Permutations reference:
	// https://www.baeldung.com/cs/array-generate-all-permutations#quickperm-algorithm
	for permutation := range quickPerm.GeneratePermutationsInt(phaseSettings) {
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

	return finalOutputSignal
}

func loopAmplifiers() int {
	file := utils.GetFile("g7/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amplifierController := utils.GetProgram(scanner)

	phaseSettings := []int{9, 8, 7, 6, 5}
	finalOutputSignal := 0

	amplifierA := &pc.Computer{}
	amplifierB := &pc.Computer{}
	amplifierC := &pc.Computer{}
	amplifierD := &pc.Computer{}
	amplifierE := &pc.Computer{}

	for permutation := range quickPerm.GeneratePermutationsInt(phaseSettings) {
		amplifierA.LoadSoftware(amplifierController, permutation[0])
		amplifierB.LoadSoftware(amplifierController, permutation[1])
		amplifierC.LoadSoftware(amplifierController, permutation[2])
		amplifierD.LoadSoftware(amplifierController, permutation[3])
		amplifierE.LoadSoftware(amplifierController, permutation[4])

		amplifierA.SetInputSignal(0)

		for amplifierE.IsRunning {
			amplifierA.ProcessUntilSuspended()

			amplifierB.SetInputSignal(amplifierA.OutputSignal)
			amplifierB.ProcessUntilSuspended()

			amplifierC.SetInputSignal(amplifierB.OutputSignal)
			amplifierC.ProcessUntilSuspended()

			amplifierD.SetInputSignal(amplifierC.OutputSignal)
			amplifierD.ProcessUntilSuspended()

			amplifierE.SetInputSignal(amplifierD.OutputSignal)
			amplifierE.ProcessUntilSuspended()

			amplifierA.SetInputSignal(amplifierE.OutputSignal)
		}

		if amplifierE.OutputSignal > finalOutputSignal {
			finalOutputSignal = amplifierE.OutputSignal
		}

	}

	return finalOutputSignal
}

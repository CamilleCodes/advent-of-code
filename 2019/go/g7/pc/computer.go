package pc

import (
	"fmt"

	"camille.codes/aoc/utils"
)

type Computer struct {
	IsRunning    bool
	IsSuspended  bool
	OutputSignal int

	memory         []int
	instructionPtr int
	phaseSetting   int
	inputSignal    int
	inputCounter   int
}

// InitializeMemory sets the values for the computer program memory (for part 1)
func (c *Computer) InitializeMemory(software []int, phaseSetting, inputSignal int) {
	c.memory = utils.Copy(software)
	c.phaseSetting = phaseSetting
	c.inputSignal = inputSignal
	c.inputCounter = 0
	c.instructionPtr = 0
}

// ProcessInstructions runs the computer program (for part 1)
func (c *Computer) ProcessInstructions() {
	c.IsRunning = true
	for c.IsRunning && c.instructionPtr < len(c.memory) {
		c.instructionPtr = c.processOpCode(c.instructionPtr)
	}
}

// LoadSoftware sets the values for the computer program memory (for part 2)
func (c *Computer) LoadSoftware(software []int, phaseSetting int) {
	c.IsRunning = true
	c.memory = utils.Copy(software)
	c.phaseSetting = phaseSetting
	c.inputCounter = 0
	c.instructionPtr = 0
}

// SetInputSignal sets the input signal for the computer (for part 2)
func (c *Computer) SetInputSignal(inputSignal int) {
	c.inputSignal = inputSignal
}

// ProcessUntilSuspended runs the computer program until it is suspended or halted (for part 2)
func (c *Computer) ProcessUntilSuspended() {
	c.IsSuspended = false
	for c.IsRunning && !c.IsSuspended {
		c.instructionPtr = c.processOpCode(c.instructionPtr)
	}
}

// Process the opcode and advance to the next instruction
func (c *Computer) processOpCode(ptr int) int {
	opCodeLength, instruction := 2, c.memory[ptr]
	opCode, modes := pullOpCodeAndModes(instruction, opCodeLength)

	switch opCode {
	case 1:
		return c.add(modes, ptr)
	case 2:
		return c.multiply(modes, ptr)
	case 3:
		return c.input(ptr)
	case 4:
		return c.output(modes, ptr)
	case 5:
		return c.jumpIfTrue(modes, ptr)
	case 6:
		return c.jumpIfFalse(modes, ptr)
	case 7:
		return c.lessThan(modes, ptr)
	case 8:
		return c.equals(modes, ptr)
	case 99:
		// halt
		c.IsRunning = false
		return len(c.memory)
	default:
		fmt.Println("invalid opcode:", opCode)
		return len(c.memory)
	}
}

// Convert an integer to a slice of integers
func convertIntToSlice(num int) []int {
	slice := make([]int, 4)

	for i := 3; num > 0; i-- {
		slice[i] = num % 10
		num /= 10
	}

	return slice
}

// Convert a slice of integers to an integer
func convertToNumber(slice []int) int {
	var num int

	for i := 0; i < len(slice); i++ {
		num = num*10 + slice[i]
	}

	return num
}

// Pull the opcode from the instructions
func pullOpCode(instructions []int, opCodeLength int) int {
	opCode := instructions[len(instructions)-opCodeLength:]
	return convertToNumber(opCode)
}

// Pull the parameter modes from the instructions
func pullParameterModes(instructions []int, opCodeLength int) []int {
	return instructions[:len(instructions)-opCodeLength]
}

// Pull the opcode and the parameter modes from the instructions
func pullOpCodeAndModes(instruction, opCodeLength int) (int, []int) {
	fullInstructions := convertIntToSlice(instruction)
	return pullOpCode(fullInstructions, opCodeLength),
		pullParameterModes(fullInstructions, opCodeLength)
}

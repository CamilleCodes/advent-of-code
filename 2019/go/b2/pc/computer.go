package pc

type Computer struct {
	memory []int
	Output int
}

// InitializeMemory sets the values for the computer program memory
func (c *Computer) InitializeMemory(inputOne, inputTwo int) {
	c.memory = []int{
		1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3,
		2, 6, 1, 19, 1, 19, 5, 23, 2, 9, 23, 27, 1, 5, 27, 31,
		1, 5, 31, 35, 1, 35, 13, 39, 1, 39, 9, 43, 1, 5, 43, 47,
		1, 47, 6, 51, 1, 51, 13, 55, 1, 55, 9, 59, 1, 59, 13, 63,
		2, 63, 13, 67, 1, 67, 10, 71, 1, 71, 6, 75, 2, 10, 75, 79,
		2, 10, 79, 83, 1, 5, 83, 87, 2, 6, 87, 91, 1, 91, 6, 95,
		1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 9, 107, 1, 10, 107, 111,
		2, 111, 13, 115, 1, 10, 115, 119, 1, 10, 119, 123, 2, 13, 123, 127,
		2, 6, 127, 131, 1, 13, 131, 135, 1, 135, 2, 139, 1, 139, 6, 0,
		99, 2, 0, 14, 0,
	}

	c.memory[1] = inputOne
	c.memory[2] = inputTwo
}

// getParameters pulls the input values needed to perform the operation
// indicated by the current instruction opcode
func (c *Computer) getParameters(index int) (int, int) {
	indexOne := c.memory[index+1]
	indexTwo := c.memory[index+2]

	return c.memory[indexOne], c.memory[indexTwo]
}

// setInstructionResult sets the result value of the operation indicated
// by the current instruction opcode
func (c *Computer) setInstructionResult(index, result int) {
	resultIndex := c.memory[index+3]
	c.memory[resultIndex] = result

}

// Process the instruction opcode, perform the operation indicated,
// and assign the value at memory address 0 as the computer output
func (c *Computer) processOpcode(opcode, index int) {
	a, b := c.getParameters(index)

	// Perform instruction operation (add or multiply)
	result := _op[opcode](a, b)
	c.setInstructionResult(index, result)

	c.Output = c.memory[0]
}

// processInstructions iterates over the puzzle input and processes
// the instructions (opcode + 3 parameters)
func (c *Computer) ProcessInstructions(opLength, endProgram int) {
	for index, opcode := range c.memory {
		if index%opLength != 0 {
			continue
		}

		if opcode == endProgram {
			break
		}

		c.processOpcode(opcode, index)
	}
}

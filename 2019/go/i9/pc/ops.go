package pc

import "fmt"

// TODO: Update to return a pointer, instead of values
func (c *Computer) determineParameter(mode, instructionPtr, offset int) int {
	switch mode {
	case 0:
		// Position mode: return the address of the parameter
		address := c.memory[instructionPtr+offset]
		return address
	case 1:
		// Immediate mode: return the address of the parameter
		return instructionPtr + offset
	case 2:
		// Relative mode: return the address of the parameter
		return c.relativeBase + c.memory[instructionPtr+offset]
	default:
		panic("Invalid mode")
	}
}

// Get the parameters for the instruction
func (c *Computer) getParameters(modes []int, addressPtr int) (int, int, int) {
	var p1, p2, p3 int

	switch len(modes) {
	case 1:
		p1 = c.determineParameter(modes[0], addressPtr, 1)
	case 2:
		p1 = c.determineParameter(modes[1], addressPtr, 1)
		p2 = c.determineParameter(modes[0], addressPtr, 2)
	case 3:
		p1 = c.determineParameter(modes[2], addressPtr, 1)
		p2 = c.determineParameter(modes[1], addressPtr, 2)
		p3 = c.determineParameter(modes[0], addressPtr, 3)
	default:
		panic("Invalid mode length")
	}

	return p1, p2, p3
}

// Add the first two parameter values and save
// the result in the position given by the third parameter
func (c *Computer) add(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)
	c.memory[p3] = c.memory[p1] + c.memory[p2]

	// Advance the pointer by the instruction length
	ptr += 4
	return ptr
}

// Multiply the first two parameter values and save
// the result in the position given by the third parameter
func (c *Computer) multiply(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)
	c.memory[p3] = c.memory[p1] * c.memory[p2]

	ptr += 4
	return ptr
}

// Input takes a single integer as input and saves it
// to the position given by its only parameter
func (c *Computer) input(modes []int, ptr int) int {
	p1, _, _ := c.getParameters(modes, ptr)
	c.memory[p1] = c.inputSignal

	ptr += 2
	return ptr
}

// Outputs the value of it's only parameter
func (c *Computer) output(modes []int, ptr int) int {
	p1, _, _ := c.getParameters(modes, ptr)

	fmt.Println("Output:", c.memory[p1])
	c.OutputSignal = c.memory[p1]

	ptr += 2
	return ptr
}

// Jump if true checks if the first parameter is non-zero
// If it is non-zero, the instruction pointer is set to the
// value from the second parameter, otherwise it simply
// moves to the next instruction
func (c *Computer) jumpIfTrue(modes []int, ptr int) int {
	p1, p2, _ := c.getParameters(modes, ptr)

	if c.memory[p1] != 0 {
		ptr = c.memory[p2]
	} else {
		ptr += 3
	}

	return ptr
}

// Jump if false checks if the first parameter is zero
// If it is zero, the instruction pointer is set to the
// value from the second parameter, otherwise it simply
// moves to the next instruction
func (c *Computer) jumpIfFalse(modes []int, ptr int) int {
	p1, p2, _ := c.getParameters(modes, ptr)

	if c.memory[p1] == 0 {
		ptr = c.memory[p2]
	} else {
		ptr += 3
	}

	return ptr
}

// Less than checks if the first parameter is less than
// the second parameter. If it is, it stores 1 in the
// position given by the third parameter. Otherwise, it
// stores 0 and then moves to the next instruction
func (c *Computer) lessThan(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)

	if c.memory[p1] < c.memory[p2] {
		c.memory[p3] = 1
	} else {
		c.memory[p3] = 0
	}

	ptr += 4
	return ptr

}

// Equals checks if the first parameter is equal to
// the second parameter. If it is, it stores 1 in the
// position given by the third parameter. Otherwise, it
// stores 0 and then moves to the next instruction
func (c *Computer) equals(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)

	if c.memory[p1] == c.memory[p2] {
		c.memory[p3] = 1
	} else {
		c.memory[p3] = 0
	}

	ptr += 4
	return ptr
}

// UpdateRelativeBase adjusts the relative base by the value of its only parameter
func (c *Computer) updateRelativeBase(modes []int, ptr int) int {
	p1, _, _ := c.getParameters(modes, ptr)
	c.relativeBase += c.memory[p1]

	ptr += 2
	return ptr
}

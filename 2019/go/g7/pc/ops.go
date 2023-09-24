package pc

// Get the parameters for the instruction
func (c *Computer) getParameters(modes []int, addressPtr int) (int, int, int) {
	var p1, p2 int
	var p3 = c.memory[addressPtr+3]

	if modes[len(modes)-1] == 1 {
		p1 = c.memory[addressPtr+1]
	} else {
		p1 = c.memory[c.memory[addressPtr+1]]
	}

	if modes[len(modes)-2] == 1 {
		p2 = c.memory[addressPtr+2]
	} else {
		p2 = c.memory[c.memory[addressPtr+2]]
	}

	return p1, p2, p3
}

// Add the first two parameter values and save
// the result in the position given by the third parameter
func (c *Computer) add(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)
	c.memory[p3] = p1 + p2

	// Advance the pointer by the instruction length
	ptr += 4
	return ptr
}

// Multiply the first two parameter values and save
// the result in the position given by the third parameter
func (c *Computer) multiply(modes []int, ptr int) int {
	p1, p2, p3 := c.getParameters(modes, ptr)
	c.memory[p3] = p1 * p2

	ptr += 4
	return ptr
}

// Input takes a single integer as input and saves it
// to the position given by its only parameter
func (c *Computer) input(ptr int) int {
	parameterAddress := ptr + 1
	savePosition := c.memory[parameterAddress]

	if c.inputCounter != 0 {
		c.memory[savePosition] = c.inputSignal
	} else {
		c.memory[savePosition] = c.phaseSetting
		c.inputCounter++
	}

	ptr += 2
	return ptr
}

// Outputs the value of it's only parameter (and suspends the computer for part 2)
func (c *Computer) output(modes []int, ptr int) int {
	var output int
	if modes[len(modes)-1] == 1 {
		output = c.memory[ptr+1]
	} else {
		output = c.memory[c.memory[ptr+1]]
	}

	c.OutputSignal = output
	c.IsSuspended = true

	ptr += 2
	return ptr
}

// Jump if true checks if the first parameter is non-zero
// If it is non-zero, the instruction pointer is set to the
// value from the second parameter, otherwise it simply
// moves to the next instruction
func (c *Computer) jumpIfTrue(modes []int, ptr int) int {
	p1, p2, _ := c.getParameters(modes, ptr)

	if p1 != 0 {
		ptr = p2
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

	if p1 == 0 {
		ptr = p2
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

	if p1 < p2 {
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

	if p1 == p2 {
		c.memory[p3] = 1
	} else {
		c.memory[p3] = 0
	}

	ptr += 4
	return ptr
}

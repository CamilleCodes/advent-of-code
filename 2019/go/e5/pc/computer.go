package pc

import (
	"fmt"
	"log"
)

type Computer struct {
	memory []int
	Output int
}

// InitializeMemory sets the values for the computer program memory
func (c *Computer) InitializeMemory(input []int) {
	c.memory = input
}

// ProcessInstructions runs the computer program
func (c *Computer) ProcessInstructions(input, endProgram int) {
	parameterLength := 0
	for address, instruction := range c.memory {
		if parameterLength > 0 {
			parameterLength--
			continue
		}

		if instruction == endProgram {
			fmt.Println("halt program")
			break
		}

		slice := convertToSlice(instruction)
		opCode := pullOpCode(slice)
		modes := pullParameterModes(slice)

		switch opCode {
		case 1:
			p1, p2, p3 := c.getParameters(modes, address)

			result := add(p1, p2)
			c.memory[p3] = result

			parameterLength = 3
		case 2:
			p1, p2, p3 := c.getParameters(modes, address)

			result := multiply(p1, p2)
			c.memory[p3] = result

			parameterLength = 3

		case 3:
			// Takes a single integer as input and saves it
			// to the position given by its only parameter
			c.memory[c.memory[address+1]] = input

			parameterLength = 1

		case 4:
			// Outputs the value of it's only parameter
			var output int
			if modes[len(modes)-1] == 1 {
				output = c.memory[address+1]
			} else {
				output = c.memory[c.memory[address+1]]
			}

			c.Output = output

			fmt.Println("output:", c.Output)

			parameterLength = 1

		default:
			log.Fatal("invalid opcode")
		}
	}
}

func (c *Computer) getParameters(modes []int, address int) (int, int, int) {
	var p1, p2 int
	var p3 = c.memory[address+3]

	if modes[len(modes)-1] == 1 {
		p1 = c.memory[address+1]
	} else {
		p1 = c.memory[c.memory[address+1]]
	}

	if modes[len(modes)-2] == 1 {
		p2 = c.memory[address+2]
	} else {
		p2 = c.memory[c.memory[address+2]]
	}

	return p1, p2, p3
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func convertToSlice(num int) []int {
	slice := make([]int, 4)

	for i := 3; num > 0; i-- {
		slice[i] = num % 10
		num /= 10
	}

	return slice
}

func convertToNumber(slice []int) int {
	var num int

	for i := 0; i < len(slice); i++ {
		num = num*10 + slice[i]
	}

	return num
}

func pullOpCode(slice []int) int {
	code := slice[len(slice)-2:]
	return convertToNumber(code)
}

func pullParameterModes(slice []int) []int {
	return slice[:len(slice)-2]
}

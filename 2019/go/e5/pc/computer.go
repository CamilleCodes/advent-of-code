package pc

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"camille.codes/aoc/utils"
)

type Computer struct {
	memory           []int
	inputInstruction int
}

// [ ] TODO: file drive with different modes of reading the data (floppy drive, disk drive, etc.)
// load function for each
// bring back diagnostic program

// LoadProgram loads the program into memory and sets the input instruction
func (c *Computer) LoadProgram(scanner *bufio.Scanner, inputInstruction int) {
	scanner.Scan()
	stringInput := strings.Split(scanner.Text(), ",")

	var input []int
	for _, v := range stringInput {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Println("error converting string to int")
			log.Fatal(err)
		}

		input = append(input, val)
	}

	c.memory = utils.Copy(input)
	c.inputInstruction = inputInstruction
}

// ProcessInstructions runs the computer program
func (c *Computer) ProcessInstructions() {
	for instructionPtr := 0; instructionPtr < len(c.memory); {
		instructionPtr = c.processOpCode(instructionPtr)
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

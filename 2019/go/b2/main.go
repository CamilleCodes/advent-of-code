package main

import "fmt"

var _op = map[int]func(int, int) int{
	1: add,
	2: multiply,
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

type puzzle struct {
	input []int
}

// Pull input values needed for the add or multiply operation
func (p *puzzle) pullInputValues(index int) (int, int) {
	indexOne := p.input[index+1]
	indexTwo := p.input[index+2]

	return p.input[indexOne], p.input[indexTwo]
}

// Set the result value of the add or multiply operation
func (p *puzzle) setOpResult(index, result int) {
	resultIndex := p.input[index+3]
	p.input[resultIndex] = result
}

// Process the opcode (1, 2, or 99) and return the current value
// at position 0 of the puzzle input
func (p *puzzle) processOpcode(opcode, index int) int {
	a, b := p.pullInputValues(index)

	// Perform operation (add or multiply)
	result := _op[opcode](a, b)
	p.setOpResult(index, result)

	return p.input[0]
}

func main() {
	const opLength = 4
	const endProgram = 99

	puzz := &puzzle{
		input: []int{
			1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3,
			2, 6, 1, 19, 1, 19, 5, 23, 2, 9, 23, 27, 1, 5, 27, 31,
			1, 5, 31, 35, 1, 35, 13, 39, 1, 39, 9, 43, 1, 5, 43, 47,
			1, 47, 6, 51, 1, 51, 13, 55, 1, 55, 9, 59, 1, 59, 13, 63,
			2, 63, 13, 67, 1, 67, 10, 71, 1, 71, 6, 75, 2, 10, 75, 79,
			2, 10, 79, 83, 1, 5, 83, 87, 2, 6, 87, 91, 1, 91, 6, 95,
			1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 9, 107, 1, 10, 107, 111,
			2, 111, 13, 115, 1, 10, 115, 119, 1, 10, 119, 123, 2, 13, 123, 127,
			2, 6, 127, 131, 1, 13, 131, 135, 1, 135, 2, 139, 1, 139, 6, 0,
			99, 2, 0, 14, 0,
		},
	}

	var result int
	for index, opcode := range puzz.input {
		if index%opLength == 0 {
			if opcode == endProgram {
				break
			}

			result = puzz.processOpcode(opcode, index)
		}
	}

	fmt.Println(result)
}

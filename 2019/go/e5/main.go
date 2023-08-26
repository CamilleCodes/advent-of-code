package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"camille.codes/aoc/e5/pc"
	"camille.codes/aoc/utils"
)

func getProgram(scanner *bufio.Scanner) []int {
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

	return input
}

func main() {
	file := utils.GetFile("e5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	diagnosticProgram := getProgram(scanner)

	c := &pc.Computer{}
	c.InitializeMemory(diagnosticProgram)

	const inputInstruction = 1
	const endProgram = 99
	c.ProcessInstructions(inputInstruction, endProgram)

	fmt.Println("diagnostic code:", c.Output)
}

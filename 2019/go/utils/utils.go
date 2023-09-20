package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFile opens the input file and returns a pointer to the file
func GetFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file")
		log.Fatal(err)
	}

	return file
}

// GetProgram reads the input file and returns a slice of integers
func GetProgram(scanner *bufio.Scanner) []int {
	scanner.Scan()
	stringInput := strings.Split(scanner.Text(), ",")

	var program []int
	for _, v := range stringInput {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Println("error converting string to int")
			log.Fatal(err)
		}

		program = append(program, val)
	}

	return program
}

// Abs returns the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// Min returns the lowest value in a slice of integers
func Min(nums []int) int {
	lowest := nums[0]
	for _, num := range nums {
		if num < lowest {
			lowest = num
		}
	}

	return lowest
}

// Copy returns a copy of a slice of integers
func Copy(nums []int) []int {
	var copy []int
	copy = append(copy, nums...)
	return copy
}

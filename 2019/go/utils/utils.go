package utils

import (
	"fmt"
	"log"
	"os"
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

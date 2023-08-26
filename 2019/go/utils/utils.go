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

// Loop backwards through the slice of the number
// If the index as we loop backwards through the slice,
// if the index is greater than or equal to length of the slice - 2
// then we can assume we are still looking at the OP codes
// and every number that we get, we can stick into a new array
// called OP codes

// As it goes backwards through the slice, we get to the ones
// that are less, ex. length - 2, are all parameter modes

// Make two arrays, one to hold the op codes and one to hold the params

// Convert the slice back to a single number
// multiply it to get back to the number (by 1 and then 10)

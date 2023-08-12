package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part 1

// getFile opens the input file and returns a pointer to the file
func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	return file
}

// getModuleMass scans the input file and returns the mass of the
// module mass listed in the input file
func getModuleMass(input string) int {
	mass, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error converting string to int")
		log.Fatal(err)
	}

	return mass
}

// getFuelRequired takes a mass and returns the amount of fuel required
// for a module of that mass
func getFuelRequired(mass int) int {
	return mass/3 - 2
}

func main() {
	file := getFile("a1/input.txt")
	defer file.Close()

	var fuelTotal = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass := getModuleMass(scanner.Text())
		fuel := getFuelRequired(mass)
		fuelTotal += fuel
	}

	fmt.Println(fuelTotal)
}

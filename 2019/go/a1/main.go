package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part 1 & 2

// getFile opens the input file and returns a pointer to the file
func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	return file
}

// getModuleMass returns the module mass listed in the input text
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

// getAdditionalFuelRequired calculates the total amount of fuel
// required for the fuel itself
func getAdditionalFuelRequired(fuel int) int {
	additionalFuel := 0
	for {
		fuel = getFuelRequired(fuel)
		if fuel <= 0 {
			break
		}

		additionalFuel += fuel
	}

	return additionalFuel
}

func main() {
	file := getFile("a1/input.txt")
	defer file.Close()

	fuelTotal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass := getModuleMass(scanner.Text())
		fuel := getFuelRequired(mass)
		fuelTotal += fuel

		additionalFuel := getAdditionalFuelRequired(fuel)
		fuelTotal += additionalFuel
	}

	fmt.Println(fuelTotal)
}

package main

import (
	"bufio"
	"fmt"
	"strings"

	"camille.codes/aoc/c3/coord"
	"camille.codes/aoc/utils"
)

func getWire(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func main() {
	file := utils.GetFile("c3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	coaxial := getWire(scanner)
	fiberOptic := getWire(scanner)
	coaxialPath := coord.CreatePath(coaxial)
	fiberOpticPath := coord.CreatePath(fiberOptic)

	distances := coaxialPath.GetIntersectionDistances(fiberOpticPath)
	steps := coaxialPath.GetIntersectionSteps(fiberOpticPath)

	fmt.Println(utils.Min(distances))
	fmt.Println(utils.Min(steps))
}

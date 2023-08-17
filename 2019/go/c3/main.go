package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"camille.codes/aoc/c3/coord"
	"camille.codes/aoc/utils"
)

func getPath(scanner *bufio.Scanner) coord.Path {
	scanner.Scan()

	directions := strings.Split(scanner.Text(), ",")

	path := createPath(directions)

	return path
}

func createPath(directions []string) coord.Path {
	path := coord.Path{}
	startPoint := coord.Point{X: 0, Y: 0}

	for _, location := range directions {
		direction := location[0]
		displacement, err := strconv.Atoi(location[1:])
		if err != nil {
			fmt.Println("error converting string to int")
			log.Fatal(err)
		}

		startPoint = path.AddPoints(direction, displacement, startPoint)
	}

	return path
}

func getLowestDistance(distances []int) int {
	sort.SliceStable(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	return distances[0]
}

func main() {
	file := utils.GetFile("c3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pathA := getPath(scanner)
	pathB := getPath(scanner)

	distances := pathA.GetIntersectionDistances(pathB)
	fmt.Println(getLowestDistance(distances))
}

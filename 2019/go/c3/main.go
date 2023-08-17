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

	path := coord.Path{}
	point := coord.Point{X: 0, Y: 0}

	for _, location := range directions {
		direction := location[0]
		displacement, err := strconv.ParseInt(location[1:], 10, 64)
		if err != nil {
			fmt.Println("error converting string to int")
			log.Fatal(err)
		}

		var i int64
		for i = 0; i < displacement; i++ {
			// TODO: Put this in a function
			switch direction {
			case 'U':
				point.Y++
			case 'D':
				point.Y--
			case 'L':
				point.X--
			case 'R':
				point.X++
			default:
				fmt.Println("invalid direction")
			}

			path = append(path, point)
		}
	}

	return path
}

func calculateDistance(pathA, pathB coord.Path) int64 {
	distances := pathA.GetIntersectionDistances(pathB)

	sort.Slice(distances, func(i, j int) bool {
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

	fmt.Println(calculateDistance(pathA, pathB))

	coord.Something()
}

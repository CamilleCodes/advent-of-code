package coord

import (
	"fmt"
	"log"
	"strconv"

	"camille.codes/aoc/utils"
)

type Point struct {
	X int
	Y int
}

// Path is a map of points to the number of steps it took to get there
type Path map[Point]int

func (p *Point) addDisplacement(direction byte) {
	switch direction {
	case 'U':
		p.Y++
	case 'D':
		p.Y--
	case 'L':
		p.X--
	case 'R':
		p.X++
	default:
		panic("invalid direction")
	}
}

func CreatePath(directions []string) Path {
	path := make(Path)
	point := Point{0, 0}
	currentStep := 0

	for _, location := range directions {
		direction := location[0]
		displacement, err := strconv.Atoi(location[1:])
		if err != nil {
			fmt.Println("error converting string to int")
			log.Fatal(err)
		}

		for i := 0; i < displacement; i++ {
			currentStep++
			point.addDisplacement(direction)
			path[point] = currentStep
		}
	}

	return path
}

// Part 1
func (path Path) GetIntersectionDistances(wirePath Path) []int {
	distances := make([]int, 0)

	for point := range wirePath {
		if path.contains(point) {
			distance := utils.Abs(point.X) + utils.Abs(point.Y)
			distances = append(distances, distance)
		}
	}

	return distances
}

// Part 2
func (path Path) GetIntersectionSteps(wirePath Path) []int {
	steps := make([]int, 0)

	for point := range wirePath {
		if path.contains(point) {
			steps = append(steps, path[point]+wirePath[point])
		}
	}

	return steps
}

func (path Path) contains(point Point) bool {
	_, ok := path[point]
	return ok
}

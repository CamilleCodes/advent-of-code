package coord

import (
	"fmt"

	"camille.codes/aoc/utils"
)

type Point struct {
	X int64
	Y int64
}

type Path []Point

func Something() {
	pointA := Point{X: 0, Y: 0}
	pointB := Point{X: 0, Y: 0}

	test := map[Point]int64{
		pointA: 1,
	}

	_, ok := test[pointB]
	if ok {
		fmt.Println("hello, I am in the map")
	}
}

func (path Path) GetIntersectionDistances(wirePath Path) []int64 {
	distances := make([]int64, 0)

	for _, point := range wirePath {
		if path.contains(point) {
			distance := utils.Abs(point.X) + utils.Abs(point.Y)
			distances = append(distances, distance)
		}
	}

	return distances
}

func (path Path) contains(point Point) bool {
	for _, location := range path {
		if point == location {
			return true
		}
	}

	return false
}

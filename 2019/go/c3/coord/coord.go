package coord

import (
	"camille.codes/aoc/utils"
)

type Point struct {
	X int
	Y int
}

func (point *Point) addDisplacement(direction byte) {
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
		panic("invalid direction")
	}
}

// TODO: Make this a map of points to the number of steps for part 2
type Path map[Point]int

func (path Path) AddPoints(
	direction byte,
	displacement int,
	point Point,
) Point {
	for i := 0; i < displacement; i++ {
		point.addDisplacement(direction)
		// TODO: Add steps for part 2
		path[point] = 0
	}

	return point
}

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

func (path Path) contains(point Point) bool {
	_, ok := path[point]
	return ok
}

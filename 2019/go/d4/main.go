package main

import (
	"fmt"

	"camille.codes/aoc/d4/password"
)

func main() {
	start, end, countPt1, countPt2 := 284639, 748759, 0, 0

	for start <= end {
		pw := password.ConvertToSlice(start)

		if password.IsValid(pw) {
			countPt1++
		}

		if password.IsStillValid(pw) {
			countPt2++
		}

		start++
	}

	fmt.Println(countPt1)
	fmt.Println(countPt2)
}

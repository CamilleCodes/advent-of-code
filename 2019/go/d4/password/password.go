package password

// Part 1: A password is valid if it has two adjacent digits that
// are the same and, from left to right, the digits never decrease
func IsValid(password []int) bool {
	adjacent := false

	for i := 1; i < len(password); i++ {
		if password[i-1] > password[i] {
			return false
		}

		if password[i-1] == password[i] {
			adjacent = true
		}
	}

	return adjacent
}

// Part 2: Doubles can no longer have repeatedness
func IsStillValid(password []int) bool {
	adjacent := false
	numAdjacent := 0

	for i := 1; i < len(password); i++ {
		// Digits never decrease
		if password[i-1] > password[i] {
			return false
		}

		// Doubles aren't repeated
		if password[i-1] == password[i] {
			numAdjacent++
		} else {
			if numAdjacent == 1 {
				adjacent = true
			}

			// Reset the counter
			numAdjacent = 0
		}
	}

	return adjacent || numAdjacent == 1
}

// Convert a 6 digit number to a slice of ints
func ConvertToSlice(num int) []int {
	slice := make([]int, 6)

	for i := 5; num > 0; i-- {
		slice[i] = num % 10
		num /= 10
	}

	return slice
}

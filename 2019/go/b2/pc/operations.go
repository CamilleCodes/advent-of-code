package pc

var _op = map[int]func(int, int) int{
	1: add,
	2: multiply,
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

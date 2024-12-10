package utils

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func MapTwoSlices(a, b []int, fn func(int, int) int) []int {
	if len(a) != len(b) {
		panic("Slices must be of the same size.")
	}

	result := make([]int, len(a))
	for i := range a {
		result[i] = fn(a[i], b[i])
	}
	return result
}

func Reduce(initial int, slice []int, fn func(int, int) int) int {
	result := initial
	for _, value := range slice {
		result = fn(result, value)
	}
	return result
}

func Filter(initial int, slice []int, fn func(int, int) bool) []int {
	result := []int{}
	for _, value := range slice {
		if fn(initial, value) {
			result = append(result, value)
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

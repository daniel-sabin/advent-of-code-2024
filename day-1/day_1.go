package day_1

import (
	"fmt"

	puzzledata "github.com/daniel-sabin/advent-of-code-2024/puzzle-data"
	"github.com/daniel-sabin/advent-of-code-2024/utils"
)

func Step1() (int, error) {
	a, b := puzzledata.GetPuzzleData()

	// Result part 1
	differences := utils.MapTwoSlices(a, b, func(x, y int) int {
		return utils.Abs(x - y)
	})
	fmt.Println("result:", utils.Reduce(0, differences, func(x, y int) int {
		return x + y
	}))

	// Result part 2
	result := 0
	for _, value := range a {
		result += value * len(utils.Filter(value, b, func(x, y int) bool {
			return x == y
		}))
	}
	fmt.Println("result:", result)

	return result, nil
}

func Step2() (int, error) {
	a, b := puzzledata.GetPuzzleData()
	// Result part 1
	differences := utils.MapTwoSlices(a, b, func(x, y int) int {
		return utils.Abs(x - y)
	})
	fmt.Println("result:", utils.Reduce(0, differences, func(x, y int) int {
		return x + y
	}))

	// Result part 2
	result := 0
	for _, value := range a {
		result += value * len(utils.Filter(value, b, func(x, y int) bool {
			return x == y
		}))
	}
	fmt.Println("result:", result)

	return result, nil
}

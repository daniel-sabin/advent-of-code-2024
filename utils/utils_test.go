package utils

import "testing"

func TestReduce(t *testing.T) {
    sum := func(a, b int) int { return a + b }
    result := reduce(0, []int{1, 2, 3, 4, 5}, sum)
    expected := 15
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
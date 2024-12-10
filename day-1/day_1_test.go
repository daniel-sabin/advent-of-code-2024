package day_1_test

import (
	"testing"

	day_1 "github.com/daniel-sabin/advent-of-code-2024/day-1"
)

func TestMain(t *testing.T) {

	t.Parallel()

	t.Run("Step 1", func(t *testing.T) {

		actual, _ := day_1.Step1()

		if actual != 2285373 {
			t.Errorf("step 1 error %d", actual)
		}

	})
}

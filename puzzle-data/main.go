package puzzledata

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/daniel-sabin/advent-of-code-2024/utils"
)

func GetPuzzleData() ([]int, []int) {
	path := "./puzzle-data/day_1.input"
	absPath, _ := filepath.Abs(path)
	fmt.Println("Absolute Path:", absPath)
	f, err := os.Open(absPath)
	utils.Check(err)
	defer f.Close()

	a := []int{}
	b := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")
		val, err := strconv.Atoi(split[0])
		utils.Check(err)

		val2, err := strconv.Atoi(split[1])
		utils.Check(err)
		a = append(a, val)
		b = append(b, val2)
	}
	sort.Ints(a)
	sort.Ints(b)
	return a, b
}

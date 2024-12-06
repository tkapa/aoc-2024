package rednose

import (
	"fmt"
	"math"
)

func RednoseSafeReports(reports [][]int) int {
	safeReports := len(reports)

	x := 0
	for x < len(reports) {
		fmt.Println(reports[x])

		isIncreasing := true
		y := 0
		for y < len(reports[x])-1 {
			value := reports[x][y]
			nextValue := reports[x][y+1]

			diff := value - nextValue

			if y == 0 && diff < 0 {
				isIncreasing = false
			}

			fmt.Println(diff)
			y++

			if (isIncreasing && diff < 0) ||
				(!isIncreasing && diff > 0) ||
				(math.Abs(float64(diff)) > 3) ||
				diff == 0 {
				safeReports--
				break
			}
		}

		x++
	}

	return safeReports
}

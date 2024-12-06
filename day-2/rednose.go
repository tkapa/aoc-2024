package rednose

import (
	"fmt"
	"math"
	"slices"
)

func RednoseSafeReports(reports [][]int) int {
	safeReports := len(reports)

	x := 0
	for x < len(reports) {
		thisRep := reports[x]
		repSafe := RednoseIsReportSafe(thisRep)

		if !repSafe {
			safeReports--
		}

		x++
	}

	return safeReports
}

func RednoseSafeReportsWithDampening(reports [][]int) int {

	safeReports := len(reports)

	x := 0
	for x < len(reports) {
		thisRep := reports[x]
		repSafe := RednoseIsReportSafe(thisRep)

		if !repSafe {
			for y := range len(thisRep) {
				cpRep := make([]int, len(thisRep))
				copy(cpRep, thisRep)
				cpRep = slices.Delete(cpRep, y, y+1)
				dmpnSafe := RednoseIsReportSafe(cpRep)

				if dmpnSafe {
					repSafe = true
					break
				}
			}
		}

		if !repSafe {
			safeReports--
		}

		x++
	}

	return safeReports
}

func RednoseIsReportSafe(report []int) bool {
	isIncreasing := true
	y := 0

	for y <= len(report)-2 {
		nextIdx := y + 1

		if nextIdx > len(report)-1 {
			y++
			continue
		}

		value := report[y]
		nextValue := report[nextIdx]
		diff := value - nextValue

		// Diff must be at least one and no more than 3
		if DiffIsLessThan3AndAbove0(diff) {
			return false
		}

		// All Increasing / Decreasing
		// Negative Diff Should Indicate an Increase (3->4 = Diff -1)
		// Positive Diff Should Indicate a Decrease (4->3 = Diff 1)
		if y == 0 && diff > 0 {
			isIncreasing = false
		}

		if isIncreasing && !HasIncreased(diff) {
			// fmt.Println("No Incr:", y, nextIdx, diff, isIncreasing)
			return false
		} else if !isIncreasing && HasIncreased(diff) {
			// fmt.Println("No Decr:", y, nextIdx, diff, isIncreasing)
			return false
		}

		y++
	}

	fmt.Println(report)
	return true
}

func DiffIsLessThan3AndAbove0(diff int) bool {
	return diff == 0 || (math.Abs(float64(diff)) > 3)
}

func HasIncreased(diff int) bool {
	return diff < 0
}

package rednose

import (
	"fmt"
	"math"
	"slices"
)

func RednoseSafeReports(reports [][]int, tryDampen bool) int {
	safeReports := len(reports)

	x := 0
	for x < len(reports) {
		thisRep := reports[x]
		repSafe, probIdx1 := RednoseIsReportSafe(thisRep)

		if !repSafe && tryDampen {
			nRep1 := make([]int, len(thisRep))
			copy(nRep1, thisRep)
			nRep1 = slices.Delete(nRep1, probIdx1, probIdx1+1)
			nRep2 := make([]int, len(thisRep))
			copy(nRep2, thisRep)
			nRep2 = slices.Delete(nRep2, probIdx1+1, probIdx1+2)

			rep1Safe, _ := RednoseIsReportSafe(nRep1)
			rep2Safe, _ := RednoseIsReportSafe(nRep2)

			repSafe = rep1Safe || rep2Safe

			fmt.Println("Skip Index ", probIdx1)
			fmt.Println(nRep1, nRep2)
			fmt.Printf("Rep Safe: %v\n", repSafe)
		}

		if !repSafe {
			safeReports--
		}

		x++
	}

	return safeReports
}

func RednoseIsReportSafe(report []int) (bool, int) {
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
			return false, y
		}

		// All Increasing / Decreasing
		// Negative Diff Should Indicate an Increase (3->4 = Diff -1)
		// Positive Diff Should Indicate a Decrease (4->3 = Diff 1)
		if y == 0 && diff > 0 {
			isIncreasing = false
		}

		if isIncreasing && !HasIncreased(diff) {
			// fmt.Println("No Incr:", y, nextIdx, diff, isIncreasing)
			return false, y
		} else if !isIncreasing && HasIncreased(diff) {
			// fmt.Println("No Decr:", y, nextIdx, diff, isIncreasing)
			return false, y
		}

		y++
	}

	fmt.Println(report)
	return true, -1
}

func DiffIsLessThan3AndAbove0(diff int) bool {
	return diff == 0 || (math.Abs(float64(diff)) > 3)
}

func HasIncreased(diff int) bool {
	return diff < 0
}

package rednose

import (
	"fmt"
	"math"
)

func RednoseSafeReports(reports [][]int, tryDampen bool) int {
	safeReports := len(reports)

	x := 0
	for x < len(reports) {
		fmt.Println(reports[x])

		repSafe, probIdx1, probIdx2 := RednoseIsReportSafe(reports[x], -1)

		if !repSafe && tryDampen {
			fixedRep1, _, _ := RednoseIsReportSafe(reports[x], probIdx1)
			fixedRep2, _, _ := RednoseIsReportSafe(reports[x], probIdx2)

			repSafe = fixedRep1 || fixedRep2
		}

		if !repSafe {
			safeReports--
		}

		x++
	}

	return safeReports
}

func RednoseIsReportSafe(report []int, skipIdx int) (bool, int, int) {
	isIncreasing := true
	y := 0

	if skipIdx > len(report) {
		return false, -1, -1
	}

	fmt.Printf("Rep: %v\n", report)

	if skipIdx > -1 {
		fmt.Printf("Skip Index: %v\n", skipIdx)
	}

	for y <= len(report)-2 {
		value := report[y]
		nextIdx := y + 1

		if y == skipIdx-1 {
			nextIdx = skipIdx + 1
		} else if y == skipIdx {
			y++
			continue
		}

		if nextIdx > len(report)-1 {
			y++
			continue
		}

		nextValue := report[nextIdx]

		diff := value - nextValue

		if (y == 0 || (skipIdx == 0 && y == 1)) && diff < 0 {
			isIncreasing = false
		}

		y++

		if (isIncreasing && diff < 0) ||
			(!isIncreasing && diff > 0) ||
			(math.Abs(float64(diff)) > 3) ||
			diff == 0 {
			return false, y, nextIdx
		}
	}

	return true, -1, -1
}

package hysteria

// Historian Hysteria
// Example Data
// 3 4
// 4 3
// 2 5
// 1 3
// 3 9
// 3 3
// Step 1, Order
// Step 2, Compare
// Step 3, Add Diff
//
// List 1: 1,2,3,3,3,4
// List 2: 3,3,3,4,5,9
// Diff Abs(x - y): 2,1,0,1,2,5
// Total Diff: 11

import (
	"fmt"
	"math"
	"slices"
)

func Hysteria(list1 []int, list2 []int) int {

	// Sort List
	fmt.Println("--- Sorting List ---")
	slices.Sort(list1)
	slices.Sort(list2)
	fmt.Println("List 1: ", list1)
	fmt.Println("List 2: ", list2)

	var totalDiff float64
	i := 0
	maxInt := len(list1)
	for i < maxInt {
		diff := list1[i] - list2[i]
		absDiff := math.Abs(float64(diff))
		totalDiff += absDiff
		i++
	}

	fmt.Printf("--- Total Diff: %v ---", totalDiff)
	return int(totalDiff)
}

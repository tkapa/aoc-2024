package hysteria

import (
	"fmt"
	"math"
	"slices"
)

func Hysteria(list1 []int, list2 []int) int {

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
	return int(totalDiff)
}

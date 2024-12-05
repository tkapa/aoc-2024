package hysteria

import (
	"fmt"
	"math"
	"slices"
)

func HysteriaCompare(list1 []int, list2 []int) int {

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

func HysteriaSimilarity(list1, list2 []int) int {

	list1Len := len(list1)
	appearances := make([][2]int, list1Len)
	similarityValue := 0

	x := 0
	// TODO: Copy Value for Duplicates
	for x < list1Len {
		value := list1[x]
		appears := 0

		y := 0
		for y < len(list2) {
			value2 := list2[y]

			if value == value2 {
				appears++
			}
			y++
		}

		appearances[x] = [2]int{value, appears}
		similarityValue += (value * appears)
		x++
	}

	return similarityValue
}

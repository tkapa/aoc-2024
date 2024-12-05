package hysteria

import (
	"fmt"
	"testing"
)

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

func TestHysteriaCompareSimple(t *testing.T) {

	// Arrange
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	// Act
	diff := HysteriaCompare(list1, list2)

	// Assert
	if diff != 11 {
		t.Fatalf("Diff was not correct, expected 11, got: %v", diff)
	}
}

func TestHysteriaSimilaritySimple(t *testing.T) {
	// Arrange
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	// Act
	similarity := HysteriaSimilarity(list1, list2)

	// Assert
	if similarity != 31 {
		t.Fatalf("Diff was not correct, expected 31, got: %v", similarity)
	}
}

func TestHysteriaReadFile(t *testing.T) {
	// Arrange
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	fileName := "./test.txt"

	// Act
	res1, res2 := HysteriaReadFile(fileName)

	// Assert
	if !compareSlice(list1, res1) {
		t.Fatalf("List 1 not equal to Res 1")
	}
	if !compareSlice(list2, res2) {
		t.Fatalf("List 2 not equal to Res 2")
	}
}

func TestDay1HysteriaCompare(t *testing.T) {
	inputFile := "./input.txt"

	list1, list2 := HysteriaReadFile(inputFile)

	diff := HysteriaCompare(list1, list2)

	if diff != 2113135 {
		t.Fatalf("List diff is not correct! Expected: 2113135 Got: %v", diff)
	}
}

func TestDay1HysteriaSimilarity(t *testing.T) {
	inputFile := "./input.txt"

	list1, list2 := HysteriaReadFile(inputFile)

	similarity := HysteriaSimilarity(list1, list2)

	// Assert
	if similarity != 19097157 {
		t.Fatalf("List similarity is not correct! Expected: 19097157 Got: %v", similarity)
	}
}

func TestCompareSlice(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{3, 4, 2, 1, 3, 3}

	if !compareSlice(list1, list2) {
		t.Fatalf("Lists not equal")
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("Length Doesn't Mat")
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

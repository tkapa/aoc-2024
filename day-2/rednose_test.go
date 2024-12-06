package rednose

import (
	"fmt"
	"testing"
)

/*
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
*/

func TestRednoseSafeReportsSimple(t *testing.T) {
	// Arrange
	matrix := make([][]int, 6)

	matrix[0] = []int{7, 6, 4, 2, 1}
	matrix[1] = []int{1, 2, 7, 8, 9}
	matrix[2] = []int{9, 7, 6, 2, 1}
	matrix[3] = []int{1, 3, 2, 4, 5}
	matrix[4] = []int{8, 6, 4, 4, 1}
	matrix[5] = []int{1, 3, 6, 7, 9}

	// Act
	safeReps := RednoseSafeReports(matrix)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestRednoseSafeReportsTestInputSimple(t *testing.T) {
	// Arrange
	matrix := RednoseReadFile("./test.txt")

	// Act
	safeReps := RednoseSafeReports(matrix)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestDay2RednoseSafeReports(t *testing.T) {
	inputFile := "./input.txt"

	matrix := RednoseReadFile(inputFile)

	safeReps := RednoseSafeReports(matrix)

	fmt.Printf("Safe Reports: %v", safeReps)

	if safeReps != 432 {
		t.Fatalf("Expected 432 Safe Reports, Found: %v", safeReps)
	}
}

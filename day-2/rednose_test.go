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

func TestRednoseSafeReportsSimpleNoDampening(t *testing.T) {
	// Arrange
	matrix := make([][]int, 6)

	matrix[0] = []int{7, 6, 4, 2, 1}
	matrix[1] = []int{1, 2, 7, 8, 9}
	matrix[2] = []int{9, 7, 6, 2, 1}
	matrix[3] = []int{1, 3, 2, 4, 5}
	matrix[4] = []int{8, 6, 4, 4, 1}
	matrix[5] = []int{1, 3, 6, 7, 9}

	// Act
	safeReps := RednoseSafeReports(matrix, false)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestRednoseSafeReportsTestInputSimpleNoDampening(t *testing.T) {
	// Arrange
	matrix := RednoseReadFile("./test.txt")

	// Act
	safeReps := RednoseSafeReports(matrix, false)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestDay2RednoseSafeReportsNoDampening(t *testing.T) {
	inputFile := "./input.txt"

	matrix := RednoseReadFile(inputFile)

	safeReps := RednoseSafeReports(matrix, false)

	fmt.Printf("Safe Reports: %v", safeReps)

	if safeReps != 432 {
		t.Fatalf("Expected 432 Safe Reports, Found: %v", safeReps)
	}
}

func TestDay2RednoseSafeReportsWithDampening(t *testing.T) {
	inputFile := "./input.txt"

	matrix := RednoseReadFile(inputFile)

	safeReps := RednoseSafeReports(matrix, true)

	fmt.Printf("Safe Reports: %v", safeReps)
}

func TestRednoseIsReportSafeNoDampening(t *testing.T) {
	// Arrange
	r0 := []int{7, 6, 4, 2, 1}
	r1 := []int{1, 2, 7, 8, 9}
	r2 := []int{9, 7, 6, 2, 1}
	r3 := []int{1, 3, 2, 4, 5}
	r4 := []int{8, 6, 4, 4, 1}
	r5 := []int{1, 3, 6, 7, 9}

	isR0Safe, _, _ := RednoseIsReportSafe(r0, -1)
	isR1Safe, _, _ := RednoseIsReportSafe(r1, -1)
	isR2Safe, _, _ := RednoseIsReportSafe(r2, -1)
	isR3Safe, _, _ := RednoseIsReportSafe(r3, -1)
	isR4Safe, _, _ := RednoseIsReportSafe(r4, -1)
	isR5Safe, _, _ := RednoseIsReportSafe(r5, -1)

	if !isR0Safe || isR1Safe || isR2Safe || isR3Safe || isR4Safe || !isR5Safe {
		fmt.Printf("R0 Expects: true, Is: %v\n", isR0Safe)
		fmt.Printf("R1 Expects: false, Is: %v\n", isR1Safe)
		fmt.Printf("R2 Expects: false, Is: %v\n", isR2Safe)
		fmt.Printf("R3 Expects: false, Is: %v\n", isR3Safe)
		fmt.Printf("R4 Expects: false, Is: %v\n", isR4Safe)
		fmt.Printf("R5 Expects: true, Is: %v\n", isR5Safe)
		t.Fatal("Something is wrong")
	}
}

func TestRednoseIsReportSafeWithDampening(t *testing.T) {
	// Arrange
	r0 := []int{7, 6, 4, 2, 1}
	r1 := []int{1, 2, 7, 8, 9}
	r2 := []int{9, 7, 6, 2, 1}
	r3 := []int{1, 3, 2, 4, 5}
	r4 := []int{8, 6, 4, 4, 1}
	r5 := []int{1, 3, 6, 7, 9}

	r3RemIdx := 1
	r4RemIdx := 3

	isR0Safe, _, _ := RednoseIsReportSafe(r0, -1)
	isR1Safe, _, _ := RednoseIsReportSafe(r1, -1)
	isR2Safe, _, _ := RednoseIsReportSafe(r2, -1)
	isR3Safe, _, _ := RednoseIsReportSafe(r3, r3RemIdx)
	isR4Safe, _, _ := RednoseIsReportSafe(r4, r4RemIdx)
	isR5Safe, _, _ := RednoseIsReportSafe(r5, -1)

	if !isR0Safe || isR1Safe || isR2Safe || !isR3Safe || !isR4Safe || !isR5Safe {
		fmt.Printf("R0 Expects: true, Is: %v\n", isR0Safe)
		fmt.Printf("R1 Expects: false, Is: %v\n", isR1Safe)
		fmt.Printf("R2 Expects: false, Is: %v\n", isR2Safe)
		fmt.Printf("R3 Expects: false, Is: %v\n", isR3Safe)
		fmt.Printf("R4 Expects: false, Is: %v\n", isR4Safe)
		fmt.Printf("R5 Expects: true, Is: %v\n", isR5Safe)
		t.Fatal("Something is wrong")
	}
}

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
	safeReps := RednoseSafeReports(matrix)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestRednoseSafeReportsSimpleWithDampening(t *testing.T) {
	// Arrange
	matrix := make([][]int, 6)

	matrix[0] = []int{7, 6, 4, 2, 1}
	matrix[1] = []int{1, 2, 7, 8, 9}
	matrix[2] = []int{9, 7, 6, 2, 1}
	matrix[3] = []int{1, 3, 2, 4, 5}
	matrix[4] = []int{8, 6, 4, 4, 1}
	matrix[5] = []int{1, 3, 6, 7, 9}

	// Act
	safeReps := RednoseSafeReportsWithDampening(matrix)

	// Assert
	if safeReps != 4 {
		t.Fatalf("Expected 4 Safe Reports, Found: %v", safeReps)
	}
}

func TestRednoseSafeReportsTestInputSimpleNoDampening(t *testing.T) {
	// Arrange
	matrix := RednoseReadFile("./test.txt")

	// Act
	safeReps := RednoseSafeReports(matrix)

	// Assert
	if safeReps != 2 {
		t.Fatalf("Expected 2 Safe Reports, Found: %v", safeReps)
	}
}

func TestDay2RednoseSafeReportsNoDampening(t *testing.T) {
	inputFile := "./input.txt"

	matrix := RednoseReadFile(inputFile)

	safeReps := RednoseSafeReports(matrix)

	fmt.Printf("Safe Reports: %v", safeReps)

	if safeReps != 432 {
		t.Fatalf("Expected 432 Safe Reports, Found: %v", safeReps)
	}
}

func TestDay2RednoseSafeReportsWithDampening(t *testing.T) {
	inputFile := "./input.txt"

	matrix := RednoseReadFile(inputFile)

	safeReps := RednoseSafeReportsWithDampening(matrix)

	fmt.Printf("Safe Reports: %v", safeReps)

	if safeReps != 488 {
		t.Fatalf("Expected 432 Safe Reports, Found: %v", safeReps)
	}
}

// func zzTestRednoseIsReportSafeNoDampening(t *testing.T) {
// 	// Arrange
// 	r0 := []int{7, 6, 4, 2, 1}
// 	r1 := []int{1, 2, 7, 8, 9}
// 	r2 := []int{9, 7, 6, 2, 1}
// 	r3 := []int{1, 3, 2, 4, 5}
// 	r4 := []int{8, 6, 4, 4, 1}
// 	r5 := []int{1, 3, 6, 7, 9}
// 	r6 := []int{3, 3, 4, 5, 6}

// 	isR0Safe, _ := RednoseIsReportSafe(r0)
// 	isR1Safe, _ := RednoseIsReportSafe(r1)
// 	isR2Safe, _ := RednoseIsReportSafe(r2)
// 	isR3Safe, _ := RednoseIsReportSafe(r3)
// 	isR4Safe, _ := RednoseIsReportSafe(r4)
// 	isR5Safe, _ := RednoseIsReportSafe(r5)
// 	isR6Safe, _ := RednoseIsReportSafe(r6)

// 	if !isR0Safe || isR1Safe || isR2Safe || isR3Safe || isR4Safe || !isR5Safe || isR6Safe {
// 		fmt.Printf("R0 Expects: true, Is: %v\n", isR0Safe)
// 		fmt.Printf("R1 Expects: false, Is: %v\n", isR1Safe)
// 		fmt.Printf("R2 Expects: false, Is: %v\n", isR2Safe)
// 		fmt.Printf("R3 Expects: false, Is: %v\n", isR3Safe)
// 		fmt.Printf("R4 Expects: false, Is: %v\n", isR4Safe)
// 		fmt.Printf("R5 Expects: true, Is: %v\n", isR5Safe)
// 		fmt.Printf("R6 Expects: false, Is: %v\n", isR6Safe)
// 		t.Fatal("Something is wrong")
// 	}
// }

// func zzTestIsReportSafeIncrease(t *testing.T) {

// 	r6 := []int{3, 3, 4, 5, 6}

// 	isR6Safe, _ := RednoseIsReportSafe(r6, 0)

// 	if !isR6Safe {
// 		t.Fatal("Reee")
// 	}
// }

// func zzTestIsReportSafeDecreased(t *testing.T) {

// 	r7 := []int{9, 9, 8, 7, 6}

// 	isR7Safe, _ := RednoseIsReportSafe(r7)

// 	if !isR7Safe {
// 		t.Fatal("Reee")
// 	}
// }

// func zzTestIsReportSafeBadEndIncrease(t *testing.T) {

// 	r7 := []int{4, 5, 7, 8, 9, 9}

// 	isR7Safe, _ := RednoseIsReportSafe(r7, 5)

// 	if !isR7Safe {
// 		t.Fatal("Reee")
// 	}
// }

// func zzTestIsReportSafeBadEndDecrease(t *testing.T) {

// 	r7 := []int{4, 5, 7, 8, 9, 9}

// 	isR7Safe, _ := RednoseIsReportSafe(r7)

// 	if !isR7Safe {
// 		t.Fatal("Reee")
// 	}
// }

// func zzTestRednoseIsReportSafeWithDampening(t *testing.T) {
// 	// Arrange
// 	r0 := []int{7, 6, 4, 2, 1}
// 	r1 := []int{1, 2, 7, 8, 9}
// 	r2 := []int{9, 7, 6, 2, 1}
// 	r3 := []int{1, 3, 2, 4, 5}
// 	r4 := []int{8, 6, 4, 4, 1}
// 	r5 := []int{1, 3, 6, 7, 9}
// 	r6 := []int{3, 3, 4, 5, 6}
// 	r7 := []int{9, 9, 8, 7, 6}

// 	isR0Safe, _ := RednoseIsReportSafe(r0)
// 	isR1Safe, _ := RednoseIsReportSafe(r1)
// 	isR2Safe, _ := RednoseIsReportSafe(r2)
// 	isR3Safe, _ := RednoseIsReportSafe(r3)
// 	isR4Safe, _ := RednoseIsReportSafe(r4)
// 	isR5Safe, _ := RednoseIsReportSafe(r5)
// 	isR6Safe, _ := RednoseIsReportSafe(r6)
// 	isR7Safe, _ := RednoseIsReportSafe(r7)

// 	if !isR0Safe || isR1Safe || isR2Safe || !isR3Safe || !isR4Safe || !isR5Safe || !isR6Safe || !isR7Safe {
// 		fmt.Printf("R0 Expects: true, Is: %v\n", isR0Safe)
// 		fmt.Printf("R1 Expects: false, Is: %v\n", isR1Safe)
// 		fmt.Printf("R2 Expects: false, Is: %v\n", isR2Safe)
// 		fmt.Printf("R3 Expects: true, Is: %v\n", isR3Safe)
// 		fmt.Printf("R4 Expects: true, Is: %v\n", isR4Safe)
// 		fmt.Printf("R5 Expects: true, Is: %v\n", isR5Safe)
// 		fmt.Printf("R6 Expects: true, Is: %v\n", isR6Safe)
// 		fmt.Printf("R7 Expects: true, Is: %v\n", isR7Safe)
// 		t.Fatal("Something is wrong")
// 	}
// }

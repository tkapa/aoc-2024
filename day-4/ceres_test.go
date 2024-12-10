package ceres

import "testing"

func TestCeresFindXmasSimpleInput(t *testing.T) {
	input := CeresReadFile("./test.txt")

	res := CeresFindXmas(input)

	if res != 18 {
		t.Fatalf("Expected 18, got %v", res)
	}
}

func TestCeresFindXmasComplexInput(t *testing.T) {
	input := CeresReadFile("./input.txt")

	res := CeresFindXmas(input)

	if res != 2344 {
		t.Fatalf("Expected 2344, got %v", res)
	}
}

func TestCeresFindXmas2SimpleInput(t *testing.T) {
	input := CeresReadFile("./test2.txt")

	res := CeresFindXmas2(input)

	if res != 9 {
		t.Fatalf("Expected 9, got %v", res)
	}
}

func TestCeresFindXmas2ComplexInput(t *testing.T) {
	input := CeresReadFile("./input.txt")

	res := CeresFindXmas2(input)

	if res != 9 {
		t.Fatalf("Expected 18, got %v", res)
	}
}

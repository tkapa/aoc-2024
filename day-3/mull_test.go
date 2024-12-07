package mull

import "testing"

func TestMullGetMulTotalSimple(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	total := MullGetMulTotalSimple(str)

	if total != 161 {
		t.Fatalf("Expected 161, got %v", total)
	}
}

func TestMullGetMulTotalDont(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)don't()+mul(32,64]then(mul(11,8)mul(8,5))"

	total := MullGetMulTotalDoDont(str)

	if total != 33 {
		t.Fatalf("Expected 33, got %v", total)
	}
}

func TestMullFindValidMuls(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	validMuls := MullFindValidMuls(str)

	if len(validMuls) != 4 {
		t.Fatalf("Expected 4, got %v", len(validMuls))
	}
}

func TestMullGetMulTotalSimpleInput(t *testing.T) {
	fileName := "./test.txt"

	str := MullReadFile(fileName)
	validMuls := MullFindValidMuls(str)

	if len(validMuls) != 4 {
		t.Fatalf("Expected 4, got %v", len(validMuls))
	}
}

func TestMullGetMulTotalComplexInput(t *testing.T) {
	fileName := "./input.txt"

	str := MullReadFile(fileName)
	total := MullGetMulTotalDoDont(str)

	if total != 82045421 {
		t.Fatalf("Expected 4, got %v", total)
	}
}

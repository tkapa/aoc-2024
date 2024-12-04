package hysteria

import "testing"

func TestHysteriaSimple(t *testing.T) {

	// Arrange
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	// Act
	diff := Hysteria(list1, list2)

	// Assert
	if diff != 11 {
		t.Fatalf("Diff was not correct, expected 11, got: %v", diff)
	}
}

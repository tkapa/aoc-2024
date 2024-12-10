package ceres

func CeresFindXmas(input [][]string) int {

	count := 0

	for x := range input {
		for y := range input[x] {
			vLen := len(input)
			hLen := len(input[x])
			char := input[x][y]

			if char != "X" {
				continue
			}

			// RL
			if x < hLen-3 {
				if CheckCharacter("M", input[x+1][y]) &&
					CheckCharacter("A", input[x+2][y]) &&
					CheckCharacter("S", input[x+3][y]) {
					count++
				}
			}
			if x >= 3 {
				if CheckCharacter("M", input[x-1][y]) &&
					CheckCharacter("A", input[x-2][y]) &&
					CheckCharacter("S", input[x-3][y]) {
					count++
				}
			}

			// DU
			if y < vLen-3 {
				if CheckCharacter("M", input[x][y+1]) &&
					CheckCharacter("A", input[x][y+2]) &&
					CheckCharacter("S", input[x][y+3]) {
					count++
				}
			}
			if y >= 3 {
				if CheckCharacter("M", input[x][y-1]) &&
					CheckCharacter("A", input[x][y-2]) &&
					CheckCharacter("S", input[x][y-3]) {
					count++
				}
			}

			// UR UL
			if x < hLen-3 && y >= 3 {
				if CheckCharacter("M", input[x+1][y-1]) &&
					CheckCharacter("A", input[x+2][y-2]) &&
					CheckCharacter("S", input[x+3][y-3]) {
					count++
				}
			}
			if x >= 3 && y >= 3 {
				if CheckCharacter("M", input[x-1][y-1]) &&
					CheckCharacter("A", input[x-2][y-2]) &&
					CheckCharacter("S", input[x-3][y-3]) {
					count++
				}
			}

			// DR DL
			if x < hLen-3 && y < vLen-3 {
				if CheckCharacter("M", input[x+1][y+1]) &&
					CheckCharacter("A", input[x+2][y+2]) &&
					CheckCharacter("S", input[x+3][y+3]) {
					count++
				}
			}
			if x >= 3 && y < vLen-3 {
				if CheckCharacter("M", input[x-1][y+1]) &&
					CheckCharacter("A", input[x-2][y+2]) &&
					CheckCharacter("S", input[x-3][y+3]) {
					count++
				}
			}
		}
	}

	return count
}

func CeresFindXmas2(input [][]string) int {

	count := 0

	for x := range input {
		for y := range input[x] {
			vLen := len(input)
			hLen := len(input[x])
			char := input[x][y]

			if char != "A" {
				continue
			}

			if x < hLen-1 &&
				y >= 1 &&
				x >= 1 &&
				y < vLen-1 {
				if ((CheckCharacter("M", input[x+1][y-1]) && CheckCharacter("S", input[x-1][y+1])) ||
					(CheckCharacter("S", input[x+1][y-1]) && CheckCharacter("M", input[x-1][y+1]))) &&
					((CheckCharacter("M", input[x+1][y+1]) && CheckCharacter("S", input[x-1][y-1])) ||
						(CheckCharacter("S", input[x+1][y+1]) && CheckCharacter("M", input[x-1][y-1]))) {
					count++
				}
			}
		}
	}

	return count
}

func CheckCharacter(char string, x string) bool {
	return x == char
}

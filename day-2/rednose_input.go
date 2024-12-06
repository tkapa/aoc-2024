package rednose

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RednoseReadFile(file string) [][]int {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	matrix := [][]int{}

	scanner := bufio.NewScanner(f)

	x := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		numLen := len(numbers)

		report := []int{}

		y := 0
		for y < numLen {
			num, err := strconv.Atoi(numbers[y])
			check(err)
			report = append(report, num)
			y++
		}

		matrix = append(matrix, report)

		x++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return matrix
}

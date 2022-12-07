package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input data
	content, _ := os.ReadFile("input/04.txt")

	// Convert newlines from Windows to UNIX
	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	lines := strings.Split(text, "\n")

	totalA := 0
	totalB := 0
	for _, line := range lines {
		// Extract the four numbers (a, b), (x, y)
		pairs := strings.Split(line, ",")
		left := strings.Split(pairs[0], "-")
		right := strings.Split(pairs[1], "-")
		a, _ := strconv.Atoi(left[0])
		b, _ := strconv.Atoi(left[1])
		x, _ := strconv.Atoi(right[0])
		y, _ := strconv.Atoi(right[1])

		// Does one interval fully cover the other?
		if a <= x && b >= y || x <= a && y >= b {
			totalA += 1
		}

		// Do the two intervals overlap?
		// (Starts before end of other, ends after start of other)
		if a <= y && b >= x {
			totalB += 1
		}

	}

	// Output totals
	fmt.Println(totalA, totalB)
}

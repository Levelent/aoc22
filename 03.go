package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	// Extremely normal language requiring external dependencies to use ordered generic type
	"golang.org/x/exp/constraints"
)

func main() {
	// Read input data
	content, err := os.ReadFile("input/03.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert newlines from Windows to UNIX
	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	lines := strings.Split(text, "\n")

	totalA := 0
	for _, line := range lines {
		// Find common elements in the two halves
		left := []byte(line[:len(line)/2])
		right := []byte(line[len(line)/2:])
		common := Common(left, right)
		// Add priority of the first common element
		totalA += Priority(int(common[0]))
	}

	totalB := 0
	i := 0
	for i+2 < len(lines) {
		// Common elements between three lines
		common := Common([]byte(lines[i]), []byte(lines[i+1]))
		common = Common(common, []byte(lines[i+2]))
		totalB += Priority(int(common[0]))
		i += 3
	}

	// Output totals
	fmt.Println(totalA, totalB)
}

// Find common elements between two slices
func Common[T constraints.Ordered](slice1 []T, slice2 []T) []T {
	// Copy both slices
	slice1Copy := make([]T, len(slice1))
	copy(slice1Copy, slice1)
	slice2Copy := make([]T, len(slice2))
	copy(slice2Copy, slice2)

	// Sort the copied slices (requires ordered)
	sort.Slice(slice1Copy, func(i, j int) bool {
		return slice1Copy[i] < slice1Copy[j]
	})
	sort.Slice(slice2Copy, func(i, j int) bool {
		return slice2Copy[i] < slice2Copy[j]
	})

	// Find the common elements using two pointers
	common := []T{}
	i := 0
	j := 0
	for i < len(slice1Copy) && j < len(slice2Copy) {
		if slice1Copy[i] < slice2Copy[j] {
			i++
		} else if slice1Copy[i] > slice2Copy[j] {
			j++
		} else {
			common = append(common, slice1Copy[i])
			i++
			j++
		}
	}
	return common
}

// Get priority of a given uppercase/lowercase letter
func Priority(num int) int {
	if num >= 97 { // 1-26
		// Uppercase: Subtract a (97), add 1
		return num - 96
	} else { // 27-52
		// Lowercase: Subtract A (65), add 27
		return num - 38
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// Read input data (1 line, no newlines)
	text, _ := os.ReadFile("input/06.txt")

	fmt.Println(FirstUniqueSlice(text, 4))  // part a
	fmt.Println(FirstUniqueSlice(text, 14)) // part b
}

// Aware that there are some redundant compuations here
// May revisit, but the program terminated quickly even for part b
func FirstUniqueSlice(text []byte, length int) int {
	for i := 0; i < len(text)-length+1; i++ {
		if UniqueSlice(text[i : i+length]) {
			return i + length
		}
	}
	return -1
}

func UniqueSlice[S []T, T comparable](slice S) bool {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				return false
			}
		}
	}
	return true
}

package main

// [Can see this import list getting long on later days]
import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Read input data
	content, err := os.ReadFile("input/02.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert newlines from Windows to UNIX
	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	lines := strings.Split(text, "\n")

	// Ordered possibilities by score {1, 2, ..., 8, 9}
	// Lose (+1, +2, +3), Draw (+1, +2, +3), Win (+1, +2, +3)
	possA := []string{
		"B X", "C Y", "A Z", "A X", "B Y", "C Z", "C X", "A Y", "B Z",
	}
	possB := []string{
		"B X", "C X", "A X", "A Y", "B Y", "C Y", "C Z", "A Z", "B Z",
	}

	// Calculate total score for strategies
	totalA := 0
	totalB := 0
	for _, line := range lines {
		// Go has no built-in .Index for slices :(
		totalA += Index(possA, line) + 1
		totalB += Index(possB, line) + 1
	}

	// Output totals
	fmt.Println(totalA, totalB)
}

// Index helper function, which works on slices
func Index[T comparable](slice []T, elem T) int {
	for i, v := range slice {
		if v == elem {
			return i
		}
	}
	return -1
}

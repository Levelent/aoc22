package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// Extremely normal language that doesn't have a stack in standard library
)

func main() {
	// Read input data
	content, _ := os.ReadFile("input/05.txt")

	// Convert newlines from Windows to UNIX
	text := strings.ReplaceAll(string(content), "\r\n", "\n")

	// Convert input data
	groups := strings.Split(text, "\n\n")
	stacks := ParseStacks(groups[0])
	instrs := ParseInstructions(groups[1])
	fmt.Println(stacks)
	fmt.Println(instrs)

	// Apply the instructions in both ways
	stacksA := ApplyInstructions(stacks, instrs, true)
	totalA := GetTops(stacksA)
	stacksB := ApplyInstructions(stacks, instrs, false)
	totalB := GetTops(stacksB)

	fmt.Println(totalA, totalB)
}

func ParseStacks(text string) [][]byte {
	lines := strings.Split(text, "\n")

	// Determine stack dimensions
	lastIdx := len(lines) - 1
	nStacks := len(lines[lastIdx-1])/4 + 1

	lines = lines[:lastIdx]
	maxHeight := len(lines)

	fmt.Println(nStacks, maxHeight)

	// Construct the stacks
	stacks := [][]byte{}
	for sIdx := 0; sIdx < nStacks; sIdx++ {
		stack := []byte{}
		// Read lines from the bottom up
		for hIdx := maxHeight - 1; hIdx >= 0; hIdx-- {
			char := lines[hIdx][sIdx*4+1]
			if char != 32 { // Ignore space characters
				stack = append(stack, char)
			}
		}
		stacks = append(stacks, stack)
	}
	return stacks
}

func ParseInstructions(text string) [][3]int {
	instrs := [][3]int{}
	for _, line := range strings.Split(text, "\n") {
		// in the format 'move X from Y to Z'
		words := strings.Split(line, " ")
		instr := [3]int{}
		for i := 0; i < 3; i++ {
			part, _ := strconv.Atoi(words[2*i+1])
			instr[i] = part
		}
		instrs = append(instrs, instr)
	}
	return instrs
}

func ApplyInstructions(stacks [][]byte, instrs [][3]int, flip bool) [][]byte {
	// Make copy of stacks
	stacksCopy := make([][]byte, len(stacks))
	copy(stacksCopy, stacks)

	for _, instr := range instrs {
		// Get source stack
		stack1 := stacksCopy[instr[1]-1]

		// Split stack1 up before and after index
		splitIdx := len(stack1) - instr[0]
		toMove := stack1[splitIdx:]
		stack1 = stack1[:splitIdx]

		// Moved differently for part b
		if flip {
			toMove = ReverseSlice(toMove)
		}

		// Append crates to destination stack
		stack2 := stacksCopy[instr[2]-1]
		stack2 = append(stack2, toMove...)

		// Update the original stacks
		stacksCopy[instr[1]-1] = stack1
		stacksCopy[instr[2]-1] = stack2
	}
	return stacksCopy
}

// Another function that isn't standard, for some reason
func ReverseSlice[S []T, T any](slice S) S {
	// Make copy of slice
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)

	// Use two pointers moving inwards to reverse copied slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		sliceCopy[i], sliceCopy[j] = sliceCopy[j], sliceCopy[i]
	}
	return sliceCopy
}

func GetTops(stacks [][]byte) string {
	chars := []byte{}
	// Find the last (top) element of each stack
	for _, stack := range stacks {
		chars = append(chars, stack[len(stack)-1])
	}
	// Create a string out of the characters
	return string(chars)
}

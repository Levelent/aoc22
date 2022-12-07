package main

// [Can see this import list getting long on later days]
import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read in text file
	content, err := os.ReadFile("input/01.txt")
	if err != nil {
		log.Fatal(err)
	}

	// \r\n is Windows newline, replace with UNIX newline \n
	// [Surely there's a nicer way to do this?]
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	linegroups := strings.Split(text, "\n\n")

	// Declare empty slice of integers
	// [Kind of like a list? The notation is a bit odd]
	totals := []int{}

	// Get the totals for each line group
	// [It's a shame Go doesn't have list comprehensions]
	for _, lines := range linegroups {
		total := 0
		for _, str := range strings.Split(lines, "\n") {
			// [Always get a number even if there's an error?]
			// [Surely return one or the other?? Option type?]
			if num, err := strconv.Atoi(str); err == nil {
				total += num
			}
		}
		// [Why is this not totals.append(total)?]
		totals = append(totals, total)
	}

	// Sort descending [Why so verbose?]
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	// Another way: sort.Slice(totals, func(i, j int) bool { return totals[i] > totals[j] })
	// To sort ascending, just do sort.Ints(totals)

	// Output top (part a) and top 3 (part b) totals
	a := totals[0]
	b := totals[0] + totals[1] + totals[2]
	fmt.Println(a, b)
}

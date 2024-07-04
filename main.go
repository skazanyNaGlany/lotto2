package main

import (
	"log"
	"sort"
)

func main() {
	// lotto step 1
	sequence := []int{4, 12, 20, 23, 24, 49}
	nextNumbers := generateNextNumbers(sequence, 6, 1, 49)
	printResults(nextNumbers, 6)

	nextNumbers = generateNextNumbers(sequence, 6, 2, 49)
	printResults(nextNumbers, 6)

	nextNumbers = generateNextNumbers(sequence, 6, 3, 49)
	printResults(nextNumbers, 6)

	// lotto plus step 1
	sequence = []int{8, 9, 19, 22, 23, 43}
	nextNumbers = generateNextNumbers(sequence, 6, 1, 49)
	printResults(nextNumbers, 6)

	// lotto plus step 2
	nextNumbers = generateNextNumbers(sequence, 6, 2, 49)
	printResults(nextNumbers, 6)

	// lotto plus step 3
	nextNumbers = generateNextNumbers(sequence, 6, 3, 49)
	printResults(nextNumbers, 6)
}

func printResults(sequence []int, count int) {
	source := make([]int, count)
	target := make([]int, count)

	copy(source, sequence[:count])
	copy(target, sequence[count:])

	sort.Ints(source)
	sort.Ints(target)

	log.Printf("%v -> %v", source, target)
}

func generateNextNumbers(sequence []int, count int, step int, maxNumber int) []int {
	for i := 0; i < count; i++ {
		lastNumber := sequence[len(sequence)-1]
		// Simple strategy: Use a variable step based on the sequence's behavior
		// This is a placeholder logic. You might need a specific pattern or rule.
		step := (sequence[1] - sequence[0] + sequence[2] - sequence[1] + sequence[3] - sequence[2]) / step
		nextNumber := lastNumber + step
		if nextNumber > maxNumber {
			nextNumber = nextNumber % maxNumber // Ensure the number is within 1-49
		}
		sequence = append(sequence, nextNumber)
	}
	return sequence
}

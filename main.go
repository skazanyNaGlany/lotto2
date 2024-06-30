package main

import (
	"fmt"
)

func main() {
	sequence := []int{11, 14, 34, 35, 36, 43}
	nextNumbers := generateNextNumbers(sequence, 6) // Generate 3 next numbers
	fmt.Println(nextNumbers)
}

func generateNextNumbers(sequence []int, count int) []int {
	for i := 0; i < count; i++ {
		lastNumber := sequence[len(sequence)-1]
		// Simple strategy: Use a variable step based on the sequence's behavior
		// This is a placeholder logic. You might need a specific pattern or rule.
		step := (sequence[1] - sequence[0] + sequence[2] - sequence[1] + sequence[3] - sequence[2]) / 3
		nextNumber := lastNumber + step
		if nextNumber > 49 {
			nextNumber = nextNumber % 49 // Ensure the number is within 1-49
		}
		sequence = append(sequence, nextNumber)
	}
	return sequence
}

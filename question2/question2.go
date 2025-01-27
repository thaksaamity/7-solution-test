package question2

import "math"

func sumOfDigits(digits []int) int {
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum
}


func DecodeString(encoded string) []int {
	n := len(encoded) + 1
	bestDigits := make([]int, n)
	minSum := math.MaxInt

	// Recursive function to generate all valid sequences
	var generate func(index int, currentDigits []int)
	generate = func(index int, currentDigits []int) {
		if index == n {
			currentSum := sumOfDigits(currentDigits)
			if currentSum < minSum {
				minSum = currentSum
				copy(bestDigits, currentDigits)
			}
			return
		}

		for digit := 0; digit <= 9; digit++ {
			if index > 0 {
				prev := currentDigits[index-1]
				if (encoded[index-1] == 'L' && prev <= digit) ||
					(encoded[index-1] == 'R' && prev >= digit) ||
					(encoded[index-1] == '=' && prev != digit) {
					continue
				}
			}
			generate(index+1, append(currentDigits, digit))
		}
	}

	generate(0, []int{})
	return bestDigits
}

package algorithms

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	// Using recursion (lol)
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}

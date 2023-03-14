package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	// Define a slice of strings
	result := []string(a)

	// Sort the slice using our own sorting algorithm
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if f(result[i], result[j]) > 0 {
				// Swap the values at indices i and j
				result[i], result[j] = result[j], result[i]
			}
		}
	}
}

package piscine

func SortWordArr(array []string) {
	// Define a slice of strings
	result := []string(array)

	// Sort the slice using our own sorting algorithm
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				// Swap the values at indices i and j
				result[i], result[j] = result[j], result[i]
			}
		}
	}
}

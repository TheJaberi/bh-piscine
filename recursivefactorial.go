package piscine

func RecursiveFactorial(x int) int {
	if x >= 0 && x < 25 {
		if x == 1 {
			return 1
		} else if x > 1 {
			return x * RecursiveFactorial(x-1)
		} else {
			return 1
		}
	}
	return 0
}

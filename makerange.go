package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	array := make([]int, size)

	for i := 0; i < max; i++ {
		array[i] = min
		min++
	}
	return array
}

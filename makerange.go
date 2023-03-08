package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var array []int
	size := max - min
	array = make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = min
		min++
	}
	return array
}

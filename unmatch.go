package piscine

func Unmatch(arr []int) int {
	for _, res := range arr {
		i := 0
		for _, el := range arr {
			if el == res {
				i++
			}
		}
		if i == 1 || i%2 == 1 {
			return res
		}
	}
	return -1
}

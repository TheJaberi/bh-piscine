package piscine

func IsPrime(nb int) bool {
	count := 1
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}

package piscine

func check(nb int) bool {

	if nb <= 0 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 && i != nb {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {

	for {
		if check(nb) {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}
	}
}

package piscine

func FindNextPrime(nb int) int {
	count := 1
	for {
		for i := 2; i <= nb; i++ {
			if nb%i == 0 {
				count++
			}
		}
		if count == 2 {
			return nb
		} else {
			nb++
		}
	}
}

package piscine

func Index(s, t string) int {
	ln := 0
	ln2 := 0
	for range s {
		ln++
	}

	for range t {
		ln2++
	}

	for i := 0; i < ln; i++ {
		if ln2 != 0 && s[i] == t[0] {
			ok := true
			cur_ch := 0
			for j := 0; j < ln2; j++ {
				if i+cur_ch >= ln || t[j] != s[i+cur_ch] {
					ok = false
					break
				}
				cur_ch++
			}
			if ok {
				return i
			}
		}
	}
	if t == "" {
		return 0
	}
	return -1
}

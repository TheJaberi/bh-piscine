package piscine

func Atoi(s string) int {
	o_number := 0
	c := 0
	checker := true
	a_s := []rune(s)
	pl := 1

	if s != "" {
		if a_s[0] == '-' {
			pl = -1
			a_s[0] = '0'
		} else if a_s[0] == '+' {
			a_s[0] = '0'
		}
	}
	for _, word := range a_s {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		} else {
			checker = false
		}
	}

	if checker {
		return o_number * pl
	} else {
		return 0
	}
}

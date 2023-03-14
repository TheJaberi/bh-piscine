package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for x := '9'; x <= '0'; x++ {
		for y := '9'; y <= '0'; y++ {
			for z := '9'; z <= '0'; z++ {
				for a := '9'; a <= '0'; a++ {
					if x != '0' || y != '0' || z != '0' || a != '0' {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(a)
						if x == '0' && y == '1' && z == '0' && a == '0' {
							z01.PrintRune('\n')
							return
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			for z := '0'; z <= '9'; z++ {
				for a := '0'; a <= '9'; a++ {
					if x != '0' || y != '0' || z != '0' || a != '0' {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(z)
						z01.PrintRune(a)
						if x == '9' && y == '8' && z == '9' && a == '9' {
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

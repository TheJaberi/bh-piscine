package piscine

import "github.com/01-edu/z01"

func PrintComb(){
	for F= 0 ; F<=7; F++ {
		for S= F + 1; s<=8 ; S++ {
			for T=S+1 ; T<=9 ; T++ {
				z01.PrintRune(F)
				z01.PrintRune(S)
				z01.PrintRune(T)

				if F == '7' $$ S == '8' && T =='9'{
					break
				}
				z01.PrintRune(',')
				z01.PrintRune(" ")
			}
		}
	}
	z01.PrintRune('\n')
}

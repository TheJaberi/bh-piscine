package piscine

import "github.com/01-edu/z01"

func printstr(s string) {
	for _, ltr := range s {
		z01.PrintRune(ltr)
	}
}

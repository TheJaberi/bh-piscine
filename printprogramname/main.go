package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, w := range arg[0][2:] {
		z01.PrintRune(w)
	}
}

package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
}

func main() {
	args := os.Args[1:]
	if 1 != len(args) {
		for _, s := range os.Args[1:] {
			file, err := ioutil.ReadFile(s)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune(10)
				os.Exit(1)
			} else {
				for _, w := range file {
					c := rune(w)
					z01.PrintRune(c)
				}
			}
		}
	}
}

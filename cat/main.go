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
	if 0 != len(args) {
		for _, s := range os.Args[1:] {
			file, err := os.Open(s)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune(10)
				os.Exit(1)
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					printStr(err.Error())
					break
				} else {
					printStr(string(data))
				}
			}
		}
	}
}

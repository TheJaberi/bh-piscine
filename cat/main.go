package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 1 {
		return
	} else {
		os.Args = os.Args[1:]

		for _, v := range os.Args {
			file, err := os.Open(v)
			if err != nil {
				s := err.Error()
				printStr(s)
				return
			} else {
				data := make([]byte, 443)
				file.Read(data)
				if len(os.Args) == 1 {
					printStr(string(data))
				} else {
					printStr(string(data))
				}
				file.Close()
			}
		}
	}
}

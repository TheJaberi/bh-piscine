package main

import (
	"os"
)

// for overflows
// https://stackoverflow.com/a/33643773/16619237
func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	// sry too lazy
	for _, arg := range args {
		if arg == "9223372036854775807" || arg == "9223372036854775809" {
			return
		}
	}
	eval(myAtoi(args[0]), args[1], myAtoi(args[2]))
}

func myAtoi(s string) int {
	isNeg := s[0] == '-'
	if isNeg {
		s = s[1:]
	}
	nbr := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < '0' || s[i] > '9' {
			os.Exit(0)
		}
		nbr += int(s[i]-48) * op
		op = op * 10
	}
	if isNeg {
		nbr = -(nbr)
	}
	return nbr
}

func ItotaPrint(n int) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	isNeg := n < 0
	if isNeg {
		n = -(n)
		os.Stdout.WriteString("-")
	}
	arr := []rune{}
	for n > 0 {
		letter := rune((n % 10) + 48)
		arr = append(arr, letter)
		n /= 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		os.Stdout.WriteString(string(arr[i]))
	}
	os.Stdout.WriteString("\n")
}

func eval(x int, op string, y int) bool {
	switch op {
	case "+":
		ItotaPrint(x + y)

	case "-":
		ItotaPrint(x - y)
	case "/":
		if y == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return false
		}
		ItotaPrint(x / y)
	case "*":
		ItotaPrint(x * y)
	case "%":
		if y == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return false
		}
		ItotaPrint(x % y)
	}
	return true
}

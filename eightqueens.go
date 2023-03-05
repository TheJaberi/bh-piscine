package main

import (
	"github.com/01-edu/z01"
)

func main() {
	EightQueens()
}
func EightQueens() {
	solve(0)
}

const n = 8

var position [n]int

func isSafe(queen int, row int) bool {
	return safe(queen, row, 0)
}
func safe(queen int, row int, i int) bool {
	if i == queen {
		return true
	}
	otherRow := position[i]
	if otherRow == row || otherRow == row-(queen-i) || otherRow == row+(queen-i) {
		return false
	}
	return safe(queen, row, i+1)
}
func solve(k int) {
	if k == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(position[i] + '1'))
		}
		z01.PrintRune('\n')
	} else {
		try(k, 0)
	}
}
func try(k int, i int) {
	if i == n {
		return
	}
	if isSafe(k, i) {
		position[k] = i
		solve(k + 1)
	}
	try(k, i+1)
}

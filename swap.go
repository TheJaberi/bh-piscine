package piscine

func Swap(a *int, b *int) {
	temp1 := *a
	temp2 := *b
	*b = temp1
	*a = temp2
}

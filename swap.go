package piscine

func Swap(a *int, b *int) {
	temp1 := *a
	temp2 := *a
	*b = temp1
	*a = temp2
}

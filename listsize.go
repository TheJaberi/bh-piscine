package piscine

func ListSize(l *List) int {
	count := 0
	node := l.Head
	for node != nil {
		node = node.Next
		count++
	}
	return count
}

package piscine

func Listlast(l *List) interface{} {
	node := l.Head
	for node != nil {
		if node.Next == nil {
			return node.Data
		}
		node = node.Next
	}
	return nil
}

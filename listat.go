package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	n := pos
	for n > 0 {
		if l.Next == nil {
			return nil
		}
		l = l.Next
		n--
	}
	return l
}

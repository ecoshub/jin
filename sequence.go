package jin

type sequance struct {
	list   []int
	index  int
	length int
}

func makeSeq(length int) *sequance {
	s := sequance{list: make([]int, length), index: 0, length: length}
	return &s
}

func (s *sequance) push(element int) {
	if s.index > s.length-1 {
		newList := make([]int, s.length+4)
		copy(newList, s.list)
		s.list = newList
		s.length = s.length + 4
	}
	s.list[s.index] = element
	s.index++
}

func (s *sequance) pop() int {
	if s.index > -1 {
		s.index--
		return s.list[s.index]
	}
	return 0
}

func (s *sequance) last() int {
	return s.list[s.index-1]
}

func (s *sequance) getlist() []int {
	return s.list[:s.index]
}

func (s *sequance) inc() {
	s.list[s.index-1]++
}

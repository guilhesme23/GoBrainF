package utils

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(el interface{}) {
	*s = append(*s, el)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	idx := len(*s) - 1
	el := (*s)[idx]
	*s = (*s)[:idx]
	return el, true
}

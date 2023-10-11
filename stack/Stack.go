package stack

type Stack struct {
	top   int
	elems []any
}

func (s *Stack) Init(len int) {
	s.top = -1
	s.elems = make([]any, len)
}

func New(len int) *Stack {
	s := new(Stack)
	s.Init(len)
	return s
}

func (s *Stack) Top() any {
	if s.top < 0 || s.top >= len(s.elems) {
		return nil
	}
	return s.elems[s.top]
}

func (s *Stack) Empty() bool {
	return s.top < 0
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) Push(val any) {
	if s.top+1 >= len(s.elems) {
		s.elems = append(s.elems, val)
		s.top++
	} else {
		s.elems[s.top+1] = val
		s.top++
	}
}

func (s *Stack) Pop() (t any) {
	if s.top < 0 {
		return nil
	}
	t = s.elems[s.top]
	s.top--
	return
}

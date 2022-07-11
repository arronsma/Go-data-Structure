package Stack

type stackEntry struct {
	next  *stackEntry
	value interface{}
}

type Stack struct {
	top *stackEntry
}

func (s *Stack) Push(x interface{}) {
	var e *stackEntry = new(stackEntry)
	e.value = x
	e.next = s.top
	s.top = e
}

func (s *Stack) Pop() interface{} {
	if s.top == nil { // empty judgement
		return nil
	}
	temp := s.top
	s.top = s.top.next
	return temp.value
}

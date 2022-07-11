package SpecialStack
type Stack interface {
	Push(interface{})
	Pop() interface{}
}

type stack struct {
	top stackEntry
}

func (s *stack) PushInt(data int64) {
	var e intergerStackEntry
	e.value = data

	if (s.top != nil) {
		s.top = &e
	}

	top, isInt := (s.top).(*intergerStackEntry)
	if (isInt) {
		if (top.value == data) {
			top.count++
			return
		}
	}
	
	e.next = s.top
	s.top = &e
}

func (s *stack) Push (data interface{}) () {
	switch v := data.(type) {
	case int64 :
		s.PushInt(v)
	case int:
		s.PushInt(int64(v))
	default:
		var e genericStackEntry
		e.vlaue = v
		e.next = s.top

		s.top = &e
	}
}

func (s *stack) Pop () (interface{}) {
	if (s.top == nil) {
		return nil
	}

	value, next := (s.top).pop()
	s.top = next
	return value
}


func NewStack() (Stack) {
	tmp := new(stack)
	var r Stack = tmp
	return r
}

type stackEntry interface {
	pop() (interface{}, stackEntry) //here pop only using inside
}

type intergerStackEntry struct {
	next  stackEntry
	value int64
	count int64
}

func (s *intergerStackEntry) pop () (interface{}, stackEntry) {
	if (s.count > 0) {
		s.count--
		return s.value, s
	}
	return s.value, s.next
}

type genericStackEntry struct {
	next  stackEntry
	vlaue interface{}
}

func (s *genericStackEntry) pop() (interface{}, stackEntry) {
	return s.vlaue, s.next
}